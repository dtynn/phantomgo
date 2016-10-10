package system

import (
	"github.com/gopherjs/gopherjs/js"
)

type OS struct {
	obj *js.Object
}

func (this *OS) Architecture() string {
	return this.obj.Get("architecture").String()
}

func (this *OS) Name() string {
	return this.obj.Get("name").String()
}

func (this *OS) Release() string {
	return this.obj.Get("release").String()
}

func (this *OS) Version() string {
	return this.obj.Get("version").String()
}
