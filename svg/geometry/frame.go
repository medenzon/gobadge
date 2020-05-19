package geometry

import (
	"fmt"
)

type Frame struct {
	Width  float64
	Height float64
}

func (frame Frame) Vectorize() string {
	width := fmt.Sprintf(`width="%f"`, frame.Width)
	height := fmt.Sprintf(`height="%f"`, frame.Height)
	return fmt.Sprintf("%s %s", width, height)
}
