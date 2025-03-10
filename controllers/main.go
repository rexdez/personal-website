package controllers

import "net/http"

type InitType interface {
	Init()
}

type Controller struct {
	InitType
}

var blogs map[string]string = map[string]string{
	"3e9" : "blog/aws-ses-golang.html",
}

func (handler *Controller) StaticHandler() {

}

func (handler *Controller) Homepage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "text/html")
}