package controllers

import (
	"net/http"
	"strings"

	"github.com/rexdez/personal-website/common"
)

type Conn interface {
	Init()
}

type Controller struct {
	Conn
}

var blogs map[string]string = map[string]string{
	"3e9" : "blog/aws-ses-golang.html",
}

func (handler *Controller) StaticHandler(root string) http.Handler{
	fs := http.FileServer(http.Dir(root))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		splits := strings.Split(root, ".")
		extension := splits[len(splits)-1]
		w.Header().Set("Content-Type", common.MimeTypes[extension])
		fs.ServeHTTP(w, r)
	})
}

func (handler *Controller) Homepage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "text/html")
}