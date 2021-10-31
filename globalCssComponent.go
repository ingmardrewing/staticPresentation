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
/* EB Garamond */
@font-face {
font-family: 'EB Garmond';
font-weight: 400;
	src: local('EB Garamond'), local('EBGaramond-Regular'),
		url('/fonts/EBGaramond-Regular.ttf') format('truetype');
}

/* EB Garamond italic
@font-face {
font-family: 'EB Garmond';
font-style: italic;
font-weight: 400;
	src: local('EB Garamond'), local('EBGaramond-Italic'),
		url('/fonts/EBGaramond-Italic.ttf') format('truetype');
}*/

/* open-sans-300 - latin-ext_latin */
@font-face {
font-family: 'Open Sans';
font-style: normal;
font-weight: 300;
	src: url('/fonts/open-sans-v15-latin-ext_latin-300.eot'); /* IE9 Compat Modes */
	src: local('Open Sans Light'), local('OpenSans-Light'),
		url('/fonts/open-sans-v15-latin-ext_latin-300.eot?#iefix') format('embedded-opentype'), /* IE6-IE8 */
		url('/fonts/open-sans-v15-latin-ext_latin-300.woff2') format('woff2'), /* Super Modern Browsers */
		url('/fonts/open-sans-v15-latin-ext_latin-300.woff') format('woff'), /* Modern Browsers */
		url('/fonts/open-sans-v15-latin-ext_latin-300.ttf') format('truetype'), /* Safari, Android, iOS */
		url('/fonts/open-sans-v15-latin-ext_latin-300.svg#OpenSans') format('svg'); /* Legacy iOS */
}

/* open-sans-700 - latin-ext_latin */
@font-face {
font-family: 'Open Sans';
font-style: normal;
font-weight: 700;
	src: url('/fonts/open-sans-v15-latin-ext_latin-700.eot'); /* IE9 Compat Modes */
	src: local('Open Sans Bold'), local('OpenSans-Bold'),
		url('/fonts/open-sans-v15-latin-ext_latin-700.eot?#iefix') format('embedded-opentype'), /* IE6-IE8 */
		url('/fonts/open-sans-v15-latin-ext_latin-700.woff2') format('woff2'), /* Super Modern Browsers */
		url('/fonts/open-sans-v15-latin-ext_latin-700.woff') format('woff'), /* Modern Browsers */
		url('/fonts/open-sans-v15-latin-ext_latin-700.ttf') format('truetype'), /* Safari, Android, iOS */
		url('/fonts/open-sans-v15-latin-ext_latin-700.svg#OpenSans') format('svg'); /* Legacy iOS */
}
body, span {
	margin: 0;
	padding: 0;
}
p, main {
	margin: 0;
	padding: 0;
	font-family: 'EB Garamond', serif;
	/* font-style: italic; */
}

h2 + div > p:first-child {
	margin-top: 5px;
}
.maincontent__h2 + a {
	display: block;
	margin-bottom: 8px;
}

body, span, h1, h2 {
	font-family: 'Open Sans', Arial, Helvetica, sans-serif;
	font-style: normal;
}
body {
	padding-top: 122px;
}
@media only screen and (max-width: 768px) {
	body{
		padding-top: 0;
	}
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
	margin-top: 50px;
}

.wrapperInner {
	margin: 0 auto;
	max-width: 800px;
}
p + p {
	text-indent: 20px;
}
`
}
