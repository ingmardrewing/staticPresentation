package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// TODO: Separate content from implementation
// Creates a new NarrativeCopyRightComponent
func NewNarrativeCopyRightComponent() *NarrativeCopyRightComponent {
	c := new(NarrativeCopyRightComponent)
	return c
}

// NarrativeCopyRightComponent
type NarrativeCopyRightComponent struct {
	abstractComponent
	wrapper
}

func (crc *NarrativeCopyRightComponent) VisitPage(p staticIntf.Page) {
	n := htmlDoc.NewNode("div", `All content including but not limited to the art, characters, story, website design & graphics are © copyright 2013-2018 Ingmar Drewing unless otherwise stated. All rights reserved. Do not copy, alter or reuse without expressed written permission.`, "class", "copyright")
	wn := crc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (crc *NarrativeCopyRightComponent) GetCss() string {
	return `
.copyright {
	background-color: black;
	color: white;
	text-align: left;
	font-family: Arial, Helvetica, sans-serif;
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
