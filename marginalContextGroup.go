package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewMarginalContextGroup(
	marginalPages []staticIntf.Page,
	cd staticIntf.CommonData) staticIntf.ContextGroup {

	cg := new(marginalContextGroup)

	cg.marginalContext = NewMarginalContext(cd)
	cg.marginalContext.SetElements(marginalPages)

	locs := ElementsToLocations(marginalPages)
	for _, l := range locs {
		cd.AddMarginal(l)
	}

	return cg
}

type marginalContextGroup struct {
	abstractContextGroup
	marginalContext staticIntf.Context
}

func (s *marginalContextGroup) GetComponents() []staticIntf.Component {
	return s.marginalContext.GetComponents()
}

func (s *marginalContextGroup) RenderPages(dir string) []fs.FileContainer {
	return s.marginalContext.RenderPages(dir)
}
