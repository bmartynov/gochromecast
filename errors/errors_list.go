package errors


const (
	//unhandled error
	ERROR_CODE = -1
	ERROR_MESSAGE = "error occured (%s)"

	//params missed
	HTTP_PARAM_MISSED_CODE = 1
	HTTP_PARAM_MISSED_MESSAGE = "required param missed (%s)"

	//downloader errors
	DOWNLOADER_NOT_FOUND_CODE = 2
	DOWNLOADER_NOT_FOUND_MESSAGE = "requested download not found (%s)"

	DOWNLOADER_FILE_NOT_FOUND_CODE = 3
	DOWNLOADER_FILE_NOT_FOUND_MESSAGE = "file not found (%s)"

	DOWNLOADER_CANNOT_CREATE_TCLIENT_CODE = 4
	DOWNLOADER_CANNOT_CREATE_TCLIENT_MESSAGE = "cannot create torrent client: %s"

	DOWNLOADER_ALREADY_EXISTS_CODE = 5
	DOWNLOADER_ALREADY_EXISTS_MESSAGE = "download already exists"

	DOWNLOADER_INVALID_URI_CODE = 6
	DOWNLOADER_INVALID_URI_MESSAGE = "invalid uri (%s)"

)