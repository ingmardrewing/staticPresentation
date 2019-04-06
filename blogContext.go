package staticPresentation

import (
	"path"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
	"github.com/ingmardrewing/staticUtil"
)

func NewBlogContext(site staticIntf.Site) staticIntf.Context {
	tool := staticUtil.NewPagesContainerCollectionTool(site)

	cg := new(blogContext)
	cg.site = site

	archivePage := cg.makeArchive()
	site.AddMain(archivePage)

	cg.archiveRenderer = NewBlogArchiveRenderer(site)
	cg.archiveRenderer.Pages(archivePage)

	cg.renderer = NewBlogRenderer(site)
	cg.renderer.Pages(tool.GetPagesByVariant(staticIntf.BLOG)...)

	cg.naviRenderer = NewBlogNaviRenderer(site)
	cg.naviRenderer.Pages(tool.GetNaviPagesByVariant(staticIntf.BLOG)...)

	cg.rssRenderer = NewRssRenderer(
		cg.getLastTenReversedPages(),
		path.Join(site.TargetDir(), "/blog/"),
		"/blog/",
		site.RssFilename())

	return cg
}

type blogContext struct {
	abstractContext
	naviRenderer    staticIntf.Renderer
	archiveRenderer staticIntf.Renderer
	rssRenderer     staticIntf.Renderer
}

func (b *blogContext) makeArchive() staticIntf.Page {
	pm := staticModel.NewPageMaker()
	pm.Title("Archive")
	pm.Category("blog archive")
	pm.PathFromDocRoot("/blog/")
	pm.FileName("archive.html")
	pm.Site(b.site)
	return pm.Make()
}

func (b *blogContext) GetComponents() []staticIntf.Component {
	components := b.renderer.Components()
	components = append(components, b.archiveRenderer.Components()...)
	return append(components, b.naviRenderer.Components()...)
}

func (b *blogContext) RenderPages() []fs.FileContainer {
	fcs := b.renderer.Render()
	fcs = append(fcs, b.naviRenderer.Render()...)
	fcs = append(fcs, b.archiveRenderer.Render()...)
	fcs = append(fcs, b.rssRenderer.Render()...)
	return fcs
}
