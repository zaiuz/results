package results_test

import "testing"
import . "github.com/zaiuz/results"
import z "github.com/zaiuz/zaiuz"
import a "github.com/stretchr/testify/assert"

var _ z.Result = &HttpResult{}

func TestNewHttpResult(t *testing.T) {
	result := NewHttpResult(123, "Content-Type", "text/plain").(*HttpResult)
	a.NotNil(t, result, "cannot create base http result.")
	a.Equal(t, 123, result.Code, "status code not saved.")
	a.NotNil(t, result.Headers, "headers map initialized.")

	contentType, ok := result.Headers["Content-Type"]
	a.True(t, ok, "given header not saved.")
	a.Equal(t, []string{"text/plain"}, contentType, "wrong header value saved.")

	// WAIT: https://github.com/stretchr/testify/issues/34
	result = NewHttpResult(123, "").(*HttpResult)
	a.Equal(t, 0, len(result.Headers), "headers list not empty when empty string given.")
}

func TestHttpResult_Execute(t *testing.T) {
	result := NewHttpResult(302, "Location", "example.com").(*HttpResult)
	result.Headers["Accept"] = []string{"text/plain"}

	RenderCheck(t, result).Code(302).Header("Location", "example.com")
}
