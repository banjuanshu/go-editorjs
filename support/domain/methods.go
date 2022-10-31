package domain

type EditorJSMethods interface {
	SetData(data interface{})
	SetStyles(styles []string)
	SetResult(result string)
	SetScripts(scripts []string)
	LoadLibrary()
	CreatePage() string
	GetHtml() string
	Separator()
	Header()
	Paragraph()
	Quote()
	Warning()
	Delimiter()
	Alert()
	List()
	Checklist()
	Table()
	AnyButton()
	Code()
	Raw()
	Image()
	LinkTool()
	Attaches()
	Embed()
	ImageGallery()
}
