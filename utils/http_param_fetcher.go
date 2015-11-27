package utils

import (
	"strings"
	"net/http"
	"github.com/bmartynov/gochromecast/errors"
)

func GetParams(r *http.Request, params... string) (response map[string]string, err error) {
	response = make(map[string]string)
	query := r.URL.Query()
	var missed_params []string

	for _, param := range params {
		if v := query.Get(param); v == "" {
			missed_params = append(missed_params, param)
		} else {
			response[param] = v
		}
	}
	if len(missed_params) > 0 {
		err = errors.New(
			errors.HTTP_PARAM_MISSED_CODE,
			errors.HTTP_PARAM_MISSED_MESSAGE,
			strings.Join(missed_params, ", "),
		)
		return
	}
	return
}