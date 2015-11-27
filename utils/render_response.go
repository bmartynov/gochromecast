package utils
import (
	"encoding/json"
)


func RenderResponse(response interface{}) string {

	b, _ := json.Marshal(response)

	return string(b)
}
