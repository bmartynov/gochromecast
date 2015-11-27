package errors
import (
	"fmt"
	"strings"
)


type GoChromeCastError struct {
	Code    int `json:"code,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

func (this GoChromeCastError) Error() string {
	return fmt.Sprintf("GoChromeCastError(%d): %s", this.Code, this.Message)
}

func New(code int, message string, attrs... interface{}) GoChromeCastError {
	var err_attrs []string
	if len(attrs) > 0 {
		for _, attr := range attrs { err_attrs = append(err_attrs, fmt.Sprint(attr)) }
		return GoChromeCastError{
			Code:code,
			Message:fmt.Sprintf(message, strings.Join(err_attrs, ", ")),
		}
	} else {
		return GoChromeCastError{
			Code:code,
			Message:message,
		}
	}
}