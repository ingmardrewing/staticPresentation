package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewMarginalContextGroup(s staticIntf.Site) staticIntf.ContextGroup {

	cg := new(marginalContextGroup)

	cg.site = s
	cg.marginalContext = NewMarginalContext(s)
	cg.marginalContext.SetElements(s.Marginals())

	locs := ElementsToLocations(s.Marginals())
	for _, l := range locs {
		s.AddMarginal(l)
	}

	return cg
}

type marginalContextGroup struct {
	abstractContextGroup
	marginalContext staticIntf.SubContext
}

func (s *marginalContextGroup) GetComponents() []staticIntf.Component {
	return s.marginalContext.GetComponents()
}

func (s *marginalContextGroup) RenderPages() []fs.FileContainer {
	return s.marginalContext.RenderPages()
}
