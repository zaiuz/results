package results

import v "github.com/zaiuz/views"
import z "github.com/zaiuz/zaiuz"

// Controller mixin to provide simpler and straightforward result building in controllers.
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

func (c *Controller) Redirect(permanent bool, url string) z.Result {
	return Redirect(permanent, url)
}

func (c *Controller) String(code int, str string) z.Result {
	return String(code, str)
}

func (c *Controller) View(view v.View, data interface{}) z.Result {
	return View(view, data)
}
