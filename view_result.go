package results

import v "github.com/zaiuz/views"
import z "github.com/zaiuz/zaiuz"

type viewResult struct {
	view v.View
	data interface{}
}

func View(view v.View, data interface{}) z.Result {
	return &viewResult{view, data}
}

func (r *viewResult) Render(context *z.Context) error {
	return r.view.Render(context, r.data)
}
