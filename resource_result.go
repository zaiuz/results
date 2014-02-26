package results

import z "github.com/zaiuz/zaiuz"

type ResourceRepresentable struct {
	representations map[string]z.Result
}

func Resource() *ResourceRepresentable {
	reps := map[string]z.Result{}
	return &ResourceRepresentable{reps}
}

func (b *ResourceRepresentable) Json(json z.Result) *ResourceRepresentable {
	return b.Mime("application/json", json)
}

func (b *ResourceRepresentable) Text(text z.Result) *ResourceRepresentable {
	return b.Mime("text/plain", text)
}

func (b *ResourceRepresentable) Html(html z.Result) *ResourceRepresentable {
	return b.Mime("text/html", html)
}

func (b *ResourceRepresentable) Mime(mime string, r z.Result) *ResourceRepresentable {
	b.representations[mime] = r
	return b
}

func (b *ResourceRepresentable) Render(c *z.Context) error {
	accepts := c.Request.Header["Accept"]
	for _, mimetype := range accepts {
		if result, ok := b.representations[mimetype]; ok {
			return result.Render(c)
		}
	}

	// TODO: Probably better to error?
	return nil
}
