package filesystem

import (
	"strings"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

func ObjectToFileSystem(obj *js.Object) *FileSystem {
	return &FileSystem{
		obj: obj,
	}
}

func NewFileSystem() *FileSystem {
	return ObjectToFileSystem(js.Global.Call("require", "fs"))
}

type FileSystem struct {
	obj *js.Object
}

func (this *FileSystem) Separator() string {
	return this.obj.Get("separator").String()
}

func (this *FileSystem) WorkingDirectory() string {
	return this.obj.Get("workingDirectory").String()
}

func (this *FileSystem) Absolute(path string) string {
	return this.obj.Call("absolute", path).String()
}

func (this *FileSystem) ChangeWorkingDirectory(path string) bool {
	return this.obj.Call("changeWorkingDirectory", path).Bool()
}

// attributes
func (this *FileSystem) Size(path string) int {
	return this.obj.Call("size", path).Int()
}

func (this *FileSystem) LastModified(path string) (time.Time, bool) {
	var t time.Time

	o := this.obj.Call("lastModified", path)
	if o == nil {
		return t, false
	}

	t, err := time.Parse(time.RFC3339Nano, o.String())
	if err != nil {
		return t, false
	}

	return t, true
}

// files
func (this *FileSystem) Open(path string, charset Charset, mode fileMode) *File {
	opt := map[string]string{}
	if modeStr := parseFileMode(mode); modeStr != "" {
		opt["mode"] = modeStr
	}

	if charset != CharsetDefault {
		opt["charset"] = string(charset)
	}

	return &File{
		obj: this.obj.Call("open", path, opt),
	}
}

func (this *FileSystem) Read(path string, charset Charset, readBinary bool) []byte {
	mode := FileModeRead
	if readBinary {
		mode |= FileModeBinary
	}

	opt := map[string]string{}
	if modeStr := parseFileMode(mode); modeStr != "" {
		opt["mode"] = modeStr
	}

	if charset != CharsetDefault {
		opt["charset"] = string(charset)
	}

	return []byte(this.obj.Call("read", path, opt).String())
}

func (this *FileSystem) Write(path, content string, charset Charset, mode fileMode) {
	opt := map[string]string{}
	if modeStr := parseFileMode(mode); modeStr != "" {
		opt["mode"] = modeStr
	}

	if charset != CharsetDefault {
		opt["charset"] = string(charset)
	}

	this.obj.Call("write", path, content, opt)
}

func (this *FileSystem) Remove(path string) {
	this.obj.Call("remove", path)
}

func (this *FileSystem) Move(src, dst string) {
	this.obj.Call("move", src, dst)
}

func (this *FileSystem) Touch(path string) {
	this.obj.Call("touch", path)
}

// dir
func (this *FileSystem) List(path string, all bool) []string {
	res := this.obj.Call("list", path)

	size := res.Length()
	list := make([]string, 0, size)
	for i := 0; i < res.Length(); i++ {
		p := res.Index(i).String()
		if !all && strings.HasPrefix(p, ".") {
			continue
		}

		list = append(list, p)
	}

	return list
}

func (this *FileSystem) CopyTree(src, dst string) {
	this.obj.Call("copyTree", src, dst)
}

func (this *FileSystem) MakeTree(path string) bool {
	return this.obj.Call("makeTree", path).Bool()
}

func (this *FileSystem) MakeDirectory(dir string) bool {
	return this.obj.Call("makeDirectory", dir).Bool()
}

func (this *FileSystem) RemoveTree(path string) {
	this.obj.Call("removeTree", path)
}

func (this *FileSystem) RemoveDirectory(dir string) {
	this.obj.Call("removeDirectory", dir)
}

func (this *FileSystem) Join(path ...string) string {
	args := make([]interface{}, len(path))
	for i, one := range path {
		args[i] = one
	}

	return this.obj.Call("join", args...).String()
}

func (this *FileSystem) Split(path string) []string {
	o := this.obj.Call("split", path)

	size := o.Length()
	res := make([]string, size)
	for i := 0; i < size; i++ {
		res[i] = o.Index(i).String()
	}

	return res
}

func (this *FileSystem) FromNativeSeparators(path string) string {
	return this.obj.Call("fromNativeSeparators", path).String()
}

func (this *FileSystem) ToNativeSeparators(path string) string {
	return this.obj.Call("toNativeSeparators", path).String()
}

// links
func (this *FileSystem) ReadLink(path string) string {
	return this.obj.Call("readLink", path).String()
}

// tests
func (this *FileSystem) Exists(path string) bool {
	return this.obj.Call("exists", path).Bool()
}

func (this *FileSystem) IsDirectory(path string) bool {
	return this.obj.Call("isDirectory", path).Bool()
}

func (this *FileSystem) IsFile(path string) bool {
	return this.obj.Call("isFile", path).Bool()
}

func (this *FileSystem) IsAbsolute(path string) bool {
	return this.obj.Call("isAbsolute", path).Bool()
}

func (this *FileSystem) IsExecutable(path string) bool {
	return this.obj.Call("isExecutable", path).Bool()
}

func (this *FileSystem) IsReadable(path string) bool {
	return this.obj.Call("isReadable", path).Bool()
}

func (this *FileSystem) IsWritable(path string) bool {
	return this.obj.Call("isWritable", path).Bool()
}

func (this *FileSystem) IsLink(path string) bool {
	return this.obj.Call("isLink", path).Bool()
}
