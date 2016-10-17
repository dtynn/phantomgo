package system

import (
	"github.com/gopherjs/gopherjs/js"
)

func ObjectToSystem(obj *js.Object) *System {
	return &System{
		obj: obj,
	}
}

func NewSystem() *System {
	return ObjectToSystem(js.Global.Call("require", "system"))
}

type System struct {
	obj *js.Object

	Args           []string          `js:"args"`
	Env            map[string]string `js:"env"`
	Pid            int               `js:"pid"`
	Platform       string            `js:"platform"`
	IsSSLSupported bool              `js:"isSSLSupported"`
	OS             *OS               `js:"os"`
}
