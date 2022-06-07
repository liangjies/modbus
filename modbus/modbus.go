package modbus

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Client 接口定义了 Modbus 客户端的接口。
type TCPClientHandler struct {
	// TCP 地址
	Address string
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
}

// 定义常量
const (
	// 定义 TCP 连接的超时时间
	tcpTimeout     = 10 * time.Second
	tcpIdleTimeout = 60 * time.Second
)

// 新建 TCP 客户端
func NewTCPClientHandler(address string) *TCPClientHandler {
	h := &TCPClientHandler{}
	h.Address = address
	h.Timeout = tcpTimeout
	h.IdleTimeout = tcpIdleTimeout
	return h
}

// 建立 TCP 连接
func (mb *TCPClientHandler) connect() error {
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
