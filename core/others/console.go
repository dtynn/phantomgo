package others

import (
	"github.com/gopherjs/gopherjs/js"
)

var consoleObj = &ConsoleObject{
	obj: js.Global.Get("console"),
}

func Console() *ConsoleObject {
	return consoleObj
}

type ConsoleObject struct {
	obj *js.Object
}

func (this *ConsoleObject) Log(msgs ...interface{}) {
	this.obj.Call("log", msgs...)
}
