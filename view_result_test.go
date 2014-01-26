package results_test

import "testing"
import "os"
import "io/ioutil"
import . "github.com/zaiuz/results"
import v "github.com/zaiuz/views"
import z "github.com/zaiuz/zaiuz"
import a "github.com/stretchr/testify/assert"

var _ z.Result = &ViewResult{}

const (
	templateContent = `{{define "root"}}Hello, {{if .Data}}{{.Data}}{{else}}World!{{end}}{{end}}`
	templateOutput  = `Hello, World!`
)

func TestNewViewResult(t *testing.T) {
	view := getTestView(t)
	result := NewViewResult(view, nil)
	a.NotNil(t, result, "cannot create view result.")
}

func TestViewResult_Render(t *testing.T) {
	view := getTestView(t)
	result := NewViewResult(view, nil)

	RenderCheck(t, result).
		Code(200).
		Header("Content-Type", "text/plain").
		Body(templateOutput)
}

func TestViewResult_Render_WithData(t *testing.T) {
	m := map[string]string{"Data": `foobar`}

	view := getTestView(t)
	result := NewViewResult(view, m)

	RenderCheck(t, result).
		Code(200).
		Header("Content-Type", "text/plain").
		BodyContains(`foobar`)
}

func getTestView(t *testing.T) v.View {
	file, e := ioutil.TempFile("", "zaiuz-results")
	a.NoError(t, e)

	filename := file.Name()
	defer os.Remove(filename) // should be ok since views are pre-loaded

	_, e = file.WriteString(templateContent)
	a.NoError(t, e)
	file.Close()

	return v.NewTextView(filename)
}
