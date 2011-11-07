package components

import (
	"io"
	"template"
	"github.com/skelterjohn/goui"
	"github.com/skelterjohn/goui/messages"
)

const ButtonHTMLTemplateFormat = `
<a href=/m/{{.MName}}>{{.Label}}</a>
`

var ButtonHTMLTemplate = template.Must(goui.ParseExecTemplate(ButtonHTMLTemplateFormat))

type ButtonData struct {
	Label string
	MName string
}

type Button struct {
	bd ButtonData
	Click <-chan []byte
}

func NewButton(label string) (me *Button) {
	me = new(Button)
	me.bd.Label = label
	me.bd.MName, me.Click = messages.GetMessenger()

	return
}

func (me *Button) Render(html io.Writer) (e error) {
	e = ButtonHTMLTemplate.Execute(html, me.bd)
	return
}

func (me *Button) Size() (w, h int) {
	return
}

func (me *Button) Update() {
	
}