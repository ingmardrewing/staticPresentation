package staticPresentation

import "github.com/ingmardrewing/staticIntf"

// Creates a new GlobalCssComponent
func NewGlobalCssComponent(r staticIntf.Renderer) *GlobalCssComponent {
	g := new(GlobalCssComponent)
	g.abstractComponent.Renderer(r)
	return g
}

type GlobalCssComponent struct {
	abstractComponent
}

func (gcc *GlobalCssComponent) GetCss() string {
	return `
/* open-sans-300 - latin-ext_latin */
@font-face {
font-family: 'Open Sans';
font-style: normal;
font-weight: 300;
	src: url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-300.eot'); /* IE9 Compat Modes */
	src: local('Open Sans Light'), local('OpenSans-Light'),
		url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-300.eot?#iefix') format('embedded-opentype'), /* IE6-IE8 */
		url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-300.woff2') format('woff2'), /* Super Modern Browsers */
		url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-300.woff') format('woff'), /* Modern Browsers */
		url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-300.ttf') format('truetype'), /* Safari, Android, iOS */
		url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-300.svg#OpenSans') format('svg'); /* Legacy iOS */
}

/* open-sans-700 - latin-ext_latin */
@font-face {
font-family: 'Open Sans';
font-style: normal;
font-weight: 700;
	src: url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-700.eot'); /* IE9 Compat Modes */
	src: local('Open Sans Bold'), local('OpenSans-Bold'),
		url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-700.eot?#iefix') format('embedded-opentype'), /* IE6-IE8 */
		url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-700.woff2') format('woff2'), /* Super Modern Browsers */
		url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-700.woff') format('woff'), /* Modern Browsers */
		url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-700.ttf') format('truetype'), /* Safari, Android, iOS */
		url('https://drewing.de/fonts/open-sans-v15-latin-ext_latin-700.svg#OpenSans') format('svg'); /* Legacy iOS */
}

body, p, span {
	margin: 0;
	padding: 0;
	font-family: 'Open Sans', Arial, Helvetica, sans-serif;
}
a {
	color: grey;
	text-decoration: none;
}
a:hover {
	text-decoration: underline;
}
.wrapperOuter {
	text-align: center;
}

.wrapperInner {
	margin: 0 auto;
	max-width: 800px;
}
p + p {
	margin-top: 10px;
}
`
}
