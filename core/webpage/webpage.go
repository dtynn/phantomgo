package webpage

import (
	"github.com/dtynn/phantomgo/core/cookiejar"
	"github.com/gopherjs/gopherjs/js"
)

func NewWebPage() *WebPage {
	obj := js.Global.Call("require", "webpage").Call("create")

	return &WebPage{
		obj: obj,
	}
}

type WebPage struct {
	obj *js.Object

	CanGoBack           bool               `js:"canGoBack"`
	CanGoForward        bool               `js:"canGoForward"`
	ClipRect            *ClipRect          `js:"clipRect"`
	Content             string             `js:"content"`
	Cookies             []cookiejar.Cookie `js:"cookies"`
	CustomHeaders       map[string]string  `js:"customHeaders"`
	Event               *Event             `js:"event"`
	FocusedFrameName    string             `js:"focusedFrameName"`
	FrameContent        string             `js:"frameContent"`
	FrameName           string             `js:"frameName"`
	FramePlainText      string             `js:"framePlainText"`
	FrameTitle          string             `js:"frameTitle"`
	FrameUrl            string             `js:"frameUrl"`
	FramesCount         int                `js:"framesCount"`
	FramesName          []string           `js:"framesName"`
	LibraryPath         string             `js:"libraryPath"`
	Loading             bool               `js:"loading"`
	LoadingProgress     int                `js:"loadingProgress"`
	NavigationLocked    bool               `js:"navigationLocked"`
	OfflineStoragePath  string             `js:"offlineStoragePath"`
	OfflineStorageQuota int                `js:"offlineStorageQuota"`
	OwnsPages           bool               `js:"ownsPages"`
	PagesWindowName     []string           `js:"pagesWindowName"`
	Pages               []WebPage          `js:"pages"`
	PaperSize           *PaperSize         `js:"paperSize"`
	PlainText           string             `js:"plainText"`
	ScrollPosition      *ScrollPosition    `js:"scrollPosition"`
	Settings            *Settings          `js:"settings"`
	Title               string             `js:"title"`
	Url                 string             `js:"url"`
	ViewportSize        *Size              `js:"viewportSize"`
	WindowName          string             `js:"windowName"`
	ZoomFactor          float64            `js:"zoomFactor"`
}
