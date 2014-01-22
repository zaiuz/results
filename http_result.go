package results

import "net/http"
import z "github.com/zaiuz/zaiuz"

type HttpResult struct {
	Code    int
	Headers http.Header
}

func NewHttpResult(code int, header string, values ...string) z.Result {
	headers := make(http.Header)
	if header != "" {
		headers[header] = values
	}

	return &HttpResult{code, headers}
}

func (r *HttpResult) Render(c *z.Context) error {
	header := c.ResponseWriter.Header()
	for k, v := range r.Headers {
		header[k] = v
	}

	c.ResponseWriter.WriteHeader(r.Code)
	return nil
}
