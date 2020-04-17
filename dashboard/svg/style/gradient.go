package style

import (
	"fmt"
	"strings"
)

const (
	pct = "%"
)

// LinearGradient represents an SVG linear gradient def.
type LinearGradient struct {
	ID     string
	Top    Point
	Bottom Point
	Stops  Stops
}

// Point represents an SVG gradient point.
type Point struct {
	X int
	Y int
}

// Stops is an array of SVG gradient color stops.
type Stops []Stop

// Stop represents an SVG gradient color stop.
type Stop struct {
	Color   string
	Opacity float64
	Offset  int
}

// Vectorize converts a LinearGradient struct to an SVG string.
func (gradient LinearGradient) Vectorize() string {
	top := gradient.Top.vectorize(1)
	bottom := gradient.Bottom.vectorize(2)
	id := gradient.ID
	stops := gradient.Stops.vectorize()
	template := "<linearGradient %s %s id=\"%s\">\n%s\n</linearGradient>"
	return fmt.Sprintf(template, top, bottom, id, stops)
}

func (point Point) vectorize(n int) string {
	x := fmt.Sprintf(`x%d="%d%s"`, n, point.X, pct)
	y := fmt.Sprintf(`y%d="%d%s"`, n, point.Y, pct)
	return fmt.Sprintf("%s %s", x, y)
}

func (stops Stops) vectorize() string {
	lines := []string{}
	for _, stop := range stops {
		lines = append(lines, stop.vectorize())
	}
	return strings.Join(lines, "\n")
}

func (stop Stop) vectorize() string {
	color := fmt.Sprintf(`stop-color="%s"`, stop.Color)
	opacity := fmt.Sprintf(`stop-opacity="%f"`, stop.Opacity)
	offset := fmt.Sprintf(`offset="%d%s"`, stop.Offset, pct)
	return fmt.Sprintf("<stop %s %s %s ></stop>", color, opacity, offset)
}
