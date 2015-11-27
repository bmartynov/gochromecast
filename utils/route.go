package utils
import (
	"io"
	"reflect"
	"net/http"
	"github.com/zenazn/goji/web"
	"github.com/bmartynov/gochromecast/responses"
)


func Route(controller interface{}, route string) interface{} {
	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {
		c.Env["Content-Type"] = "application/json"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		methodValue := reflect.ValueOf(controller).MethodByName(route)

		methodInterface := methodValue.Interface()

		method := methodInterface.(func(c *web.C, r *http.Request) (interface{}, error))

		body, err := method(&c, r)

		response := responses.JsonResponse{}

		response.Method = route

		response.Set(body, err)

		json_response := RenderResponse(response)

		io.WriteString(w, json_response)
	}
	return fn
}


func BuildMux(controller interface{}, routes map[string][]string) *web.Mux {
	downloader_mux := web.New()

	for c_method, route  := range routes {
		switch route[0] {
		case "GET":
			downloader_mux.Get(route[1], Route(controller, c_method))
		case "POST":
			downloader_mux.Post(route[1], Route(controller, c_method))
		}
	}
	return downloader_mux
}