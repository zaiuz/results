package results

import "fmt"
import z "github.com/zaiuz/zaiuz"

type redirectResult struct {
	Permanent bool
	Url       string
}

// TODO: Properly url-escape and query-escape formatted things.
//   Otherwise doc the assumption that the url is safe (since this is mostly used
//   to format model ids into the string anyway.)
func Redirect(url string, args ...interface{}) z.Result {
	return &redirectResult{false, fmt.Sprintf(url, args...)}
}

func PermanentRedirect(url string, args ...interface{}) z.Result {
	return &redirectResult{true, fmt.Sprintf(url, args...)}
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
