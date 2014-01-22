package results_test

import "testing"
import "net/http/httptest"
import "github.com/zaiuz/testutil"
import z "github.com/zaiuz/zaiuz"
import a "github.com/stretchr/testify/assert"

type ResultExpectable struct {
	t       *testing.T
	context *z.Context
	result  z.Result
}

func RenderCheck(t *testing.T, result z.Result) *ResultExpectable {
	response, request := testutil.NewTestRequestPair()
	context := z.NewContext(response, request)

	result.Render(context)
	return &ResultExpectable{t, context, result}
}

func (r *ResultExpectable) Code(code int) *ResultExpectable {
	a.Equal(r.t, code, r.recorder().Code, "response has wrong status code.")
	return r
}

func (r *ResultExpectable) Header(key, value string) *ResultExpectable {
	v, ok := r.recorder().Header()[key]
	a.True(r.t, ok, "response does not have `%s` header.", key)
	a.Equal(r.t, []string{value}, v, "value of `%s` header is wrong.", key)
	return r
}

func (r *ResultExpectable) Body(body string) *ResultExpectable {
	raw := r.recorder().Body.Bytes()
	a.Equal(r.t, body, string(raw), "response body does not match.")
	return r
}

// TODO: ResultExpectable.BodyPattern()

func (r *ResultExpectable) recorder() *httptest.ResponseRecorder {
	return r.context.ResponseWriter.(*httptest.ResponseRecorder)
}
