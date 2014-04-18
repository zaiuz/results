package results_test

import "testing"
import . "github.com/zaiuz/results"
import z "github.com/zaiuz/zaiuz"
import a "github.com/stretchr/testify/assert"

func TestFunc_Render(t *testing.T) {
	h := newContextHolder()
	execute := func(c *z.Context) error {
		h.calledContext = c
		return nil
	}

	result := RenderFunc(execute)
	a.NotNil(t, result, "cannot create result from a function.")

	e := result.Render(h.context)
	a.NoError(t, e)
	a.Equal(t, h.context, h.calledContext, "function given wrong context instance.")
}

func TestDud_Render(t *testing.T) {
	result := Dud()
	a.NotNil(t, result, "cannot create dud result.")

	test := func() {
		e := result.Render(nil)
		a.NoError(t, e)
	}

	a.NotPanics(t, test, "dud result panics on nil context.")
}
