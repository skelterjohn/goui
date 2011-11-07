package goui

import (
	"io"
)

type Component interface {
	// write out the client-side code for this component
	Render(html io.Writer) (e error)
	// send a message to the client-side code with changes that need to be updated
	Update()
	// the pixel space that this component takes up (for use w/ layouts)
	Size() (w, h int)
}

func ComponentWriter(c Component) func(io.Writer) error {
	return func(html io.Writer) (err error) {
		return c.Render(html)
	}
}
