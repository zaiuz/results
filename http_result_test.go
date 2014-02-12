package results_test

import "testing"
import . "github.com/zaiuz/results"
import a "github.com/stretchr/testify/assert"

func TestHttp_Render(t *testing.T) {
	result := Http(302, "Location", "example.com")
	a.NotNil(t, result, "cannot create http result.")

	RenderCheck(t, result).Code(302).Header("Location", "example.com")
}
