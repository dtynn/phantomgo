package webpage

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
)

type Settings struct {
	obj *js.Object

	JavascriptEnabled             bool   `js:"javascriptEnabled"`
	JavascriptCanCloseWindows     bool   `js:"javascriptCanCloseWindows"`
	JavascriptCanOpenWindows      bool   `js:"javascriptCanOpenWindows"`
	LoadImages                    bool   `js:"loadImages"`
	LocalToRemoteUrlAccessEnabled bool   `js:"localToRemoteUrlAccessEnabled"`
	UserAgent                     string `js:"userAgent"`
	UserName                      string `js:"userName"`
	Password                      string `js:"password"`
	XSSAuditingEnabled            bool   `js:"XSSAuditingEnabled"`
	WebSecurityEnabled            bool   `js:"webSecurityEnabled"`
	ResourceTimeout               int64  `js:"resourceTimeout"`
}

func (this *Settings) SetResourceTimeout(timeout time.Duration) {
	this.obj.Set("resourceTimeout", timeout/time.Millisecond)
}
