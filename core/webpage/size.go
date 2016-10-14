package webpage

import (
	"github.com/gopherjs/gopherjs/js"
)

type Size struct {
	obj *js.Object

	Width  int `js:"width"`
	Height int `js:"height"`
}
