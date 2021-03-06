package results_test

import "testing"
import . "github.com/zaiuz/results"
import a "github.com/stretchr/testify/assert"

func TestRedirect_Render_Temporary(t *testing.T) {
	result := Redirect("/%s/url", "another")
	a.NotNil(t, result, "cannot create temporary redir result.")

	RenderCheck(t, result).
		Code(302).
		HeaderContains("Location", "/another/url").
		EmptyBody()
}

func TestRedirect_Render_Permanent(t *testing.T) {
	result := PermanentRedirect("/%s/url", "another")
	a.NotNil(t, result, "cannot create permanent redir result.")

	RenderCheck(t, result).
		Code(301).
		HeaderContains("Location", "/another/url").
		EmptyBody()
}

func TestRedirect_Render_WithArgs(t *testing.T) {
	result := Redirect("/section/%s/subsection/%d", "hello", 1)
	a.NotNil(t, result, "cannot create redir result with args.")

	RenderCheck(t, result).
		HeaderContains("Location", "/section/hello/subsection/1").
		EmptyBody()
}
