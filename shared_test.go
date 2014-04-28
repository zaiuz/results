package results_test

import "testing"
import "net/http/httptest"
import "github.com/zaiuz/testutil"
import z "github.com/zaiuz/zaiuz"
import a "github.com/stretchr/testify/assert"

type ContextHolder struct {
	context       *z.Context
	calledContext *z.Context
}

func NewContextHolder() *contextHolder {
	context := testutil.NewTestContext()
	return &contextHolder{context, nil}
}

type ResultExpectable struct {
	t       *testing.T
	context *z.Context
	result  z.Result
	err     error
}

func RenderCheck(t *testing.T, result z.Result) *ResultExpectable {
	return RenderCheck2(t, result, nil)
}

func RenderCheck2(t *testing.T, r z.Result, headers map[string][]string) *ResultExpectable {
	response, request := testutil.NewTestRequestPair()
	for key, values := range headers {
		request.Header[key] = values
	}

	context := z.NewContext(response, request)
	e := r.Render(context)
	return &ResultExpectable{t, context, r, e}

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

func (r *ResultExpectable) HeaderContains(key, query string) *ResultExpectable {
	v, ok := r.recorder().Header()[key]
	a.True(r.t, ok, "response does not have `%s` header.", key)
	a.Contains(r.t, v[0], query, "header `%s` does not contains `%s`", key, query)
	return r
}

func (r *ResultExpectable) EmptyBody() *ResultExpectable {
	a.NoError(r.t, r.err)
	raw := r.recorder().Body.Bytes()

	a.Equal(r.t, len(raw), 0, "response unexpectedly contains body.")
	return r
}

func (r *ResultExpectable) Body(body string) *ResultExpectable {
	a.NoError(r.t, r.err)

	raw := r.recorder().Body.Bytes()
	a.Equal(r.t, body, string(raw), "response body does not match.")
	return r
}

func (r *ResultExpectable) BodyContains(query string) *ResultExpectable {
	a.NoError(r.t, r.err)

	raw := r.recorder().Body.Bytes()
	a.Contains(r.t, string(raw), query, "response body does not contains expected string.")
	return r
}

// TODO: ResultExpectable.BodyPattern()

func (r *ResultExpectable) recorder() *httptest.ResponseRecorder {
	return r.context.ResponseWriter.(*httptest.ResponseRecorder)
}
