package phantom

import (
	"github.com/dtynn/phantomgo/core/cookiejar"
	"github.com/dtynn/phantomgo/core/handlers"
	"github.com/gopherjs/gopherjs/js"
)

var phantomObject = &PhantomObject{
	obj: js.Global.Get("phantom"),
}

func GetPhantomObject() *PhantomObject {
	return phantomObject
}

type PhantomObject struct {
	obj *js.Object

	Version *Version           `js:"version"`
	Cookies []cookiejar.Cookie `js:"cookies"`
}

func (this *PhantomObject) Exit(code int) {
	this.obj.Call("exit", code)
}

func (this *PhantomObject) OnError() *js.Object {
	return this.obj.Get("onError")
}

func (this *PhantomObject) SetOnError(handler handlers.OnErrorHandler) {
	this.obj.Set("onError", handler)
}

func (this *PhantomObject) AddCookie(cookie cookiejar.Cookie) bool {
	return this.obj.Call("addCookie", cookie).Bool()
}

func (this *PhantomObject) ClearCookies() {
	this.obj.Call("clearCookies")
}
