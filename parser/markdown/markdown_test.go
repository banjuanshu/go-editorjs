package markdown

import (
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support"
	"gitlab.com/rodrigoodhin/go-editorjs-parser/support/domain"
	"strconv"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderBlock(t *testing.T) {
	header1 := "# "

	for i := 1; i <= 6; i++ {
		level := strconv.Itoa(i)

		input1 := `{
    "blocks": [
        {
            "type": "header",
            "data": {
                "level": ` + level + `,
                "text": "Level ` + level + ` Header"
            }
        }
    ]
}`

		editorJSON1 := support.ParseEditorJSON(input1)
		content1 := support.PrepareData(editorJSON1.Blocks[0])

		expected1 := header1 + `Level ` + level + ` Header`
		actual1 := Header(content1.(*domain.EditorJSDataHeader))
		assert.Equal(t, expected1, actual1)

		header1 = "#" + header1
	}

	header2 := "# "

	for i := 1; i <= 6; i++ {
		level := strconv.Itoa(i)

		input1 := `{
    "blocks": [
        {
            "type": "header",
            "data": {
                "level": ` + level + `,
                "text": "Level ` + level + ` Header",
				"anchor": "Anchor Text ` + level + `"
            }
        }
    ]
}`

		editorJSON1 := support.ParseEditorJSON(input1)
		content1 := support.PrepareData(editorJSON1.Blocks[0])

		expected1 := header2 + `Level ` + level + ` Header`
		actual1 := Header(content1.(*domain.EditorJSDataHeader))
		assert.Equal(t, expected1, actual1)

		header2 = "#" + header2
	}
}

func TestParagraphBlock(t *testing.T) {
	input1 := `{
    "blocks": [
        {
			"type": "paragraph",
            "data": {
                "text": "I am a paragraph!"
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	content1 := support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `I am a paragraph!`
	actual1 := Paragraph(content1.(*domain.EditorJSDataParagraph))
	assert.Equal(t, expected1, actual1)

	input2 := `{
    "blocks": [
        {
			"type": "paragraph",
            "data": {
                "alignment": "center",
                "text": "I am a paragraph!"
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	content2 := support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `I am a paragraph!`
	actual2 := Paragraph(content2.(*domain.EditorJSDataParagraph))
	assert.Equal(t, expected2, actual2)
}

func TestQuoteBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
			"type": "quote",
            "data": {
                "alignment": "center",
                "caption": "Lao Tzu",
                "text": "The journey of a thousand miles begins with one step."
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	content := support.PrepareData(editorJSON.Blocks[0])

	expected := `> The journey of a thousand miles begins with one step.
>
> --- Lao Tzu`
	actual := Quote(content.(*domain.EditorJSDataQuote))
	assert.Equal(t, expected, actual)
}

func TestWarningBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
			"type": "warning",
            "data": {
                "message": "Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff.",
                "title": "Note:"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	content := support.PrepareData(editorJSON.Blocks[0])

	expected := `> Note:
>
> --- Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff.`
	actual := Warning(content.(*domain.EditorJSDataWarning))
	assert.Equal(t, expected, actual)
}

func TestDelimiterBlock(t *testing.T) {
	expected := `
***
`
	actual := Delimiter()
	assert.Equal(t, expected, actual)
}

func TestAlertBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
			"type": "alert",
            "data": {
                "message": "Something happened that you should know about.",
                "type": "primary"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	content := support.PrepareData(editorJSON.Blocks[0])

	expected := `
| Something happened that you should know about. |
| --- |`
	actual := Alert(content.(*domain.EditorJSDataAlert))
	assert.Equal(t, expected, actual)
}

func TestListBlock(t *testing.T) {
	input1 := `{
    "blocks": [
        {
			"type": "list",
            "data": {
                "items": [
                    "This is a block-styled editor",
                    "Clean output data",
                    "Simple and powerful API"
                ],
                "style": "unordered"
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	content1 := support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `- This is a block-styled editor
- Clean output data
- Simple and powerful API`
	actual1 := List(content1.(*domain.EditorJSDataList))
	assert.Equal(t, expected1, actual1)

	input2 := `{
    "blocks": [
        {
			"type": "list",
            "data": {
                "items": [
                    {
                        "content": "Cars",
                        "items": [
                            {
                                "content": "BMW",
                                "items": [
                                    {
                                        "content": "Z3",
                                        "items": []
                                    },
                                    {
                                        "content": "Z4",
                                        "items": []
                                    }
                                ]
                            },
                            {
                                "content": "Audi",
                                "items": [
                                    {
                                        "content": "A3",
                                        "items": []
                                    },
                                    {
                                        "content": "A1",
                                        "items": []
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "content": "Motorcycle",
                        "items": [
                            {
                                "content": "Ducati",
                                "items": [
                                    {
                                        "content": "916",
                                        "items": []
                                    }
                                ]
                            },
                            {
                                "content": "Yamanha",
                                "items": [
                                    {
                                        "content": "DT 180",
                                        "items": []
                                    }
                                ]
                            },
                            {
                                "content": "Honda",
                                "items": [
                                    {
                                        "content": "VFR 750R",
                                        "items": []
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "style": "ordered"
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	content2 := support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `1. Cars
    1. BMW
        1. Z3
        2. Z4
    2. Audi
        1. A3
        2. A1
2. Motorcycle
    1. Ducati
        1. 916
    2. Yamanha
        1. DT 180
    3. Honda
        1. VFR 750R`
	actual2 := List(content2.(*domain.EditorJSDataList))
	assert.Equal(t, expected2, actual2)
}

func TestChecklistBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "checklist",
            "data": {
                "items": [
                    {
                        "checked": true,
                        "text": "This is a block-styled editor"
                    },
                    {
                        "checked": false,
                        "text": "Clean output data"
                    },
                    {
                        "checked": true,
                        "text": "Simple and powerful API"
                    }
                ]
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	content := support.PrepareData(editorJSON.Blocks[0])

	expected := `- [x] This is a block-styled editor
- [ ] Clean output data
- [x] Simple and powerful API`
	actual := Checklist(content.(*domain.EditorJSDataChecklist))
	assert.Equal(t, expected, actual)
}

func TestTableBlock(t *testing.T) {
	input1 := `{
    "blocks": [
        {
            "type": "table",
            "data": {
                "content": [
                    [
                        "Kine",
                        "Pigs",
                        "Chicken"
                    ],
                    [
                        "1 pcs",
                        "3 pcs",
                        "12 pcs"
                    ],
                    [
                        "100$",
                        "200$",
                        "150$"
                    ]
                ],
                "withHeadings": true
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	content1 := support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `| Kine | Pigs | Chicken |
|---|---|---|
| 1 pcs | 3 pcs | 12 pcs |
| 100$ | 200$ | 150$ |`
	actual1 := Table(content1.(*domain.EditorJSDataTable))
	assert.Equal(t, expected1, actual1)

	input2 := `{
    "blocks": [
        {
            "type": "table",
            "data": {
                "content": [
                    [
                        "Kine",
                        "1 pcs",
                        "100$"
                    ],
                    [
                        "Pigs",
                        "3 pcs",
                        "200$"
                    ],
                    [
                        "Chickens",
                        "12 pcs",
                        "150$"
                    ]
                ]
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	content2 := support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `| | | |
|---|---|---|
| Kine | 1 pcs | 100$ |
| Pigs | 3 pcs | 200$ |
| Chickens | 12 pcs | 150$ |`
	actual2 := Table(content2.(*domain.EditorJSDataTable))
	assert.Equal(t, expected2, actual2)
}

func TestAnyButtonBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "AnyButton",
            "data": {
                "link": "https://editorjs.io/",
                "text": "editorjs official"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	content := support.PrepareData(editorJSON.Blocks[0])

	expected := `[editorjs official](https://editorjs.io/)`
	actual := AnyButton(content.(*domain.EditorJSDataAnyButton))
	assert.Equal(t, expected, actual)
}

func TestCodeBlock(t *testing.T) {
	input1 := `{
    "blocks": [
        {
            "type": "code",
            "data": {
                "code": "body {\n font-size: 14px;\n line-height: 16px;\n}"
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	content1 := support.PrepareData(editorJSON1.Blocks[0])

	expected1 := "```\nbody {\n font-size: 14px;\n line-height: 16px;\n}\n```"
	actual1 := Code(content1.(*domain.EditorJSDataCode))
	assert.Equal(t, expected1, actual1)

	input2 := `{
    "blocks": [
        {
            "type": "code",
            "data": {
                "code": "body {\n font-size: 14px;\n line-height: 16px;\n}",
                "languageCode": "css"
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	content2 := support.PrepareData(editorJSON2.Blocks[0])

	expected2 := "```css\nbody {\n font-size: 14px;\n line-height: 16px;\n}\n```"
	actual2 := Code(content2.(*domain.EditorJSDataCode))
	assert.Equal(t, expected2, actual2)
}

func TestRawBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "raw",
            "data": {
                "html": "<div style=\"background: #000; color: #fff; font-size: 30px; padding: 50px;\">Any HTML code</div>"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	content := support.PrepareData(editorJSON.Blocks[0])

	expected := "```\n<div style=\"background: #000; color: #fff; font-size: 30px; padding: 50px;\">Any HTML code</div>\n```"
	actual := Raw(content.(*domain.EditorJSDataRaw))
	assert.Equal(t, expected, actual)
}

func TestImageBlock(t *testing.T) {
	input1 := `{
    "blocks": [
        {
            "type": "image",
            "data": {
                "caption": "Roadster // tesla.com",
                "file": {
                    "url": "https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg"
                },
                "stretched": false,
                "withBackground": true,
                "withBorder": true
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	content1 := support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `![Roadster // tesla.com](https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg)`
	actual1 := Image(content1.(*domain.EditorJSDataImage))
	assert.Equal(t, expected1, actual1)

	input2 := `{
    "blocks": [
        {
            "type": "image",
            "data": {
                "caption": "Roadster // tesla.com",
                "stretched": true,
                "url": "https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg",
                "withBackground": false,
                "withBorder": false
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	content2 := support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `![Roadster // tesla.com](https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg)`
	actual2 := Image(content2.(*domain.EditorJSDataImage))
	assert.Equal(t, expected2, actual2)
}

func TestLinkToolBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "linkTool",
            "data": {
                "link": "https://codex.so",
                "meta": {
                    "description": "Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.",
                    "image": {
                        "url": "https://codex.so/public/app/img/meta_img.png"
                    },
                    "site_name": "CodeX",
                    "title": "CodeX Team"
                }
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	content := support.PrepareData(editorJSON.Blocks[0])

	expected := `---

# CodeX Team

![CodeX Team](https://codex.so/public/app/img/meta_img.png)

*Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.*

[codex.so](https://codex.so)

---`
	actual := LinkTool(content.(*domain.EditorJSDataLinkTool))
	assert.Equal(t, expected, actual)
}

func TestAttachesBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "attaches",
            "data": {
                "file": {
                    "extension": "jpg",
                    "name": "hero.jpg",
                    "size": 260096,
                    "url": "https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg"
                },
                "title": "Hero"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	content := support.PrepareData(editorJSON.Blocks[0])

	expected := `---

### hero.jpg

###### 254 KiB

[Download](https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg)

---`
	actual := Attaches(content.(*domain.EditorJSDataAttaches))
	assert.Equal(t, expected, actual)
}

func TestEmbedBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "embed",
            "data": {
                "caption": "Lamborghini Aventador SVJ",
                "embed": "https://www.youtube.com/embed/viW44cUfxCE",
                "height": 315,
                "service": "Youtube",
                "source": "https://www.youtube.com/watch?v=viW44cUfxCE",
                "width": 560
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	content := support.PrepareData(editorJSON.Blocks[0])

	expected := `---

### Lamborghini Aventador SVJ

[Watch on Youtube](https://www.youtube.com/watch?v=viW44cUfxCE)

---`
	actual := Embed(content.(*domain.EditorJSDataEmbed))
	assert.Equal(t, expected, actual)
}

func TestImageGalleryBlock(t *testing.T) {
	input := `{
    "blocks": [
        {
            "type": "imageGallery",
            "data": {
                "bkgMode": false,
                "editImages": true,
                "layoutDefault": true,
                "layoutHorizontal": false,
                "layoutSquare": false,
                "layoutWithFixedSize": false,
                "layoutWithGap": false,
                "urls": [
                    "https://www.nawpic.com/media/2020/ocean-nawpic-15.jpg",
                    "https://www.nawpic.com/media/2020/ocean-nawpic-18.jpg",
                    "https://wallpapercave.com/wp/6L4TVMP.jpg",
                    "https://wallpapercave.com/wp/wp9810772.jpg",
                    "https://wallpapercave.com/wp/wp9121482.jpg",
                    "https://wallpapercave.com/wp/wp9100484.jpg",
                    "https://cdn.wallpapersafari.com/94/22/4H3mFp.jpg"
                ]
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	content := support.PrepareData(editorJSON.Blocks[0])

	expected := `---
![Image 0](https://www.nawpic.com/media/2020/ocean-nawpic-15.jpg)
![Image 1](https://www.nawpic.com/media/2020/ocean-nawpic-18.jpg)
![Image 2](https://wallpapercave.com/wp/6L4TVMP.jpg)
![Image 3](https://wallpapercave.com/wp/wp9810772.jpg)
![Image 4](https://wallpapercave.com/wp/wp9121482.jpg)
![Image 5](https://wallpapercave.com/wp/wp9100484.jpg)
![Image 6](https://cdn.wallpapersafari.com/94/22/4H3mFp.jpg)
---`
	actual := ImageGallery(content.(*domain.EditorJSDataImageGallery))
	assert.Equal(t, expected, actual)
}
