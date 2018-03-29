package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new GalleryComponent
func NewGalleryComponent() *GalleryComponent {
	gc := new(GalleryComponent)
	return gc
}

type GalleryComponent struct {
	wrapper
}

func (gal *GalleryComponent) VisitPage(p staticIntf.Page) {
	inner := htmlDoc.NewNode("div", "", "class", "maincontent__inner")
	for i := 0; i < 5; i++ {
		title := htmlDoc.NewNode("span", "At The Zoo", "class", "portfoliothumb__title")
		subtitle := htmlDoc.NewNode("span", "Digital drawing", "class", "portfoliothumb__details")

		label := htmlDoc.NewNode("div", "", "class", "portfoliothumb__label")
		label.AddChild(title)
		label.AddChild(subtitle)

		img := htmlDoc.NewNode("img", "", "class", "portfoliothumb__image", "src", "https://s3.amazonaws.com/drewingdeblog/blog/wp-content/uploads/2017/12/02152842/atthezoo-400x400.png")

		div := htmlDoc.NewNode("a", "", "class", "portfoliothumb", "href", "https://drewing.de")
		div.AddChild(img)
		div.AddChild(label)

		inner.AddChild(div)
	}

	m := htmlDoc.NewNode("main", "", "class", "maincontent")
	m.AddChild(inner)
	wn := gal.wrap(m)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}
