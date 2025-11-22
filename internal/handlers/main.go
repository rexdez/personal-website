package handlers

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

type Handlers struct {
	system.SysConn
}

func (conn *Handlers) StaticHandler(root string) http.Handler{
	fs := http.FileServer(http.Dir(root))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		splits := strings.Split(root, ".")
		extension := splits[len(splits)-1]
		w.Header().Set("Content-Type", common.MimeTypes[extension])
		fs.ServeHTTP(w, r)
	})
}

