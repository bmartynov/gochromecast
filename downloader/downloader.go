package downloader

import (
	"os"
	"path/filepath"
	"sync"
	"strings"
	"net/http"
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/bmartynov/gochromecast/errors"
	"github.com/bmartynov/gochromecast/responses"
	"github.com/bmartynov/gochromecast/utils"
	"fmt"
)

const (
	PERCENT_TO_READY = float32(5)
)

type DownloaderInfoFile struct {
	Length int64 `json:"length"`
	Path   string `json:"name"`
}

type DownloaderInfoResponse struct {
	DownloadId string `json:"stream_id"`
	Length     int64 `json:"length"`
	Progress   float32 `json:"progress"`
	Name       string `json:"name"`
	Files      []DownloaderInfoFile `json:"files"`
}

type DownloaderPlayResponse struct {
	DownloadId string `json:"stream_id"`
	Length     int `json:"length"`
	Name       string `json:"name"`
	Files      []DownloaderInfoFile `json:"files"`
}

//Downloader struct. holds info about client and torrent
type Downloader struct {
	sync.RWMutex
	Id       string
	Client   *torrent.Client
	Torrent  torrent.Torrent
	CurrProgress float32
	running  bool
}

func (this *Downloader) GetFileFoPlay(path string) (t_file torrent.File, err error) {
	this.RLock()
	defer this.RUnlock()

	for _, t_file = range this.Torrent.Files() {
		if t_file.Path() == path {
			return
		}
	}
	return t_file, errors.New(
		errors.DOWNLOADER_FILE_NOT_FOUND_CODE,
		errors.DOWNLOADER_FILE_NOT_FOUND_MESSAGE,
	)
}

func (this *Downloader) Progress() float32 {
	this.RLock()
	defer this.RUnlock()
	return float32(this.Torrent.BytesCompleted()) / float32(this.Torrent.Length()) * 100
}

func (this *Downloader) ReadyForPlay() bool {
	if this.Progress() >= PERCENT_TO_READY {
		return true
	}
	return false
}

func (this *Downloader) metaInfoFromUrl(url string) (metaInfo *metainfo.MetaInfo, err error) {
	var response = &http.Response{}
	defer response.Body.Close()

	response, err = http.Get(url)
	if err != nil { return }

	metaInfo, err = metainfo.Load(response.Body)
	if err != nil { return }

	return metaInfo, nil
}

func (this *Downloader) TorrentAdd(uri string) error {
	if strings.HasPrefix(uri, "magnet:") {
		t, err := this.torrentAddFromMagnet(uri)
		if err != nil { return err }

		this.Torrent = t

		return nil
	}
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		t, err := this.torrentAddFromUrl(uri)
		if err != nil { return err }

		this.Torrent = t

		return nil
	}
	return errors.New(
		errors.DOWNLOADER_INVALID_URI_CODE,
		errors.DOWNLOADER_INVALID_URI_MESSAGE,
		uri,
	)
}

func (this *Downloader) torrentAddFromMagnet(magnet string) (t torrent.Torrent, err error) {
	t, err = this.Client.AddMagnet(magnet)
	return
}

func (this *Downloader) torrentAddFromUrl(url string) (t torrent.Torrent, err error) {
	var metaInfo = &metainfo.MetaInfo{}

	metaInfo, err = this.metaInfoFromUrl(url)
	if err != nil { return  }

	t, _, err = this.Client.AddTorrentSpec(torrent.TorrentSpecFromMetaInfo(metaInfo))

	return
}

func (this *Downloader) Start() {
	fmt.Println("this.Torrent", fmt.Sprintf("%+v", this.Torrent))
	go func() {
		<-this.Torrent.GotInfo()

		this.Torrent.DownloadAll()

		for i := 0; i < len(this.Torrent.Pieces) / 100 * int(PERCENT_TO_READY); i++ {
			this.Torrent.Pieces[i].Priority = torrent.PiecePriorityReadahead
		}
	}()
	this.running = true
}

func (this *Downloader) Stop() {
	this.Torrent.Drop()
	this.Client.Close()
	this.running = false
}

func (this *Downloader) Info() (response responses.DownloaderInfoResponse) {
	this.RLock()
	defer this.RUnlock()
	response.DownloadId = this.Id
	response.Length = this.Torrent.Length()
	response.Progress = this.Progress()
	response.Name = this.Torrent.Name()
	for _, f := range this.Torrent.Files() {
		response.Files = append(response.Files, responses.DownloaderInfoFile{
			Length:f.Length(),
			Path: f.Path(),
		})
	}
	return response
}

func New(uri string) (*Downloader, string, error) {
	var err error

	var id string
	var client *torrent.Client
	var downloader = &Downloader{}

	id =  utils.CalcMd5(uri)

	downloader.Id = id

	client, err = torrent.NewClient(&torrent.Config{
		DataDir: filepath.Join(os.TempDir(), "gochromecast"),
		NoUpload: true,
		Seed: true,
	})

	if err != nil {
		return downloader, id,  errors.New(
			errors.DOWNLOADER_CANNOT_CREATE_TCLIENT_CODE,
			errors.DOWNLOADER_CANNOT_CREATE_TCLIENT_MESSAGE,
			err,
		)
	}
	downloader.Client = client
	return downloader, id, nil
}