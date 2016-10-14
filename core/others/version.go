package others

import (
	"github.com/gopherjs/gopherjs/js"
)

type Version struct {
	*js.Object

	Major int `js:"major"`
	Minor int `js:"minor"`
	Patch int `js:"patch"`
}
