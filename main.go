package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rexdez/personal-website/controllers"
	"github.com/rexdez/personal-website/system"
)

func main() {
	conn := system.SysConn{}
	conn.Init()
	handler := controllers.Controller{SysConn: conn}
	router := InitRouter(handler)

	server := &http.Server {
		Addr: ":6432",
		Handler: router,
		ReadTimeout: time.Second*20,
		WriteTimeout: time.Second*30,
		IdleTimeout: time.Minute,
	}

	fmt.Printf(`=====================================
Starting Go Server at: http://127.0.0.1:6432
All System Checked & Connected...
=====================================
`)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("server could not be started: %v", err))
	}
	
}