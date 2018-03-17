package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

type navigationalContext struct {
	abstractContext
	naviRenderer staticIntf.Renderer
}

func (n *navigationalContext) GetComponents() []staticIntf.Component {
	components := n.renderer.GetComponents()
	return append(components, n.naviRenderer.GetComponents()...)
}

func (n *navigationalContext) RenderPages() []fs.FileContainer {
	pages := n.renderer.Render()
	return append(pages, n.naviRenderer.Render()...)
}
