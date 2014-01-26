package results

import v "github.com/zaiuz/views"
import z "github.com/zaiuz/zaiuz"

type ViewResult struct {
	view v.View
	data interface{}
}

func NewViewResult(view v.View, data interface{}) z.Result {
	return &ViewResult{view, data}
}

func (r *ViewResult) Render(context *z.Context) error {
	return r.view.Render(context, r.data)
}
