package staticPresentation

import (
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"
)

func NewHomeContextGroup(s staticIntf.Site) staticIntf.Context {

	cg := new(abstractContext)
	cg.site = s
	cg.renderer = NewHomePageRenderer(s)

	tool := staticUtil.NewPagesContainerCollectionTool(s)
	cg.renderer.Pages(tool.GetPagesByVariant(staticIntf.HOME)...)

	return cg
}
