package phantom

import (
	"github.com/dtynn/phantomgo/core/handlers"
	"github.com/dtynn/phantomgo/core/others"
	"github.com/gopherjs/gopherjs/js"
)

var phantomObject = &PhantomObject{
	Object: js.Global.Get("phantom"),
}

func GetPhantomObject() *PhantomObject {
	return phantomObject
}

type PhantomObject struct {
	*js.Object

	Version *others.Version `js:"version"`
	Cookies []others.Cookie `js:"cookies"`
}

func (this *PhantomObject) Exit(code int) {
	this.Call("exit", code)
}

func (this *PhantomObject) OnError() *js.Object {
	return this.Get("onError")
}

func (this *PhantomObject) SetOnError(handler handlers.OnErrorHandler) {
	this.Set("onError", handler)
}

func (this *PhantomObject) AddCookie(cookie others.Cookie) bool {
	return this.Call("addCookie", cookie).Bool()
}

func (this *PhantomObject) ClearCookies() {
	this.Call("clearCookies")
}
