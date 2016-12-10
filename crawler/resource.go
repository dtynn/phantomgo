package crawler

import (
	"fmt"
	"sync"
	"time"

	"github.com/dtynn/phantomgo/core/webpage"
)

type ResourceRequest struct {
	headers webpage.Headers

	RawUrl string
	Method string
}

type ResourceResponse struct {
	headers webpage.Headers

	ContentType   string
	ContentLength int
	RedirectUrl   string
	StatusCode    int
	StatusText    string
}

type ResourceError struct {
	ErrCode int
	ErrText string
}

func (this *ResourceError) Error() string {
	return fmt.Sprintf("[RESOURCE ERROR] %d | %q", this.ErrCode, this.ErrText)
}

type Resource struct {
	Id int

	// request infos
	RawUrl         string
	Method         string
	RequestHeaders webpage.Headers

	// response & errors
	Response *ResourceResponse

	Timeout *ResourceError
	Error   *ResourceError

	Finished bool

	RequestedAt time.Time
	FinishedAt  time.Time
}

func (this *Resource) Err() error {
	if this.Timeout != nil {
		return this.Timeout
	}

	if this.Error != nil {
		return this.Error
	}

	return nil
}

func NewResourceManager(page *webpage.WebPage) *ResourceManager {
	mgr := &ResourceManager{
		page:       page,
		resoureMap: map[int]*Resource{},
	}

	page.OnResourceRequested(mgr.OnResourceRequested)
	page.OnResourceReceived(mgr.OnResourceReceived)
	page.OnResourceTimeout(mgr.OnResourceTimeout)
	page.OnResourceError(mgr.OnResourceError)

	return mgr
}

type ResourceManager struct {
	page *webpage.WebPage

	resoureMap map[int]*Resource

	mutex sync.RWMutex
	wg    sync.WaitGroup
}

func (this *ResourceManager) OnResourceRequested(request *webpage.ResourceRequest, netword *webpage.ResourceNetworkRequest) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// resource exists
	if _, ok := this.resoureMap[request.Id]; ok {
		return
	}

	this.wg.Add(1)

	res := &Resource{
		Id: request.Id,
	}

	res.RawUrl = request.Url
	res.Method = request.Method
	res.RequestHeaders = request.Headers

	res.RequestedAt = time.Now()

	this.resoureMap[res.Id] = res
}

func (this *ResourceManager) OnResourceReceived(resp *webpage.ResourceResponse) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	res, ok := this.resoureMap[resp.Id]
	if !ok {
		return
	}

	switch resp.Stage {
	case webpage.ResourceResponseStageStart:

		res.Response = &ResourceResponse{
			headers:       resp.Headers,
			StatusCode:    resp.Status,
			StatusText:    resp.StatusText,
			ContentType:   resp.ContentType,
			ContentLength: resp.BodySize,
			RedirectUrl:   resp.RedirectUrl,
		}

	case webpage.ResourceResponseStageEnd:
		res.Finished = true
		res.FinishedAt = time.Now()

		this.wg.Done()
	}
}

func (this *ResourceManager) OnResourceError(err *webpage.ResourceError) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	res, ok := this.resoureMap[err.Id]
	if !ok {
		return
	}

	res.Error = &ResourceError{
		ErrCode: err.ErrorCode,
		ErrText: err.ErrorString,
	}
}

func (this *ResourceManager) OnResourceTimeout(request *webpage.TimeoutRequest) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	res, ok := this.resoureMap[request.Id]
	if !ok {
		return
	}

	res.Timeout = &ResourceError{
		ErrCode: request.ErrorCode,
		ErrText: request.ErrorString,
	}
}

func (this *ResourceManager) Len() int {
	return len(this.resoureMap)
}

func (this *ResourceManager) ResourceById(id int) *Resource {
	this.mutex.RLock()
	res := this.resoureMap[id]
	this.mutex.RUnlock()

	return res
}

func (this *ResourceManager) Wait() {
	this.wg.Wait()
}
