package phantom

import (
	"github.com/dtynn/phantomgo/core/handlers"
	"github.com/gopherjs/gopherjs/js"
)

func GetPhantomObject() *PhantomObject {
	obj := js.Global.Get("phantom")
	return &PhantomObject{
		obj: obj,
	}
}

type PhantomObject struct {
	obj *js.Object
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
