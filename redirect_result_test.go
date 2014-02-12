package results_test

import "testing"
import . "github.com/zaiuz/results"
import z "github.com/zaiuz/zaiuz"
import a "github.com/stretchr/testify/assert"

var _ z.Result = &RedirectResult{}

func TestRedirectResult_Render_Temporary(t *testing.T) {
	result := NewRedirectResult(false, "/another/url")
	a.NotNil(t, result, "cannot create temporary redir result.")

	RenderCheck(t, result).
		Code(302).
		HeaderContains("Location", "/another/url").
		EmptyBody()
}

func TestRedirectResult_Render_Permanent(t *testing.T) {
	result := NewRedirectResult(true, "/another/url")
	a.NotNil(t, result, "cannot create permanent redir result.")

	RenderCheck(t, result).
		Code(301).
		HeaderContains("Location", "/another/url").
		EmptyBody()
}
