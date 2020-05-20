package svg

import (
	"fmt"
	"io"
	"strings"
)

// Canvas is the SVG drawing canvas.
type Canvas struct {
	Writer io.Writer
}

// New creates a new SVG drawing canvas.
func New(writer io.Writer) *Canvas {
	return &Canvas{writer}
}

func (canvas *Canvas) Open(width float64, height float64) error {

	open, close := `<svg `, `>`
	_width := fmt.Sprintf(`width="%f"`, width)
	_height := fmt.Sprintf(` height="%f"`, height)
	version := `version="1.1"`
	xmlns := `xmlns="http://www.w3.org/2000/svg"`
	xlink := `xmlns:xlink="http://www.w3.org/1999/xlink"`

	elements := []string{open, _width, _height, version, xmlns, xlink, close}
	output := strings.Join(elements, " ")

	_, err := fmt.Fprint(canvas.Writer, output)

	return err
}

func (canvas *Canvas) Close() error {
	_, err := fmt.Fprint(canvas.Writer, `</svg>`)
	return err
}
