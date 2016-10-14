package system

import (
	"github.com/gopherjs/gopherjs/js"
)

type OS struct {
	obj *js.Object

	Architecture string `js:"architecture"`
	Name         string `js:"name"`
	Release      string `js:"release"`
	Version      string `js:"version"`
}
