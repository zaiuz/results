package results

import "net/http"
import z "github.com/zaiuz/zaiuz"

type httpResult struct {
	code    int
	headers http.Header
}

func Http(code int, header string, values ...string) z.Result {
	headers := make(http.Header)
	if header != "" {
		headers[header] = values
	}

	return &httpResult{code, headers}
}

func (r *httpResult) Render(c *z.Context) error {
	header := c.ResponseWriter.Header()
	for k, v := range r.headers {
		header[k] = v
	}

	c.ResponseWriter.WriteHeader(r.code)
	return nil
}
