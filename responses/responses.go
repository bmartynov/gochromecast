package responses


type ServeResponse bool

type ServeStop bool

type DownloaderAddResponse struct {
	DownloadId string `json:"download_id"`
}

type DownloaderRemoveResponse struct {}

type DownloaderStartResponse struct {}

type DownloaderStopResponse struct {}

type DownloaderInfoFile struct {
	Length int64 `json:"length"`
	Path   string `json:"name"`
}

type DownloaderInfoResponse struct {
	DownloadId string `json:"download_id"`
	Length     int64 `json:"length"`
	Progress   float32 `json:"progress"`
	Name       string `json:"name"`
	Files      []DownloaderInfoFile `json:"files"`
}

type DownloaderPlayResponse struct {
	DownloadId string `json:"download_id"`
	Length     int `json:"length"`
	Name       string `json:"name"`
	Files      []DownloaderInfoFile `json:"files"`
}