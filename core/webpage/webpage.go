package webpage

import (
	"time"

	"github.com/dtynn/phantomgo/core/cookiejar"
	"github.com/dtynn/phantomgo/core/utils"
	"github.com/gopherjs/gopherjs/js"
)

func ObjectToWebPage(obj *js.Object) *WebPage {
	return &WebPage{
		obj: obj,
	}
}

func NewWebPage() *WebPage {
	return ObjectToWebPage(js.Global.Call("require", "webpage").Call("create"))
}

type WebPage struct {
	obj *js.Object

	CanGoBack           bool                 `js:"canGoBack"`
	CanGoForward        bool                 `js:"canGoForward"`
	ClipRect            *ClipRect            `js:"clipRect"`
	Content             string               `js:"content"`
	Cookies             []cookiejar.Cookie   `js:"cookies"`
	CookieJar           *cookiejar.CookieJar `js:"cookieJar"`
	CustomHeaders       map[string]string    `js:"customHeaders"`
	Event               *Event               `js:"event"`
	FocusedFrameName    string               `js:"focusedFrameName"`
	FrameContent        string               `js:"frameContent"`
	FrameName           string               `js:"frameName"`
	FramePlainText      string               `js:"framePlainText"`
	FrameTitle          string               `js:"frameTitle"`
	FrameUrl            string               `js:"frameUrl"`
	FramesCount         int                  `js:"framesCount"`
	FramesName          []string             `js:"framesName"`
	LibraryPath         string               `js:"libraryPath"`
	Loading             bool                 `js:"loading"`
	LoadingProgress     int                  `js:"loadingProgress"`
	NavigationLocked    bool                 `js:"navigationLocked"`
	OfflineStoragePath  string               `js:"offlineStoragePath"`
	OfflineStorageQuota int                  `js:"offlineStorageQuota"`
	OwnsPages           bool                 `js:"ownsPages"`
	PagesWindowName     []string             `js:"pagesWindowName"`
	Pages               []WebPage            `js:"pages"`
	PaperSize           *PaperSize           `js:"paperSize"`
	PlainText           string               `js:"plainText"`
	ScrollPosition      *ScrollPosition      `js:"scrollPosition"`
	Settings            *PageSettings        `js:"settings"`
	Title               string               `js:"title"`
	Url                 string               `js:"url"`
	ViewportSize        *Size                `js:"viewportSize"`
	WindowName          string               `js:"windowName"`
	ZoomFactor          float64              `js:"zoomFactor"`
}

// cookies
func (this *WebPage) name() {

}

func (this *WebPage) AddCookie(cookie cookiejar.Cookie) bool {
	return this.obj.Call("addCookie", cookie).Bool()
}

func (this *WebPage) ClearCookies() {
	this.obj.Call("clearCookies")
}

func (this *WebPage) DeleteCookie(name string) bool {
	return this.obj.Call("deleteCookie", name).Bool()
}

// frames & pages
func (this *WebPage) ChildFramesCount() int {
	return this.obj.Call("childFramesCount").Int()
}

func (this *WebPage) ChildFramesName() []string {
	return utils.ObjectToStringSlice(this.obj.Call("childFramesName"))
}

func (this *WebPage) CurrentFrameName() string {
	return this.obj.Call("currentFrameName").String()
}

func (this *WebPage) GetPage(windowName string) *WebPage {
	obj := this.obj.Call("getPage", windowName)
	return &WebPage{
		obj: obj,
	}
}

func (this *WebPage) SwitchToChildFrame(frameName string) bool {
	return this.obj.Call("switchToChildFrame", frameName).Bool()
}

func (this *WebPage) SwitchToFocusedFrame() {
	this.obj.Call("switchToFocusedFrame")
}

func (this *WebPage) SwitchToFrame(frameName string) bool {
	return this.obj.Call("switchToFrame", frameName).Bool()
}

func (this *WebPage) SwitchToMainFrame() {
	this.obj.Call("switchToMainFrame")
}

func (this *WebPage) SwitchToParentFrame() bool {
	return this.obj.Call("switchToParentFrame").Bool()
}

func (this *WebPage) SetContent(content, url string) {
	this.obj.Call("setContent", content, url)
}

func (this *WebPage) UploadFile(selector string, filenames []string) {
	this.obj.Call("uploadFile", selector, filenames)
}

// actions & navigations
func (this *WebPage) Close() {
	this.obj.Call("close")
}

func (this *WebPage) GoBack() bool {
	return this.obj.Call("goBack").Bool()
}

func (this *WebPage) GoForward() bool {
	return this.obj.Call("goForward").Bool()
}

func (this *WebPage) Go(historyRelativeIndex int) bool {
	return this.obj.Call("go", historyRelativeIndex).Bool()
}

func (this *WebPage) OpenUrl(url string, request *Request, onLoadFinished OnLoadFinished) {
	var req js.M
	if request != nil {
		req = js.M{}

		if request.Operation != "" {
			req["operation"] = request.Operation
		}

		if request.Encoding != "" {
			req["encoding"] = request.Encoding
		}

		if len(request.Headers) > 0 {
			req["headers"] = request.Headers
		}

		if len(request.Data) > 0 {
			req["data"] = string(request.Data)
		}
	}

	this.OnLoadFinished(onLoadFinished)
	this.obj.Call("openUrl", url, req, this.Settings)
}

func (this *WebPage) Open(url, method string, data []byte, headers map[string]string, onLoadFinished OnLoadFinished) {
	this.OpenUrl(url, &Request{
		Operation: method,
		Data:      data,
		Headers:   headers,
	}, onLoadFinished)
}

func (this *WebPage) Reload() {
	this.obj.Call("reload")
}

func (this *WebPage) Stop() {
	this.obj.Call("stop")
}

// evaluates
func (this *WebPage) EvaluateAsync(fn *js.Object, delay time.Duration, args ...interface{}) {
	params := make([]interface{}, 2, 2+len(args))
	params[0] = fn
	params[1] = delay / time.Millisecond
	params = append(params, args...)

	this.obj.Call("evaluateAsync", params...)
}

func (this *WebPage) EvaluateJavaScript(text string) *js.Object {
	return this.obj.Call("evaluateJavaScript", text)
}

func (this *WebPage) Evaluate(fn *js.Object, args ...interface{}) *js.Object {
	params := make([]interface{}, 1, 1+len(args))
	params[0] = fn
	params = append(params, args...)

	return this.obj.Call("evaluate", params...)
}

func (this *WebPage) StopJavaScript() {
	this.obj.Call("stopJavaScript")
}

func (this *WebPage) IncludeJs(url string, callback *js.Object) {
	this.obj.Call("includeJs", url, callback)
}

func (this *WebPage) InjectJs(filename string) bool {
	return this.obj.Call("injectJs", filename).Bool()
}

// events
func (this *WebPage) SendMouseEvent(typ mouseEventType, position *MousePosition, button *mouseButton) {
	args := make([]interface{}, 4)
	args[0] = typ
	if position != nil {
		args[1] = position.X
		args[2] = position.Y
	}

	if button != nil {
		args[3] = *button
	}

	this.obj.Call("sendEvent", args...)
}

func (this *WebPage) SendKeyEvent(typ keyboardEventType, key, modifier int) {
	this.obj.Call("sendEvent", typ, key, nil, nil, modifier)
}

func (this *WebPage) SendKeysEvent(typ keyboardEventType, keys string, modifier int) {
	this.obj.Call("sendEvent", typ, keys, nil, nil, modifier)
}

// render page
func (this *WebPage) RenderBase64(format renderFormat) string {
	return this.obj.Call("renderBase64", format).String()
}

func (this *WebPage) Render(filename string, options *RenderOptions) {
	var opt js.M
	if options != nil {
		opt = js.M{}

		if options.Format != "" {
			opt["format"] = options.Format
		}

		quality := options.Quality
		if quality < 0 {
			quality = 0
		} else if quality > 100 {
			quality = 100
		}

		opt["quality"] = quality
	}

	this.obj.Call("render", filename, opt)
}

// others
func (this *WebPage) SetProxy(proxyUrl string) {
	this.obj.Call("setProxy", proxyUrl)
}

// handlers
func (this *WebPage) OnAlert(fn OnAlert) {
	this.obj.Set("onAlert", fn)
}

func (this *WebPage) OnCallback(fn OnCallback) {
	this.obj.Set("onCallback", fn)
}

func (this *WebPage) OnClosing(fn OnClosing) {
	this.obj.Set("onClosing", fn)
}

func (this *WebPage) OnConfirm(fn OnConfirm) {
	this.obj.Set("onConfirm", fn)
}

func (this *WebPage) OnConsoleMessage(fn OnConsoleMessage) {
	this.obj.Set("onConsoleMessage", fn)
}

func (this *WebPage) OnError(fn OnError) {
	this.obj.Set("onError", fn)
}

func (this *WebPage) OnFilePicker(fn OnFilePicker) {
	this.obj.Set("onFilePicker", fn)
}

func (this *WebPage) OnInitialized(fn OnInitialized) {
	this.obj.Set("onInitialized", fn)
}

func (this *WebPage) OnLoadFinished(fn OnLoadFinished) {
	this.obj.Set("onLoadFinished", fn)
}

func (this *WebPage) OnLoadStarted(fn OnLoadStarted) {
	this.obj.Set("onLoadStarted", fn)
}

func (this *WebPage) OnNavigationRequested(fn OnNavigationRequested) {
	this.obj.Set("onNavigationRequested", fn)
}

func (this *WebPage) OnPageCreated(fn OnPageCreated) {
	this.obj.Set("onPageCreated", fn)
}

func (this *WebPage) OnPrompt(fn OnPrompt) {
	this.obj.Set("onPrompt", fn)
}

func (this *WebPage) OnResourceError(fn OnResourceError) {
	this.obj.Set("onResourceError", fn)
}

func (this *WebPage) OnResourceReceived(fn OnResourceReceived) {
	this.obj.Set("onResourceReceived", fn)
}

func (this *WebPage) OnResourceRequested(fn OnResourceRequested) {
	this.obj.Set("onResourceRequested", fn)
}

func (this *WebPage) OnResourceTimeout(fn OnResourceTimeout) {
	this.obj.Set("onResourceTimeout", fn)
}

func (this *WebPage) OnUrlChanged(fn OnUrlChanged) {
	this.obj.Set("onUrlChanged", fn)
}
