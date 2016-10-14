package others

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
)

func NewCookie() Cookie {
	obj := js.Global.Get("Object").New()
	return Cookie{
		obj: obj,
	}
}

type Cookie struct {
	obj *js.Object

	Domain   string `js:"domain"`
	Path     string `js:"path"`
	Expiry   int64  `js:"expiry"`
	HttpOnly bool   `js:"httponly"`
	Secure   bool   `js:"secure"`
	Name     string `js:"name"`
	Value    string `js:"value"`
}

func (this *Cookie) SetExpiry(expiry time.Time) {
	this.obj.Set("expiry", expiry.Unix()*1000)
}

func (this *Cookie) Expires() string {
	obj := this.obj.Get("expires")
	if obj == js.Undefined {
		return ""
	}
	return obj.String()
}
