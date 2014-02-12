package results_test

import "testing"
import . "github.com/zaiuz/results"
import a "github.com/stretchr/testify/assert"

func TestRedirect_Render_Temporary(t *testing.T) {
	result := Redirect(false, "/another/url")
	a.NotNil(t, result, "cannot create temporary redir result.")

	RenderCheck(t, result).
		Code(302).
		HeaderContains("Location", "/another/url").
		EmptyBody()
}

func TestRedirect_Render_Permanent(t *testing.T) {
	result := Redirect(true, "/another/url")
	a.NotNil(t, result, "cannot create permanent redir result.")

	RenderCheck(t, result).
		Code(301).
		HeaderContains("Location", "/another/url").
		EmptyBody()
}
