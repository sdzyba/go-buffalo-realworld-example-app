package errutil

import "encoding/json"

type ErrorResponse struct {
	Errors map[string]interface{} `json:"errors"`
}

func (e *ErrorResponse) Error() string {
	res, _ := json.Marshal(e)
	return string(res)
}
