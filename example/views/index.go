// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package views

import (
	"github.com/mgutz/razor/razor"
)

// Index is generated
func Index(name string) razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	title := "Razor View Engine"
	_buffer.WriteString("\n\n<h1>Hello ")
	_buffer.WriteSafe(name)
	_buffer.WriteString("!<h1>")

	_buffer = Layout(title, _buffer)
	return _buffer
}
