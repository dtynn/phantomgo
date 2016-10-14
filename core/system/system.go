package system

import (
	"github.com/gopherjs/gopherjs/js"
)

func NewSystem() *System {

	return &System{
		obj: js.Global.Call("require", "system"),
	}
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
