package service


import (
	"log"
	"sync"
	"github.com/anacrolix/torrent"
	"github.com/bmartynov/gochromecast/downloader"
	"github.com/bmartynov/gochromecast/downloader_factory"
	"github.com/bmartynov/gochromecast/responses"
)

var (
	dsMutex = sync.RWMutex{}
	Ds = New()
)

type DownloaderService struct {
	sync.RWMutex
	df *downloader_factory.DownloaderFactory
}

func (this *DownloaderService) DownloadAdd(uri string) (id string, err error) {
	log.Println("DownloadAdd: uri -> ", uri)
	id, err = this.df.Add(uri)
	if err != nil { return id, err }
	return id, nil
}

func (this *DownloaderService) DownloadRemove(downloader_id string) error {
	log.Println("DownloadRemove: downloader_id -> ", downloader_id)
	return this.df.Remove(downloader_id)
}

func (this *DownloaderService) DownloadStart(downloader_id string) error {
	log.Println("DownloadStart: downloader_id -> ", downloader_id)
	return this.df.Start(downloader_id)
}

func (this *DownloaderService) DownloadStop(downloader_id string) error {
	log.Println("DownloadStop: downloader_id -> ", downloader_id)
	return this.df.Stop(downloader_id)
}

func (this *DownloaderService) DownloadStopAll() error {
	log.Println("DownloadStopAll")
	this.df.StopAll()
	return nil
}

func (this *DownloaderService) DownloadInfo(downloader_id string) (dif responses.DownloaderInfoResponse, err error) {
	log.Println("DownloadInfo: downloader_id -> ", downloader_id)
	var d = &downloader.Downloader{}
	d, err = this.df.Get(downloader_id)
	if err != nil { return }
	return d.Info(), nil
}

func (this *DownloaderService) DownloadInfoAll() (difa []responses.DownloaderInfoResponse) {
	log.Println("DownloadInfoAll")
	var ds []*downloader.Downloader
	ds = this.df.GetAll()
	for _, d := range ds {
		difa = append(difa, d.Info())
	}
	return difa
}

func (this *DownloaderService) DownloadPlay(downloader_id, file_path string) (sc SeekAbleContent, path string, err error) {
	log.Println("DownloadPlay: downloader_id -> ", downloader_id, " file_path -> ", file_path)
	this.RLock()
	defer this.RUnlock()

	var downloader = &downloader.Downloader{}
	var file torrent.File

	downloader, err = this.df.Get(downloader_id)
	if err != nil { return }

	file, err = downloader.GetFileFoPlay(file_path)
	if err != nil { return }

	sc, err = NewReader(downloader.Torrent, file)
	if err != nil { return }
	return sc, file.DisplayPath(), nil
}


func New() *DownloaderService {
	var ds DownloaderService
	ds.df = downloader_factory.GetInstance()
	return &ds
}

func GetInstance() *DownloaderService {
	dsMutex.RLock()
	defer dsMutex.RUnlock()
	return Ds
}