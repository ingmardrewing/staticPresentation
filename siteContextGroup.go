package staticPresentation

import "github.com/ingmardrewing/staticIntf"

func NewSiteContextGroup(
	pages []staticIntf.Page,
	cd staticIntf.CommonData) staticIntf.ContextGroup {

	cg := new(siteContextGroup)

	cg.pagesContext = NewPagesContext(cd)
	cg.pagesContext.SetElements(pages)

	return cg
}

type siteContextGroup struct {
	abstractContextGroup
	marginalContext staticIntf.Context
}

func ElementsToLocations(elements []staticIntf.Page) []staticIntf.Location {
	locs := []staticIntf.Location{}
	for _, p := range elements {
		locs = append(locs, p)
	}
	return locs
}
