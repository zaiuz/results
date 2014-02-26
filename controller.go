package results

import v "github.com/zaiuz/views"
import z "github.com/zaiuz/zaiuz"

// Vanilla mixin for convenience to add to your controller to get simpler access to result
// building methods.
type Controller struct{}

func (c *Controller) Func(render func(c *z.Context) error) z.Result {
	return Func(render)
}

func (c *Controller) Dud() z.Result {
	return Dud()
}

func (c *Controller) Http(code int, header string, values ...string) z.Result {
	return Http(code, header, values...)
}

func (c *Controller) Json(code int, object interface{}) z.Result {
	return Json(code, object)
}

func (c *Controller) Redirect(permanent bool, url string, args ...interface{}) z.Result {
	return Redirect(permanent, url, args...)
}

func (c *Controller) String(code int, str string) z.Result {
	return String(code, str)
}

func (c *Controller) View(view v.View, data interface{}) z.Result {
	return View(view, data)
}
