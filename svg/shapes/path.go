package shapes

import (
	"fmt"
)

var path = `
<path id="block"
		fill="#555555"
		fill-rule="nonzero"
		d="M3,0 L125,0 L125,0 L125,20 L3,20 C1.34314575,20 2.02906125e-16,18.6568542 0,17 L0,3 C-2.02906125e-16,1.34314575 1.34314575,3.04359188e-16 3,0 Z">
</path>

<path id="color"
		fill="#2B2B2B"
		fill-rule="nonzero"
		d="M125,0 L227,0 C228.656854,-3.04359188e-16 230,1.34314575 230,3 L230,17 C230,18.6568542 228.656854,20 227,20 L125,20 L125,20 L125,0 Z">
</path>
`

type Path struct {
	ID       string
	Fill     string
	FillRule string
	D        string
}

func (path Path) Vectorize() string {
	template := `<path id="%s" fill="%s" fill-rule="%s" d="%s"> </path>`
	return fmt.Sprintf(template, path.ID, path.Fill, path.FillRule, path.D)
}
