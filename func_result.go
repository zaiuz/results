package results

import z "github.com/zaiuz/zaiuz"

type funcResult struct{
	render func(c *z.Context) error
}

func (result *funcResult) Render(c *z.Context) error {
	return result.render(c)
}

func ResultFunc(render func(c *z.Context) error) z.Result {
	return &funcResult{render}
}

func DudResult() z.Result {
	return ResultFunc(func(c *z.Context) error {
		return nil // no-op
	})
}
