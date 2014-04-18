package results

import z "github.com/zaiuz/zaiuz"

type deferResult struct {
	resolve z.Action
}

func (result *deferResult) Render(c *z.Context) error {
	r := result.resolve(c)
	if r == nil {
		return nil
	}

	return r.Render(c)
}

func Defer(resolve z.Action) z.Result {
	return &deferResult{resolve}
}
