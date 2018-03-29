package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Create new  FaviconComponent
func NewFaviconComponent() *FaviconComponent {
	return new(FaviconComponent)
}

type FaviconComponent struct {
	abstractComponent
}

func (f *FaviconComponent) VisitPage(p staticIntf.Page) {
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("link", "", "rel", "icon", "type", "image/png", "sizes", "16x16", "href", "/icons/favicon-16x16.png"),
		htmlDoc.NewNode("link", "", "rel", "icon", "type", "image/png", "sizes", "32x32", "href", "/icons/favicon-32x32.png"),
		htmlDoc.NewNode("link", "", "rel", "icon", "type", "image/png", "sizes", "192x192", "href", "/icons/android-192x192.png"),
		htmlDoc.NewNode("link", "", "rel", "apple-touch-icon", "type", "image/png", "sizes", "180x180", "href", "/icons/apple-touch-icon-180x180.png"),
		htmlDoc.NewNode("meta", "", "name", "msapplication-config", "content", "/icons/browserconfig.xml")}
	p.AddHeaderNodes(m)
}
