package results_test

import "testing"
import . "github.com/zaiuz/results"
import z "github.com/zaiuz/zaiuz"
import a "github.com/stretchr/testify/assert"

var _ z.Result = &StringResult{}

var TestString = "Hello, World!"

func TestStringResult_Render(t *testing.T) {
	result := NewStringResult(200, TestString)
	a.NotNil(t, result, "cannot create string result.")

	RenderCheck(t, result).
		Code(200).
		Header("Content-Type", "text/plain").
		Body(TestString)
}
