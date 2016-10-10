package filesystem

import (
	"github.com/gopherjs/gopherjs/js"
)

type File struct {
	obj *js.Object
}

func (this *File) AtEnd() bool {
	return this.obj.Call("atEnd").Bool()
}

func (this *File) Close() {
	this.obj.Call("close")
}

func (this *File) Flush() {
	this.obj.Call("flush")
}

func (this *File) GetEncoding() string {
	return this.obj.Call("getEncoding").String()
}

func (this *File) Read(size int) string {
	return this.obj.Call("read", size).String()
}

func (this *File) ReadLine() string {
	return this.obj.Call("readLine").String()
}

func (this *File) Seek(pos int) bool {
	return this.obj.Call("seek", pos).Bool()
}

func (this *File) SetEncoding(encoding string) bool {
	return this.obj.Call("setEncoding", encoding).Bool()
}

func (this *File) Write(data string) bool {
	return this.obj.Call("write", data).Bool()
}

func (this *File) WriteLine(data string) bool {
	return this.obj.Call("writeLine", data).Bool()
}
