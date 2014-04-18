package results

import v "github.com/zaiuz/views"
import z "github.com/zaiuz/zaiuz"

// Vanilla mixin for convenience to add to your controller to get simpler access to result
// building methods.
type Controller struct{}

func (c *Controller) RenderFunc(render func(c *z.Context) error) z.Result {
	return RenderFunc(render)
}

func (c *Controller) Defer(resolve z.Action) z.Result {
	return Defer(resolve)
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

func (c *Controller) Redirect(url string, args ...interface{}) z.Result {
	return Redirect(url, args...)
}

func (c *Controller) PermanentRedirect(url string, args ...interface{}) z.Result {
	return PermanentRedirect(url, args...)
}

func (c *Controller) String(code int, str string, args ...interface{}) z.Result {
	return String(code, str, args...)
}

func (c *Controller) View(view v.View, data interface{}) z.Result {
	return View(view, data)
}

func (c *Controller) Resource() *ResourceRepresentable {
	return Resource()
}
