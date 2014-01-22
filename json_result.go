package results

import "encoding/json"
import z "github.com/zaiuz/zaiuz"

type JsonResult struct {
	*HttpResult
	Object interface{}
}

func NewJsonResult(code int, object interface{}) z.Result {
	http := NewHttpResult(code, "Content-Type", "application/json").(*HttpResult)
	return &JsonResult{http, object}
}

func (r *JsonResult) Render(c *z.Context) error {
	e := r.HttpResult.Render(c)
	if e != nil {
		return e
	}

	raw, e := json.Marshal(r.Object)
	if e != nil {
		return e
	}

	_, e = c.ResponseWriter.Write(raw)
	return e
}
