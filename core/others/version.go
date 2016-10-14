package others

import (
	"github.com/gopherjs/gopherjs/js"
)

type Version struct {
	obj *js.Object

	Major int `js:"major"`
	Minor int `js:"minor"`
	Patch int `js:"patch"`
}
