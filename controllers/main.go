package controllers

import (
	"net/http"
	"strings"

	"github.com/rexdez/personal-website/common"
	"github.com/rexdez/personal-website/system"
	"github.com/rexdez/personal-website/templates"
)

type Conn interface {
	Init()
}

type Controller struct {
	system.SysConn
}

var blogs map[string]string = map[string]string{
	"3e9" : "blog/aws-ses-golang.html",
}

func (conn *Controller) StaticHandler(root string) http.Handler{
	fs := http.FileServer(http.Dir(root))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		splits := strings.Split(root, ".")
		extension := splits[len(splits)-1]
		w.Header().Set("Content-Type", common.MimeTypes[extension])
		fs.ServeHTTP(w, r)
	})
}

func (conn *Controller) Homepage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "text/html")
	err := templates.Homepage.Execute(w, nil)
	if err != nil {
		go conn.Logger.Println(err)
	}
}