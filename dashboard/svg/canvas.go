package svg

import (
	"fmt"
	"io"
)

// Canvas is the SVG drawing canvas.
type Canvas struct {
	Writer io.Writer
}

// New creates a new SVG drawing canvas.
func New(writer io.Writer) *Canvas {
	return &Canvas{writer}
}

func (canvas *Canvas) Open() error {
	_, err := fmt.Fprint(canvas.Writer, `<svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">`)
	return err
}

func (canvas *Canvas) Close() error {
	_, err := fmt.Fprint(canvas.Writer, `</svg>`)
	return err
}
