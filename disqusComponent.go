package staticPresentation

import (
	"fmt"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new DisqusComponent
func NewDisqusComponent(r staticIntf.Renderer) *DisqusComponent {
	d := new(DisqusComponent)
	d.abstractComponent.Renderer(r)
	return d
}

type DisqusComponent struct {
	abstractComponent
	wrapper
	configuredJs string
}

func (dc *DisqusComponent) GetJs() string {
	return dc.configuredJs
}

func (dc *DisqusComponent) VisitPage(p staticIntf.Page) {
	dc.configuredJs = fmt.Sprintf(`var disqus_config = function () { this.page.title= "%s"; this.page.url = '%s'; this.page.identifier =  '%s'; }; (function() { var d = document, s = d.createElement('script'); s.src = 'https://%s.disqus.com/embed.js'; s.setAttribute('data-timestamp', +new Date()); (d.head || d.body).appendChild(s); })();`, p.Title(),
		p.Url(), p.DisqusId(), p.Site().DisqusId())
	n := htmlDoc.NewNode("div", " ", "id", "disqus_thread", "class", "disqus")
	js := htmlDoc.NewNode("script", dc.configuredJs)
	wn := dc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn, js})
}
