package goui

import (
	"bytes"
	"template"
	"io"
)

//i'd really like to not buffer, here, but it is what it is
func compWriter(c Component) (r io.Reader, err error) {
	b := bytes.NewBuffer(make([]byte, 1024))
	if c != nil {
		err = c.Render(b)
	}
	r = b
	return
}

var execMap = template.FuncMap{ "ComponentWriter": compWriter }

func ParseExecTemplate(format string) (t *template.Template, err error) {
	execTemplate := template.New("ComponentTemplate").Funcs(execMap)	
	return execTemplate.Parse(format)
}
