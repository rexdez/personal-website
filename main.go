package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/rexdez/personal-website/system"
	"github.com/rexdez/personal-website/controllers"
)

func main() {
	conn := &system.SysConn{}
	conn.Init()
	handler := controllers.Controller{InitType: conn}
	router := InitRouter(handler)

	server := &http.Server {
		Addr: ":6432",
		Handler: router,
		ReadTimeout: time.Second*20,
		WriteTimeout: time.Second*30,
		IdleTimeout: time.Minute,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("server could not be started: %v", err))
	}
}