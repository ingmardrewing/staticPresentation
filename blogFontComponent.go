package staticPresentation

import "github.com/ingmardrewing/staticIntf"

// Creates a new CssLinkComponent
func NewBlogFontComponent(r staticIntf.Renderer) *BlogFontComponent {
	bfc := new(BlogFontComponent)
	bfc.abstractComponent.Renderer(r)
	return bfc
}

type BlogFontComponent struct {
	abstractComponent
}

func (bfc *BlogFontComponent) GetCss() string {
	return `@import url('https://fonts.googleapis.com/css?family=Open+Sans:300,700');
`
}
