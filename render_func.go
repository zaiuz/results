package results

import z "github.com/zaiuz/zaiuz"

type funcResult struct {
	render func(c *z.Context) error
}

func (result *funcResult) Render(c *z.Context) error {
	return result.render(c)
}

func RenderFunc(render func(c *z.Context) error) z.Result {
	return &funcResult{render}
}

func Dud() z.Result {
	return RenderFunc(func(c *z.Context) error {
		return nil // no-op
	})
}
