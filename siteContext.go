package staticPresentation

import (
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"
)

func NewSiteContextGroup(s staticIntf.Site) staticIntf.Context {

	tool := staticUtil.NewPagesContainerCollectionTool(s)

	cg := new(abstractContext)
	cg.site = s
	cg.renderer = NewPagesRenderer(s)
	cg.renderer.Pages(tool.GetPagesByVariant(staticIntf.PAGES)...)

	return cg
}
