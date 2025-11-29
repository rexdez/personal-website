package main

import (
	"log"
	"os"
	"sync"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SysConn struct {
	Logger *log.Logger
	Pool *pgxpool.Pool
}

var InitLoggerOnce sync.Once

func (conn *SysConn) Init() {
	conn.Logger = InitializeLogger()
}

func InitializeLogger() *log.Logger {
	var logger *log.Logger
	InitLoggerOnce.Do(func() {
		file, err :=  os.OpenFile("errors.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
		if err != nil {
			panic("Unable to initiate logger")
		}
		logger = log.New(file, "PW_ROHIT", log.Ldate | log.Ltime | log.Lshortfile)
	})
	return logger
}