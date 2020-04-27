package shapes

import (
	"fmt"
)

type Point Coordinate

func (point Point) Vectorize() string {
	x := fmt.Sprintf(`x="%f"`, point.X)
	y := fmt.Sprintf(`y="%f"`, point.Y)
	return fmt.Sprintf("%s %s", x, y)
}
