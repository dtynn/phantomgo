package handlers

import (
	"github.com/gopherjs/gopherjs/js"
)

func NewErrorStack(obj *js.Object) *ErrorStack {
	return &ErrorStack{
		obj: obj,
	}
}

type ErrorStack struct {
	obj *js.Object
}

func (this *ErrorStack) File() string {
	return this.obj.Get("file").String()
}

func (this *ErrorStack) Line() int {
	return this.obj.Get("line").Int()
}

func (this *ErrorStack) Function() string {
	return this.obj.Get("function").String()
}

type OnErrorHandler func(msg string, stack []*ErrorStack)
