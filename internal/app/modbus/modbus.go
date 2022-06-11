package modbus

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"modbus-spyder/internal/app/global"
	"modbus-spyder/internal/app/service"
	"modbus-spyder/internal/app/utils"
	"net"
	"sync"
	"time"

	"go.uber.org/zap"
)

func RunSpyder(addr string, ctx context.Context, wg *sync.WaitGroup) {
	for {
		// 建立 TCP 连接
		mb := NewTCPClientHandler(addr)
		err := mb.Connect()
		if err != nil {
			fmt.Println(err)
		} else {
			// 接收数据--阻塞运行
			if err = mb.Receive(ctx); err != nil {
				fmt.Println(err)
			}
			defer mb.conn.Close() // 关闭连接
		}
		// 结束协程
		select {
		case <-ctx.Done(): //等待通知
			wg.Done() // 协程计数器减1
			return
		default:
		}
		global.SYS_LOG.Info("采集点发生重连", zap.Any("addr", addr))
		// 等待10秒重连
		time.Sleep(10 * time.Second)
	}

}

// Client 接口定义了 Modbus 客户端的接口。
type TCPClientHandler struct {
	// TCP 地址
	Address string
	// 从机地址
	SlaveID byte
	// 连接 & 读取超时
	Timeout time.Duration
	// 超过闲置时间关闭连接
	IdleTimeout time.Duration
	// Transmission logger
	Logger *log.Logger
	// TCP 连接
	conn net.Conn
	// 上次活动时间
	lastActivity time.Time
	// 上次解析成功时间
	lastSuccess time.Time
}

// 定义常量
const (
	// 定义 TCP 连接的超时时间
	tcpTimeout     = 10 * time.Second
	tcpIdleTimeout = 60 * time.Second
	// 数据包最大长度
	tcpMaxLength = 75
)

// 新建 TCP 客户端
func NewTCPClientHandler(address string) *TCPClientHandler {
	h := &TCPClientHandler{}
	h.Address = address
	h.Timeout = tcpTimeout
	h.IdleTimeout = tcpIdleTimeout
	lastSuccess, _ := time.Parse("2006-01-02 15:04:05", "1970-01-02 00:00:00")
	h.lastSuccess = lastSuccess
	return h
}

// 建立 TCP 连接
func (mb *TCPClientHandler) Connect() error {
	if mb.conn == nil {
		fmt.Println("connecting to", mb.Address)
		dialer := net.Dialer{Timeout: mb.Timeout}
		conn, err := dialer.Dial("tcp", mb.Address)
		if err != nil {
			return err
		}
		mb.conn = conn
	}
	return nil
}

// 接收数据
// 以一个串口服务器来解析数据
func (mb *TCPClientHandler) Receive(ctx context.Context) (err error) {

	// 设置超时时间
	mb.lastActivity = time.Now() // 设置上次活动时间
	var timeout time.Time
	if mb.Timeout > 0 {
		timeout = mb.lastActivity.Add(mb.Timeout)
	}
	if err = mb.conn.SetDeadline(timeout); err != nil {
		return
	}
	// 定义一次读多少数据
	var data [tcpMaxLength]byte
	// 定义一个缓冲区
	var buf []byte
	for {
		// 读取数据
		if _, err = io.ReadFull(mb.conn, data[:tcpMaxLength]); err != nil {
			return
		}
		// 数据存入缓冲区
		buf = append(buf, data[:tcpMaxLength]...)
		// 地址和功能码03可以确定
		// 从缓冲数据里查找这两个数据
		for i, v := range buf {
			if v == 3 && i != 0 {
				// 数据包长度
				dataLen := int(buf[i+1])
				// 待校验数据长度不足，需要继续读取
				if len(buf) < i+dataLen+4 {
					break
				}
				// 获取buf中校验码
				checksum := buf[i+dataLen+2 : i+dataLen+4]
				// CRC16校验数据是否有效
				var crc utils.CRC
				crc.Reset()
				crc.PushBytes(buf[i-1 : i+dataLen+2])
				if bytes.Equal(checksum, crc.Value()) {
					// 校验成功，解析数据
					mb.SlaveID = buf[i-1]
					datas := MsgParsing(buf[i-1:i+dataLen+4], mb)
					// if service.PutData(datas) {
					// 	mb.lastSuccess = time.Now()
					// }
					if service.PutDataInMongoDB(datas) {
						mb.lastSuccess = time.Now()
					}

					//fmt.Println("datas:", datas)
					buf = buf[i+dataLen+4:]
					break
				} else {
					continue
				}

			}
		}
		// 检测数据长度,长度超长释放
		if len(buf) > 250 {
			buf = []byte{}
		}

		// 设置超时时间
		mb.lastActivity = time.Now() // 设置上次活动时间
		var timeout time.Time
		if mb.Timeout > 0 {
			timeout = mb.lastActivity.Add(mb.Timeout)
		}
		if err = mb.conn.SetDeadline(timeout); err != nil {
			return
		}
		// 检测是否需要关闭协程
		select {
		case <-ctx.Done(): //等待通知
			return
		default:
		}
	}
}

/*
// 发送数据
func (mb *tcpTransporter) Send(aduRequest []byte) (aduResponse []byte, err error) {
	// 如果连接不存在，则建立连接
	if err = mb.connect(); err != nil {
		return
	}
	// 发送数据
	log.Println("modbus: sending % x", aduRequest)
	if _, err = mb.conn.Write(aduRequest); err != nil {
		return
	}
	// Read header first
	var data [tcpMaxLength]byte
	if _, err = io.ReadFull(mb.conn, data[:tcpHeaderSize]); err != nil {
		return
	}
	// Read length, ignore transaction & protocol id (4 bytes)
	length := int(binary.BigEndian.Uint16(data[4:]))
	if length <= 0 {
		mb.flush(data[:])
		err = fmt.Errorf("modbus: length in response header '%v' must not be zero", length)
		return
	}
	if length > (tcpMaxLength - (tcpHeaderSize - 1)) {
		mb.flush(data[:])
		err = fmt.Errorf("modbus: length in response header '%v' must not greater than '%v'", length, tcpMaxLength-tcpHeaderSize+1)
		return
	}
	// Skip unit id
	length += tcpHeaderSize - 1
	if _, err = io.ReadFull(mb.conn, data[tcpHeaderSize:length]); err != nil {
		return
	}
	aduResponse = data[:length]
	mb.logf("modbus: received % x\n", aduResponse)
	return
}
*/
