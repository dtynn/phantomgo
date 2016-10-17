package webpage

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	LoadStatusSuccess = "success"
	LoadStatusFail    = "fail"
)

type navigationType string

const (
	NavigationTypeUndefined       navigationType = "Undefined"
	NavigationTypeLinkClicked                    = "LinkClicked"
	NavigationTypeFormSubmitted                  = "FormSubmitted"
	NavigationTypeBackOrForward                  = "BackOrForward"
	NavigationTypeReload                         = "Reload"
	NavigationTypeFormResubmitted                = "FormResubmitted"
	NavigationTypeOther                          = "Other"
)

type ErrorStack struct {
	obj *js.Object

	File     string `js:"file"`
	Line     int    `js:"line"`
	Function string `js:"function"`
}

type OnAlert func(msg string)

type OnCallback func(data *js.Object)

type OnClosing func(page *WebPage)

type OnConfirm func(msg string) bool

type OnConsoleMessage func(msg string, lineNum string, sourceId string)

type OnError func(msg string, stack []ErrorStack)

type OnFilePicker func(oldFile string) string

type OnInitialized func()

type OnLoadFinished func(status string)

type OnLoadStarted func()

type OnNavigationRequested func(url string, typ navigationType, willNavigate bool, fromMainFrame bool)

type OnPageCreated func(newPage *WebPage)

type OnPrompt func(msg, defaultValue string) string

type OnResourceError func(err *ResourceError)

type OnResourceReceived func(resp *ResourceResponse)

type OnResourceRequested func(request *ResourceRequest, netword *ResourceNetworkRequest)

type OnResourceTimeout func(request *TimeoutRequest)

type OnUrlChanged func(targetUrl string)
