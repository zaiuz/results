package results_test

import "testing"
import . "github.com/zaiuz/results"
import z "github.com/zaiuz/zaiuz"
import a "github.com/stretchr/testify/assert"

func TestDefer_Resolve(t *testing.T) {
	h := newContextHolder()
	resolve := func(c *z.Context) z.Result {
		h.calledContext = c
		return Dud()
	}

	result := Defer(resolve)
	a.NotNil(t, result, "cannot create deferred result.")

	e := result.Render(h.context)
	a.NoError(t, e)
	a.Equal(t, h.context, h.calledContext, "resolve function not given the right context.")
}

func TestDefer_Resolve_Nil(t *testing.T) {
	h := newContextHolder()
	resolve := func(c *z.Context) z.Result {
		h.calledContext = c
		return nil
	}

	result := Defer(resolve)
	test := func() {
		e := result.Render(h.context)
		a.NoError(t, e)
	}

	a.NotPanics(t, test, "panic when resolved to a nil result.")
}

func TestDefer_Render(t *testing.T) {
	resolve := func(c *z.Context) z.Result {
		return String(403, "OK")
	}

	result := Defer(resolve)
	RenderCheck(t, result).
		Code(403).
		Body("OK")
}
