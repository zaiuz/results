package results

import z "github.com/zaiuz/zaiuz"

type RedirectResult struct {
	Permanent bool
	Url       string
}

func NewRedirectResult(permanent bool, url string) z.Result {
	return &RedirectResult{permanent, url}
}

func (result *RedirectResult) Render(c *z.Context) error {
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
