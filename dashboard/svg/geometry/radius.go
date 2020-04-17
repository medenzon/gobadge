package geometry

import (
	"fmt"
)

type Radius Coordinate

func (radius Radius) Vectorize() string {
	x := fmt.Sprintf(`rx="%f"`, radius.X)
	y := fmt.Sprintf(`ry="%f"`, radius.Y)
	return fmt.Sprintf("%s %s", x, y)
}
