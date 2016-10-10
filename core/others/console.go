package others

import (
	"github.com/gopherjs/gopherjs/js"
)

func GetConsole() *Console {
	obj := js.Global.Get("console")
	return &Console{
		obj: obj,
	}
}

type Console struct {
	obj *js.Object
}

func (this *Console) Log(msgs ...interface{}) {
	this.obj.Call("log", msgs...)
}
