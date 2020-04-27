package geometry

import (
	"fmt"
)

type Coordinate struct {
	X float64
	Y float64
}

func (coordinate Coordinate) Vectorize() string {
	x := fmt.Sprintf(`x="%f"`, coordinate.X)
	y := fmt.Sprintf(`y="%f"`, coordinate.Y)
	return fmt.Sprintf("%s %s", x, y)
}
