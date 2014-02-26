package results_test

import "testing"
import "os"
import "io/ioutil"
import . "github.com/zaiuz/results"
import v "github.com/zaiuz/views"
import a "github.com/stretchr/testify/assert"

const (
	templateContent = `{{define "root"}}Hello, {{if .Data}}{{.Data}}{{else}}World!{{end}}{{end}}`
	templateOutput  = `Hello, World!`
)

func TestView_Render(t *testing.T) {
	view := getTestView(t)
	result := View(view, nil)
	a.NotNil(t, result, "cannot create view result.")

	RenderCheck(t, result).
		Code(200).
		Header("Content-Type", "text/html").
		Body(templateOutput)
}

func TestView_Render_WithData(t *testing.T) {
	m := map[string]string{"Data": `foobar`}

	view := getTestView(t)
	result := View(view, m)

	RenderCheck(t, result).
		Code(200).
		Header("Content-Type", "text/html").
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

	return v.NewHtmlView(filename)
}
