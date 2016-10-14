package webpage

import (
	"github.com/gopherjs/gopherjs/js"
)

type Event struct {
	obj *js.Object

	Key      map[string]int `js:"key"`
	Modifier map[string]int `js:"modifier"`
}
