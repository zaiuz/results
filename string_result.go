package results

import z "github.com/zaiuz/zaiuz"

type stringResult struct {
	*httpResult
	str string
}

func String(code int, str string) z.Result {
	http := Http(code, "Content-Type", "text/plain").(*httpResult)
	return &stringResult{http, str}
}

func (r *stringResult) Render(c *z.Context) error {
	e := r.httpResult.Render(c)
	if e != nil {
		return e
	}

	raw := []byte(r.str)
	_, e = c.ResponseWriter.Write(raw)
	return e
}
