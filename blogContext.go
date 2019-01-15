package staticPresentation

import (
	"path"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"
)

func NewBlogContext(s staticIntf.Site) staticIntf.Context {

	tool := staticUtil.NewPagesContainerCollectionTool(s)

	cg := new(blogContext)
	cg.site = s

	cg.renderer = NewBlogRenderer(s)
	cg.renderer.Pages(tool.GetPagesByVariant(staticIntf.BLOG)...)

	cg.naviRenderer = NewBlogNaviRenderer(s)
	cg.naviRenderer.Pages(tool.GetNaviPagesByVariant(staticIntf.BLOG)...)

	return cg
}

type blogContext struct {
	abstractContext
	naviRenderer staticIntf.Renderer
}

func (b *blogContext) GetComponents() []staticIntf.Component {
	components := b.renderer.Components()
	return append(components, b.naviRenderer.Components()...)
}

func (b *blogContext) RenderPages() []fs.FileContainer {
	fcs := b.renderer.Render()
	fcs = append(fcs, b.naviRenderer.Render()...)

	rr := NewRssRenderer(
		b.getLastTenReversedPages(),
		path.Join(b.site.TargetDir(), "/blog/"),
		"/blog/",
		b.site.RssFilename())
	rssFc := rr.Render()

	if rssFc != nil {
		fcs = append(fcs, rssFc)
	}

	return fcs
}
