package results_test

import "testing"
import . "github.com/zaiuz/results"
import z "github.com/zaiuz/zaiuz"
import a "github.com/stretchr/testify/assert"

var _ z.Result = &JsonResult{}

var TestMap = map[string]string{"hello": "world"}
var TestMapJson = "{\"hello\":\"world\"}"

func TestJsonResult_Render(t *testing.T) {
	result := NewJsonResult(200, TestMap)
	a.NotNil(t, result, "cannot create json result.")

	RenderCheck(t, result).
		Code(200).
		Header("Content-Type", "application/json").
		Body(TestMapJson)
}
