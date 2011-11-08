package components

import (
	"io"
	"template"
	"github.com/skelterjohn/goui"
)

const LabelHTMLTemplateFormat = `
Label: {{.Text}}
`

var LabelHTMLTemplate = template.Must(goui.ParseExecTemplate(LabelHTMLTemplateFormat))

type LabelData struct {
	Text string
}

type Label struct {
	data LabelData
}

func NewLabel(text string) (me *Label) {
	me = new(Label)
	me.data.Text = text
	return
}

func (me *Label) SetText(text string) {
	me.data.Text = text
}

func (me *Label) Render(html io.Writer) (e error) {
	e = LabelHTMLTemplate.Execute(html, me.data)
	return
}

func (me *Label) Update() {

}

func (me *Label) Size() (w, h int) {
	return
}
