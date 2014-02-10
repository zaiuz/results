package results

import z "github.com/zaiuz/zaiuz"

type StringResult struct {
	*HttpResult
	String string
}

func NewStringResult(code int, str string) z.Result {
	http := NewHttpResult(code, "Content-Type", "text/plain").(*HttpResult)
	return &StringResult{http, str}
}

func (r *StringResult) Render(c *z.Context) error {
	e := r.HttpResult.Render(c)
	if e != nil {
		return e
	}

	raw := []byte(r.String)
	_, e = c.ResponseWriter.Write(raw)
	return e
}
