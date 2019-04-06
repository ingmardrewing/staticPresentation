package staticPresentation

import (
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"
)

func NewHomeContext(site staticIntf.Site) staticIntf.Context {
	tool := staticUtil.NewPagesContainerCollectionTool(site)
	renderer := NewHomePageRenderer(site)
	renderer.Pages(tool.GetPagesByVariant(staticIntf.HOME)...)

	return &abstractContext{
		site:     site,
		renderer: renderer,
	}
}
