package system

import (
	"github.com/gopherjs/gopherjs/js"
)

func NewSystem() *System {
	obj := js.Global.Call("require", "system")

	return &System{
		obj: obj,
	}
}

type System struct {
	obj *js.Object
}

func (this *System) Args() []string {
	o := this.obj.Get("args")

	size := o.Length()
	res := make([]string, size)
	for i := 0; i < size; i++ {
		res[i] = o.Index(i).String()
	}

	return res
}

func (this *System) Env() map[string]string {
	res := map[string]string{}

	o := this.obj.Get("env")

	keys := js.Keys(o)
	for _, key := range keys {
		res[key] = o.Get(key).String()
	}

	return res
}

func (this *System) OS() *OS {
	o := this.obj.Get("os")
	return &OS{
		obj: o,
	}
}

func (this *System) Pid() int {
	return this.obj.Get("pid").Int()
}

func (this *System) Platform() string {
	return this.obj.Get("platform").String()
}
