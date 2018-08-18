package staticPresentation

// Creates a new GlobalCssComponent
func NewGlobalCssComponent() *GlobalCssComponent {
	return new(GlobalCssComponent)
}

type GlobalCssComponent struct {
	abstractComponent
}

func (gcc *GlobalCssComponent) GetCss() string {
	return `
body, p, span {
	margin: 0;
	padding: 0;
	font-family: Arial, Helvetica, sans-serif;
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
