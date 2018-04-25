package staticPresentation

import (
	"path"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewPortfolioContext(s staticIntf.Site) staticIntf.Context {
	cg := new(portfolioContext)
	cg.site = s

	cg.renderer = NewPortfolioRenderer(s)
	cg.renderer.Pages(s.Portfolio()...)

	return cg
}

type portfolioContext struct {
	abstractContext
}

func (b *portfolioContext) RenderPages() []fs.FileContainer {
	fcs := b.renderer.Render()
	rr := NewRssRenderer(
		b.getLastTenReversedPages(),
		path.Join(b.site.TargetDir(), "/portfolio/"),
		"/portfolio/",
		b.site.RssFilename())
	rssFc := rr.Render()

	if rssFc != nil {
		fcs = append(fcs, rssFc)
	}

	return fcs
}
