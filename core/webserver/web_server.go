package webserver

import (
	"github.com/gopherjs/gopherjs/js"
)

func ObjectToWebServer(obj *js.Object) *WebServer {
	return &WebServer{
		obj: obj,
	}
}

func NewWebServer() *WebServer {
	return ObjectToWebServer(js.Global.Call("require", "webserver").Call("create"))
}

type WebServer struct {
	obj *js.Object
}
