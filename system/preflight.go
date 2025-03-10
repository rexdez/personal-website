package system

import (
	"log"
	"os"
	"sync"
)

var initLoggerOnce sync.Once

func (conn *SysConn) Init(){
	conn.logger = InitializeLogger()
}

func InitializeLogger() *log.Logger {
	var logger *log.Logger
	initLoggerOnce.Do(func() {
		file, err :=  os.OpenFile("errors.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
		if err != nil {
			panic("Unable to initiate logger")
		}
		logger = log.New(file, "PW_ROHIT", log.Ldate | log.Ltime | log.Lshortfile)
	})
	return logger
}