package staticPresentation

import (
	"strconv"
	"time"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creats a new CopyRightComponent
func NewCopyRightComponent(r staticIntf.Renderer) *CopyRightComponent {
	c := new(CopyRightComponent)
	c.abstractComponent.Renderer(r)
	return c
}

type CopyRightComponent struct {
	abstractComponent
	wrapper
}

// TODO: Separate content from form:
func (crc *CopyRightComponent) VisitPage(p staticIntf.Page) {
	year := p.PublishedTime("2006")
	if len(year) != 4 {
		year = strconv.Itoa(time.Now().Year())
	}
	n := htmlDoc.NewNode("div", `<a rel="license" class="copyright__cc" href="https://creativecommons.org/licenses/by-nc-nd/3.0/"></a><p class="copyright__license">&copy; `+year+` by Ingmar Drewing </p><p class="copyright__license">Except where otherwise noted, content on this site is licensed under a <a rel="license" href="https://creativecommons.org/licenses/by-nc-nd/3.0/">Creative Commons Attribution-NonCommercial-NoDerivs 3.0 Unported (CC BY-NC-ND 3.0) license</a>. Unless otherwise noted the author of the content on this page is Ingmar Drewing</p><p class="copyright__license">Soweit nicht anders explizit ausgewiesen, stehen die Inhalte auf dieser Website unter der <a rel="license" href="https://creativecommons.org/licenses/by-nc-nd/3.0/">Creative Commons Namensnennung-NichtKommerziell-KeineBearbeitung (CC BY-NC-ND 3.0)</a> Lizenz. </p>`, "class", "copyright")
	wn := crc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (crc *CopyRightComponent) GetCss() string {
	return `
.copyright {
	background-color: black;
	color: white;
	text-align: left;
	font-size: 14px;
	padding: 20px 20px 50px;
	margin-top: 20px;
}
.copyright__license {
	margin-top: 20px;
	margin-bottom: 20px;
}
.copyright__cc {
    display: block;
    border-width: 0;
    background-image: url(https://i.creativecommons.org/l/by-nc-nd/3.0/88x31.png);
    width: 88px;
    height: 31px;
    margin-right: 15px;
    margin-bottom: 5px;
}
`
}
