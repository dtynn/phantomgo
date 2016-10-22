package utils

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
)

func ObjectToStringSlice(obj *js.Object) []string {
	if obj == nil || obj == js.Undefined {
		return nil
	}

	size := obj.Length()
	strs := make([]string, obj.Length())
	for i := 0; i < size; i++ {
		strs[i] = obj.Index(i).String()
	}

	return strs
}

func DateObjectToTime(obj *js.Object) *time.Time {
	if obj == nil || obj == js.Undefined {
		return nil
	}

	t := time.Unix(obj.Call("getTime").Int64()/1000, 0)

	return &t
}
