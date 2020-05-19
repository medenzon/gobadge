package shapes

import (
	"fmt"
	"strings"

	geo "gobadge/svg/geometry"
)

type Rect struct {
	ID       string
	Fill     string
	FillRule string
	Origin   geo.Coordinate
	Radius   geo.Radius
	Frame    geo.Frame
}

func (path Rect) Vectorize() string {

	elements := strings.Join([]string{
		fmt.Sprintf(`id="%s"`, path.ID),
		fmt.Sprintf(`fill="%s"`, path.Fill),
		fmt.Sprintf(`fill-rule="%s"`, path.FillRule),
		path.Origin.Vectorize(),
		path.Radius.Vectorize(),
		path.Frame.Vectorize(),
	}, " ")

	template := `<rect %s></rect>`

	return fmt.Sprintf(template, elements)
}
