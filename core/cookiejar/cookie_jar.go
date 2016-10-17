package cookiejar

import (
	"github.com/gopherjs/gopherjs/js"
)

func ObjectToCookieJar(obj *js.Object) *CookieJar {
	return &CookieJar{
		obj: obj,
	}
}

type CookieJar struct {
	obj *js.Object
}
