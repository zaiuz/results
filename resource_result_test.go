package results_test

import "testing"
import . "github.com/zaiuz/results"
import a "github.com/stretchr/testify/assert"
import z "github.com/zaiuz/zaiuz"

func TestResource(t *testing.T) {
	resource := Resource()
	a.NotNil(t, resource, "resource should returns a non-nil builder.")
}

func TestResource_Render_Json(t *testing.T) {
	result := newTestResourceResult(t)
	headers := map[string][]string{
		"Accept": []string{"application/json"},
	}

	RenderCheck2(t, result, headers).
		Code(200).
		Header("Content-Type", "application/json").
		Body("\"json result\"")
}

func TestResource_Render_Text(t *testing.T) {
	result := newTestResourceResult(t)
	headers := map[string][]string{
		"Accept": []string{"text/plain"},
	}

	RenderCheck2(t, result, headers).
		Code(200).
		Header("Content-Type", "text/plain").
		Body("string result")
}

func TestResource_Render_Html(t *testing.T) {
	result := newTestResourceResult(t)
	headers := map[string][]string{
		"Accept": []string{"text/html"},
	}

	RenderCheck2(t, result, headers).
		Code(200).
		Header("Content-Type", "text/html").
		Body(templateOutput)
}

func TestResource_Render_Mime(t *testing.T) {
	result := newTestResourceResult(t)
	headers := map[string][]string{
		"Accept": []string{"custom/type"},
	}

	RenderCheck2(t, result, headers).
		Code(200).
		Header("Content-Type", "text/plain").
		Body("custom result")
}

func newTestResourceResult(t *testing.T) z.Result {
	return Resource().
		Json(Json(200, "json result")).
		Text(String(200, "string result")).
		Html(View(getTestView(t), nil)).
		Mime("custom/type", String(200, "custom result"))
}
