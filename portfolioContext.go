package staticPresentation

import (
	"path"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"
)

func NewPortfolioContext(site staticIntf.Site) staticIntf.Context {

	cg := new(portfolioContext)
	cg.site = site

	tool := staticUtil.NewPagesContainerCollectionTool(site)
	cg.renderer = NewPortfolioRenderer(site)
	cg.renderer.Pages(tool.GetPagesByVariant(staticIntf.PORTFOLIO)...)

	cg.rssRenderer = NewRssRenderer(
		cg.getLastTenReversedPages(),
		path.Join(site.TargetDir(), "/portfolio/"),
		"/portfolio/",
		site.RssFilename())

	return cg
}

/*
 * portfolioContext
 */
type portfolioContext struct {
	abstractContext
	rssRenderer staticIntf.Renderer
}

func (b *portfolioContext) RenderPages() []fs.FileContainer {
	fileContainers := b.renderer.Render()
	fileContainers = append(fileContainers, b.rssRenderer.Render()...)
	return fileContainers
}
