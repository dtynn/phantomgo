package webpage

import (
	"time"

	"github.com/dtynn/phantomgo/core/utils"
	"github.com/gopherjs/gopherjs/js"
)

type ResourceError struct {
	obj *js.Object

	Id          int    `js:"id"`
	Url         string `js:"url"`
	ErrorCode   int    `js:"errorCode"`
	ErrorString string `js:"errorString"`
}

type ResourceResponse struct {
	obj *js.Object

	Id          int               `js:"id"`
	Url         string            `js:"url"`
	Headers     map[string]string `js:"headers"`
	BodySize    int               `js:"bodySize"`
	ContentType string            `js:"contentType"`
	RedirectUrl string            `js:"redirectURL"`
	Stage       string            `js:"stage"`
	Status      int               `js:"status"`
	StatusText  string            `js:"statusText"`
}

func (this *ResourceResponse) Time() *time.Time {
	return utils.DateObjectToTime(this.obj.Get("time"))
}

type ResourceRequest struct {
	obj *js.Object

	Id      int               `js:"id"`
	Method  string            `js:"method"`
	Url     string            `js:"url"`
	Headers map[string]string `js:"headers"`
}

func (this *ResourceRequest) Time() *time.Time {
	return utils.DateObjectToTime(this.obj.Get("time"))
}

type ResourceNetworkRequest struct {
	obj *js.Object
}

func (this *ResourceNetworkRequest) Abort() {
	this.obj.Call("abort")
}

func (this *ResourceNetworkRequest) ChangeUrl(url string) {
	this.obj.Call("changeUrl", url)
}

func (this *ResourceNetworkRequest) SetHeader(key, val string) {
	this.obj.Call("setHeader", key, val)
}

type TimeoutRequest struct {
	obj *js.Object

	Id          int               `js:"id"`
	Method      string            `js:"method"`
	Url         string            `js:"url"`
	Headers     map[string]string `js:"headers"`
	ErrorCode   int               `js:"errorCode"`
	ErrorString string            `js:"errorString"`
}

func (this *TimeoutRequest) Time() *time.Time {
	return utils.DateObjectToTime(this.obj.Get("time"))
}
