package webpage

import (
	"github.com/gopherjs/gopherjs/js"
)

type ScrollPosition struct {
	obj *js.Object

	Left int `js:"left"`
	Top  int `js:"top"`
}
