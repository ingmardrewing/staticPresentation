package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"
)

func NewMarginalContextGroup(s staticIntf.Site) staticIntf.Context {
	tool := staticUtil.NewPagesContainerCollectionTool(s)

	cg := new(marginalContext)
	cg.site = s
	cg.marginalContext = NewMarginalRenderer(s)
	cg.marginalContext.Pages(tool.GetPagesByVariant(staticIntf.MARGINALS)...)
	return cg
}

type marginalContext struct {
	abstractContext
	marginalContext staticIntf.Renderer
}

func (s *marginalContext) GetComponents() []staticIntf.Component {
	return s.marginalContext.Components()
}

func (s *marginalContext) RenderPages() []fs.FileContainer {
	return s.marginalContext.Render()
}
