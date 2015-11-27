package downloader_factory

import (
	"sync"
	"github.com/bmartynov/gochromecast/utils"
	"github.com/bmartynov/gochromecast/errors"
	"github.com/bmartynov/gochromecast/downloader"
	"fmt"
)

var (
	dfMutex = sync.RWMutex{}
	df = New()
)

type DownloaderFactory struct {
	sync.RWMutex
	downloaders map[string]*downloader.Downloader
}

//add downloader
func (this *DownloaderFactory) Add(uri string) (id string, err error) {
	this.Lock()
	defer this.Unlock()
	id = utils.CalcMd5(uri)
	if _, ok := this.downloaders[id]; ok {
		return id, errors.New(
			errors.DOWNLOADER_ALREADY_EXISTS_CODE,
			errors.DOWNLOADER_ALREADY_EXISTS_MESSAGE,
		)
	}

	downloader, id, err := downloader.New(uri)

	if err != nil { return id, err }

	err = downloader.TorrentAdd(uri)

	if err != nil { return id, err }
	this.downloaders[id] = downloader

	return id, nil
}
//stop download and remove downloader
func (this *DownloaderFactory) Remove(downloader_id string) error {
	this.Lock()
	defer this.Unlock()
	if downloader, ok := this.downloaders[downloader_id]; !ok {
		return errors.New(
			errors.DOWNLOADER_NOT_FOUND_CODE,
			errors.DOWNLOADER_NOT_FOUND_MESSAGE,
			downloader_id,
		)
	} else { downloader.Stop() }
	delete(this.downloaders, downloader_id)

	return nil
}
//start download by id
func (this *DownloaderFactory) Start(downloader_id string) error {
	this.Lock()
	this.Unlock()
	downloader, err := this.Get(downloader_id)

	fmt.Println(fmt.Sprintf("%+v", downloader), err)

	if err != nil { return err }

	downloader.Start()

	return nil
}
//start download by id
func (this *DownloaderFactory) Stop(downloader_id string) error {
	this.Lock()
	this.Unlock()
	downloader, err := this.Get(downloader_id)
	if err != nil { return err }

	downloader.Stop()

	return nil
}
//get downlod by id
func (this *DownloaderFactory) Get(downloader_id string) (*downloader.Downloader, error) {
	this.RLock()
	defer this.RUnlock()
	if d, ok := this.downloaders[downloader_id]; ok {
		return d, nil
	} else {
		return d, errors.New(
			errors.DOWNLOADER_NOT_FOUND_CODE,
			errors.DOWNLOADER_NOT_FOUND_MESSAGE,
			downloader_id,
		)
	}
}
//get all downloads
func (this *DownloaderFactory) GetAll() (downloaders []*downloader.Downloader) {
	this.RLock()
	defer this.RUnlock()
	for _, d := range this.downloaders {
		downloaders = append(downloaders, d)
	}
	return downloaders
}
//stop all downloads
func (this *DownloaderFactory) StopAll() {
	this.Lock()
	defer this.Unlock()
	for downloader_id, _ := range this.downloaders {this.downloaders[downloader_id].Stop()}
	for downloader_id, _  := range this.downloaders {delete(this.downloaders, downloader_id)}
}
//create new factory
func New() *DownloaderFactory {
	df := &DownloaderFactory{}
	df.downloaders = make(map[string]*downloader.Downloader)

	return df
}
//get download factory instance
func GetInstance() *DownloaderFactory {
	dfMutex.RLock()
	defer dfMutex.RUnlock()
	return df
}