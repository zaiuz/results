package results

import "fmt"
import z "github.com/zaiuz/zaiuz"

type redirectResult struct {
	Permanent bool
	Url       string
}

func Redirect(permanent bool, url string, args ...interface{}) z.Result {
	// TODO: properly url-escape and query-escape formatted things.
	url = fmt.Sprintf(url, args...)
	return &redirectResult{permanent, url}
}

func (result *redirectResult) Render(c *z.Context) error {
	target, e := c.Request.URL.Parse(result.Url)
	if e != nil {
		return e
	}

	code := 302
	if result.Permanent {
		code = 301
	}

	header := c.ResponseWriter.Header()
	header["Location"] = []string{target.String()}
	c.ResponseWriter.WriteHeader(code)
	return nil
}
