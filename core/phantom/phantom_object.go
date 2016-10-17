package phantom

import (
	"github.com/dtynn/phantomgo/core/cookiejar"
	"github.com/dtynn/phantomgo/core/filesystem"
	"github.com/dtynn/phantomgo/core/system"
	"github.com/dtynn/phantomgo/core/webpage"
	"github.com/dtynn/phantomgo/core/webserver"
	"github.com/gopherjs/gopherjs/js"
)

var phantomObject = &PhantomObject{
	obj: js.Global.Get("phantom"),
}

func GetPhantomObject() *PhantomObject {
	return phantomObject
}

type PhantomObject struct {
	obj *js.Object

	DefaultSettings *webpage.PageSettings `js:"defaultSettings"`
	LibraryPath     string                `js:"libraryPath"`
	OutputEncoding  string                `js:"outputEncoding"`
	Version         *Version              `js:"version"`
	Page            *webpage.WebPage      `js:"page"`
	CookiesEnabled  bool                  `js:"cookiesEnabled"`
	Cookies         []cookiejar.Cookie    `js:"cookies"`
	WebdriverMode   bool                  `js:"webdriverMode"`
	RemoteDebugPort bool                  `js:"remoteDebugPort"`
}

func (this *PhantomObject) CreateCookieJar(filename string) *cookiejar.CookieJar {
	return cookiejar.ObjectToCookieJar(this.obj.Call("createCookieJar", filename))
}

func (this *PhantomObject) CreateWebPage() *webpage.WebPage {
	return webpage.ObjectToWebPage(this.obj.Call("createWebPage"))
}

func (this *PhantomObject) CreateWebServer() *webserver.WebServer {
	return webserver.ObjectToWebServer(this.obj.Call("createWebServer"))
}

func (this *PhantomObject) CreateFilesystem() *filesystem.FileSystem {
	return filesystem.ObjectToFileSystem(this.obj.Call("createFilesystem"))
}

func (this *PhantomObject) CreateSystem() *system.System {
	return system.ObjectToSystem(this.obj.Call("createSystem"))
}

func (this *PhantomObject) LoadModule(moduleSource, filename string) {
	this.obj.Call("loadModule", moduleSource, filename)
}

func (this *PhantomObject) InjectJs(filepath string) bool {
	return this.obj.Call("injectJs", filepath).Bool()
}

func (this *PhantomObject) AddCookie(cookie cookiejar.Cookie) bool {
	return this.obj.Call("addCookie", cookie).Bool()
}

func (this *PhantomObject) DeleteCookie(cookieName string) bool {
	return this.obj.Call("deleteCookie", cookieName).Bool()
}

func (this *PhantomObject) ClearCookies() {
	this.obj.Call("clearCookies")
}

func (this *PhantomObject) Proxy() string {
	return this.obj.Call("proxy").String()
}

func (this *PhantomObject) SetProxy(host string, port int, proxyType, user, password string) {
	this.obj.Call("setProxy", host, port, proxyType, user, password)
}

func (this *PhantomObject) Exit(code int) {
	this.obj.Call("exit", code)
}

func (this *PhantomObject) DebugExit(code int) {
	this.obj.Call("debugExit", code)
}

func (this *PhantomObject) ResolveRelativeUrl(url, base string) string {
	return this.obj.Call("resolveRelativeUrl", url, base).String()
}

func (this *PhantomObject) FullyDecodeUrl(url string) string {
	return this.obj.Call("fullyDecodeUrl", url).String()
}

func (this *PhantomObject) OnError(fn webpage.OnError) {
	this.obj.Set("onError", fn)
}

func (this *PhantomObject) AboutToExit(fn func(code int)) {
	this.obj.Set("aboutToExit", fn)
}

func (this *PhantomObject) OnInitialized(fn func()) {
	this.obj.Set("onInitialized", fn)
}

func (this *PhantomObject) OnPrintConsoleMessage(fn func(msg string)) {
	this.obj.Set("printConsoleMessage", fn)
}
