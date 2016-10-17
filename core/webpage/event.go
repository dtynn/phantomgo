package webpage

import (
	"github.com/gopherjs/gopherjs/js"
)

type mouseEventType string

const (
	MouseEventMouseUp          mouseEventType = "mouseup"
	MouseEventMouseDown                       = "mousedown"
	MouseEventMouseMove                       = "mousemove"
	MouseEventMouseDoubleClick                = "mousedoubleclick"
	MouseEventClick                           = "click"
	MouseEventDoubleClick                     = "doubleclick"
)

type mouseButton string

const (
	MouseButtonLeft   mouseButton = "left"
	MouseButtonMiddle             = "middle"
	MouseButtonRight              = "right"
)

type keyboardEventType string

const (
	KeyboardEventKeyDown  keyboardEventType = "keydown"
	KeyboardEventKeyUp                      = "keyup"
	KeyboardEventKeyPress                   = "keypress"
)

type MousePosition struct {
	X int
	Y int
}

type Event struct {
	obj *js.Object

	Key      map[string]int `js:"key"`
	Modifier map[string]int `js:"modifier"`
}
