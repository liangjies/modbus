package utils

import (
	"modbus-spyder/internal/app/global"
	"os"

	"go.uber.org/zap"
)

//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//@function: CreateDir
//@description: 批量创建文件夹
//@param: dirs ...string
//@return: err error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.SYS_LOG.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.SYS_LOG.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
