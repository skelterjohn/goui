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
	data  ButtonData
	dirty bool
	
	Click <-chan []byte
}

func NewButton(label string) (me *Button) {
	me = new(Button)
	me.data.Label = label
	me.data.MName, me.Click = messages.GetMessenger()

	return
}

func (me *Button) SetData(data ButtonData) {
	me.data = data
	me.dirty = true
}

func (me *Button) GetData() ButtonData {
	return me.data
}

func (me *Button) Render(html io.Writer) (e error) {
	e = ButtonHTMLTemplate.Execute(html, me.data)
	return
}

func (me *Button) Size() (w, h int) {
	return
}

func (me *Button) Update() {

}
