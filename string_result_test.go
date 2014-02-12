package results_test

import "testing"
import . "github.com/zaiuz/results"
import a "github.com/stretchr/testify/assert"

func TestString_Render(t *testing.T) {
	const Str = "Hello, World!"

	result := String(200, Str)
	a.NotNil(t, result, "cannot create string result.")

	RenderCheck(t, result).
		Code(200).
		Header("Content-Type", "text/plain").
		Body(Str)
}
