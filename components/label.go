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
	ld LabelData
}

func NewLabel(text string) (me *Label) {
	me = new(Label)
	me.ld.Text = text
	return
}

func (me *Label) SetText(text string) {
	me.ld.Text = text
}

func (me *Label) Render(html io.Writer) (e error) {
	e = LabelHTMLTemplate.Execute(html, me.ld)
	return
}

func (me *Label) Update() {

}

func (me *Label) Size() (w, h int) {
	return
}
