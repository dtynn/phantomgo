package handlers

import (
	"github.com/gopherjs/gopherjs/js"
)

type ErrorStack struct {
	*js.Object

	File     string `js:"file"`
	Line     int    `js:"line"`
	Function string `js:"function"`
}

type OnErrorHandler func(msg string, stack []*ErrorStack)
