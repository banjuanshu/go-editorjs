package domain

type EditorJS struct {
	Blocks []EditorJSBlock `json:"blocks"`
}

type EditorJSBlock struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type EditorJSDataHeader struct {
	Text   string `json:"text,omitempty"`
	Level  int    `json:"level,omitempty"`
	Anchor string `json:"anchor,omitempty"`
}

type EditorJSDataParagraph struct {
	Text      string `json:"text,omitempty"`
	Alignment string `json:"alignment,omitempty"`
}

type EditorJSDataQuote struct {
	Text      string `json:"text,omitempty"`
	Caption   string `json:"caption,omitempty"`
	Alignment string `json:"alignment,omitempty"`
}

type EditorJSDataWarning struct {
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}

type EditorJSDataAlert struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

type EditorJSDataList struct {
	Style string        `json:"style,omitempty"`
	Items []interface{} `json:"items,omitempty"`
}

type EditorJSDataChecklist struct {
	Items []ChecklistItem `json:"items,omitempty"`
}

type EditorJSDataTable struct {
	WithHeadings bool       `json:"withHeadings,omitempty"`
	Content      [][]string `json:"content,omitempty"`
}

type EditorJSDataAnyButton struct {
	Link string `json:"link,omitempty"`
	Text string `json:"text,omitempty"`
}

type EditorJSDataCode struct {
	Code         string `json:"code,omitempty"`
	LanguageCode string `json:"languageCode,omitempty"`
}

type EditorJSDataRaw struct {
	Html string `json:"html,omitempty"`
}

type EditorJSDataImage struct {
	File           FileData `json:"file,omitempty"`
	URL            string   `json:"url,omitempty"`
	Caption        string   `json:"caption,omitempty"`
	WithBorder     bool     `json:"withBorder,omitempty"`
	WithBackground bool     `json:"withBackground,omitempty"`
	Stretched      bool     `json:"stretched,omitempty"`
}

type EditorJSDataLinkTool struct {
	Link string   `json:"link,omitempty"`
	Meta MetaData `json:"meta,omitempty"`
}

type EditorJSDataAttaches struct {
	File  FileData `json:"file,omitempty"`
	Title string   `json:"title,omitempty"`
}

type EditorJSDataEmbed struct {
	Service string `json:"service,omitempty"`
	Source  string `json:"source,omitempty"`
	Embed   string `json:"embed,omitempty"`
	Width   int    `json:"width,omitempty"`
	Height  int    `json:"height,omitempty"`
	Caption string `json:"caption,omitempty"`
}

type EditorJSDataImageGallery struct {
	URLs                []string `json:"urls,omitempty"`
	BkgMode             bool     `json:"bkgMode,omitempty"`
	LayoutDefault       bool     `json:"layoutDefault,omitempty"`
	LayoutHorizontal    bool     `json:"layoutHorizontal,omitempty"`
	LayoutSquare        bool     `json:"layoutSquare,omitempty"`
	LayoutWithGap       bool     `json:"layoutWithGap,omitempty"`
	LayoutWithFixedSize bool     `json:"layoutWithFixedSize,omitempty"`
}

type NestedListItem struct {
	Content string           `json:"content,omitempty"`
	Items   []NestedListItem `json:"items,omitempty"`
}

type ChecklistItem struct {
	Text    string `json:"text,omitempty"`
	Checked bool   `json:"checked,omitempty"`
}

type FileData struct {
	URL       string  `json:"url,omitempty"`
	Size      float64 `json:"size,omitempty"`
	Name      string  `json:"name,omitempty"`
	Extension string  `json:"extension,omitempty"`
}

type MetaData struct {
	Title       string        `json:"title,omitempty"`
	SiteName    string        `json:"site_name,omitempty"`
	Description string        `json:"description,omitempty"`
	Image       MetaDataImage `json:"image,omitempty"`
}

type MetaDataImage struct {
	URL string `json:"url,omitempty"`
}
