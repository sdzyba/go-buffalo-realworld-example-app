package errmap

import "encoding/json"

type Errs struct {
	Map map[string]interface{} `json:"errors"`
}

func NewErrs() *Errs {
	return &Errs{Map: map[string]interface{}{}}
}

func (e *Errs) Error() string {
	res, _ := json.Marshal(e)

	return string(res)
}
