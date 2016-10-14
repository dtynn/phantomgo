package webpage

import (
	"github.com/gopherjs/gopherjs/js"
)

type ClipRect struct {
	obj *js.Object

	Top    int `js:"top"`
	Left   int `js:"left"`
	Width  int `js:"width"`
	Height int `js:"height"`
}
