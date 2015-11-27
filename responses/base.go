package responses
import (
	"github.com/bmartynov/gochromecast/errors"
)

type JsonResponse struct {
	Method  string `json:"method"`
	Payload interface{}  `json:"result,omitempty"`
	Status  int          `json:"status"`
	Error   error        `json:"message,omitempty"`
}

//first argument - response, can be nil
//second argument - error, can be nil
func (this *JsonResponse) Set(response interface{}, err interface{}) {
	if err != nil {
		this.Status = 1
		switch err.(type) {

		case string:
			this.Error = errors.New(
				errors.ERROR_CODE,
				errors.ERROR_MESSAGE,
				err,
			)
		case error:
			this.Error = err.(error)
		}
	} else {
		this.Status = 0
		this.Payload = response
	}

}
