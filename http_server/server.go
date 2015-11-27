package http_server

import (
	"io"
	"net/http"
	"github.com/zenazn/goji/web"
	"github.com/bmartynov/gochromecast/utils"
	"github.com/bmartynov/gochromecast/responses"
	"github.com/bmartynov/gochromecast/service"
	"time"
)

type GoChromeCastHttpServer struct {
	ds *service.DownloaderService
}

func (this *GoChromeCastHttpServer) DownloadAdd(c *web.C, r *http.Request) (interface{}, error) {
	var response responses.DownloaderAddResponse

	if params, err := utils.GetParams(r, "uri"); err != nil {
		return response, err

	} else {
		id, err := this.ds.DownloadAdd(params["uri"])
		response.DownloadId = id

		if err != nil { return response, err }
		return response, nil
	}
}

func (this *GoChromeCastHttpServer) DownloadRemove(c *web.C, r *http.Request) (interface{}, error) {
	var response responses.DownloaderRemoveResponse

	if params, err := utils.GetParams(r, "id"); err != nil {
		return response, err
	} else {
		err := this.ds.DownloadRemove(params["id"])
		if err != nil { return response, err }

		return response, nil
	}
}

func (this *GoChromeCastHttpServer) DownloadStart(c *web.C, r *http.Request) (interface{}, error) {
	var response responses.DownloaderStartResponse

	if params, err := utils.GetParams(r, "id"); err != nil {
		return response, err
	} else {
		err := this.ds.DownloadStart(params["id"])
		if err != nil { return response, err }

		return response, nil
	}
}

func (this *GoChromeCastHttpServer) DownloadStop(c *web.C, r *http.Request) (interface{}, error) {
	var response responses.DownloaderStopResponse

	if params, err := utils.GetParams(r, "id"); err != nil {
		return response, err
	} else {
		err := this.ds.DownloadStop(params["id"])
		if err != nil { return response, err }

		return response, nil
	}
}

func (this *GoChromeCastHttpServer) DownloadInfo(c *web.C, r *http.Request) (interface{}, error) {
	if params, err := utils.GetParams(r, "id"); err != nil {
		return struct{}{}, err
	} else {
		return this.ds.DownloadInfo(params["id"])
	}
}

func (this *GoChromeCastHttpServer) DownloadInfoAll(c *web.C, r *http.Request) (interface{}, error) {
	return this.ds.DownloadInfoAll(), nil
}

func (this *GoChromeCastHttpServer) DownloadPlay(c web.C, w http.ResponseWriter, r *http.Request) {
	if params, err := utils.GetParams(r, "id", "file_path"); err != nil {
		w.Header().Set("Content-Type", "application/json")
		r := responses.JsonResponse{}
		r.Method = "DownloadPlay"
		r.Set(nil, err)
		io.WriteString(w, utils.RenderResponse(r))
	} else {
		if sc, display_path, err := this.ds.DownloadPlay(params["id"], params["file_path"]); err != nil {
			w.Header().Set("Content-Type", "application/json")
			r := responses.JsonResponse{}
			r.Method = "DownloadPlay"
			r.Set(nil, err)
			io.WriteString(w, utils.RenderResponse(r))
		} else {
			defer func() {
				if err == nil { sc.Close() }
			}()
			w.Header().Set("Content-Disposition", "attachment; filename=\"" + params["file_path"] + "\"")
			http.ServeContent(w, r, display_path, time.Now(), sc)
		}
	}
}

func New() *web.Mux {
	gchs := GoChromeCastHttpServer{}

	gchs.ds = service.GetInstance()


	var routes = map[string][]string{
		"DownloadAdd": []string{"GET", "/download/add"},
		"DownloadStart": []string{"GET", "/download/start"},
		"DownloadStop": []string{"GET", "/download/stop"},
		"DownloadRemove": []string{"GET", "/download/remove"},
		"DownloadInfoAll": []string{"GET", "/download/info"},
	}

	downloader_mux := utils.BuildMux(&gchs, routes)
	downloader_mux.Get("/download/play", gchs.DownloadPlay)
	return downloader_mux
}
