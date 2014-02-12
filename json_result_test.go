package results_test

import "testing"
import . "github.com/zaiuz/results"
import a "github.com/stretchr/testify/assert"

func TestJson_Render(t *testing.T) {
	input := map[string]string{"hello": "world"}
	output := "{\"hello\":\"world\"}"

	result := Json(200, input)
	a.NotNil(t, result, "cannot create json result.")

	RenderCheck(t, result).
		Code(200).
		Header("Content-Type", "application/json").
		Body(output)
}
