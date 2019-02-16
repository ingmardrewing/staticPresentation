package staticPresentation

import (
	"fmt"
	"time"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// TODO: Separate content from implementation
// Creates a new NarrativeCopyRightComponent
func NewNarrativeCopyRightComponent(r staticIntf.Renderer) *NarrativeCopyRightComponent {
	c := new(NarrativeCopyRightComponent)
	c.abstractComponent.Renderer(r)
	return c
}

// NarrativeCopyRightComponent
type NarrativeCopyRightComponent struct {
	abstractComponent
	wrapper
}

func (crc *NarrativeCopyRightComponent) VisitPage(p staticIntf.Page) {
	currentYear := time.Now().Year()
	txt := fmt.Sprintf("All content including but not limited to the art, characters, story, website design & graphics are © copyright 2013-%d Ingmar Drewing unless otherwise stated. All rights reserved. Do not copy, alter or reuse without expressed written permission.", currentYear)
	n := htmlDoc.NewNode("div", txt, "class", "copyright")
	wn := crc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (crc *NarrativeCopyRightComponent) GetCss() string {
	return `
.copyright {
	background-color: black;
	color: white;
	text-align: left;
	font-size: 14px;
	padding: 20px 20px 50px;
	margin-top: 20px;
}
.copyright__license {
	margin-top: 20px;
	margin-bottom: 20px;
}
.copyright__cc {
    display: block;
    border-width: 0;
    background-image: url(https://i.creativecommons.org/l/by-nc-nd/3.0/88x31.png);
    width: 88px;
    height: 31px;
    margin-right: 15px;
    margin-bottom: 5px;
}
`
}
