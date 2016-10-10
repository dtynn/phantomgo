package childprocess

import (
	"github.com/gopherjs/gopherjs/js"
)

type Opt struct {
	Encoding string
}

type ChildProcess struct {
	obj *js.Object
}

func (this *ChildProcess) Spawn(cmd string, args []string, option *Opt) *Context {
	opt := map[string]string{}
	if option != nil && option.Encoding != "" {
		opt["encoding"] = option.Encoding
	}

	o := this.obj.Call("spawn", cmd, args, opt)
	return &Context{
		obj: o,
	}
}
