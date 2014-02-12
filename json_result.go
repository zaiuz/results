package results

import "encoding/json"
import z "github.com/zaiuz/zaiuz"

type jsonResult struct {
	*httpResult
	Object interface{}
}

func Json(code int, object interface{}) z.Result {
	http := Http(code, "Content-Type", "application/json").(*httpResult)
	return &jsonResult{http, object}
}

func (r *jsonResult) Render(c *z.Context) error {
	e := r.httpResult.Render(c)
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
