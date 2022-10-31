package sample

import (
	"github.com/banjuanshu/go-editorjs/support"
	"github.com/banjuanshu/go-editorjs/support/domain"
	"github.com/matryer/is"
	"strconv"

	"testing"
)

var obj = Init(true)

func TestHeaderBlock(t *testing.T) {
	is := is.New(t)

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
		obj.Data = support.PrepareData(editorJSON1.Blocks[0]).(*domain.EditorJSDataHeader)

		expected1 := `<h` + level + `  class="">Level ` + level + ` Header</h` + level + `>`

		obj.Header()

		actual1 := obj.Result[i-1]

		is.Equal(expected1, actual1) // Header 1 is different from expected
	}

	obj.Result = []string{}

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
		obj.Data = support.PrepareData(editorJSON1.Blocks[0])

		expected1 := `<h` + level + ` id="anchor-text-` + level + `" class="">Level ` + level + ` Header</h` + level + `>`

		obj.Header()

		actual1 := obj.Result[i-1]

		is.Equal(expected1, actual1) // Header 2 is different from expected
	}
}

func TestParagraphBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<p class=" ">I am a paragraph!</p>`

	obj.Paragraph()

	actual1 := obj.Result[0]

	is.Equal(expected1, actual1) // Paragraph 1 is different from expected

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<p class=" alignment_text_center">I am a paragraph!</p>`

	obj.Paragraph()

	actual2 := obj.Result[0]

	is.Equal(expected2, actual2) // Paragraph 2 is different from expected
}

func TestQuoteBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<figure class="quote_figure alignment_text_center">
<blockquote class="quote_blockquote">
The journey of a thousand miles begins with one step.
</blockquote>
<figcaption class="quote_figcaption">
Lao Tzu
</figcaption>
</figure>`

	obj.Quote()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Quote is different from expected
}

func TestWarningBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="warning_msg">
<b>
Note:
</b>
Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff.
</div>`

	obj.Warning()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Warning is different from expected
}

func TestDelimiterBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	expected := `<div class="delimiter_block">***</div>`

	obj.Delimiter()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Delimiter is different from expected
}

func TestAlertBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="alert_box alert_primary">
Something happened that you should know about.
</div>`

	obj.Alert()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Alert is different from expected
}

func TestListBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<ul class="list_group">
<li class="list_item">This is a block-styled editor</li>
<li class="list_item">Clean output data</li>
<li class="list_item">Simple and powerful API</li>
</ul>`

	obj.List()

	actual1 := obj.Result[0]

	is.Equal(expected1, actual1) // List 1 is different from expected

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<ol class="list_group">
<li class="list_item">Cars</li>
<ol class="">
<li class="list_item">BMW</li>
<ol class="">
<li class="list_item">Z3</li>
</li>
<li class="list_item">Z4</li>
</li>
</ol>
</li>
<li class="list_item">Audi</li>
<ol class="">
<li class="list_item">A3</li>
</li>
<li class="list_item">A1</li>
</li>
</ol>
</li>
</ol>
</li>
<li class="list_item">Motorcycle</li>
<ol class="">
<li class="list_item">Ducati</li>
<ol class="">
<li class="list_item">916</li>
</li>
</ol>
</li>
<li class="list_item">Yamanha</li>
<ol class="">
<li class="list_item">DT 180</li>
</li>
</ol>
</li>
<li class="list_item">Honda</li>
<ol class="">
<li class="list_item">VFR 750R</li>
</li>
</ol>
</li>
</ol>
</li>
</ol>`

	obj.List()

	actual2 := obj.Result[0]

	is.Equal(expected2, actual2) // List 2 is different from expected
}

func TestChecklistBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="checklist_block">
<div class="checklist_item">
<span class="checklist_item_checkbox checklist_checked">&#10004;</span>
<span class="checklist_item_text">This is a block-styled editor</span>
</div>
<div class="checklist_item">
<span class="checklist_item_checkbox">&nbsp;-&nbsp;</span>
<span class="checklist_item_text">Clean output data</span>
</div>
<div class="checklist_item">
<span class="checklist_item_checkbox checklist_checked">&#10004;</span>
<span class="checklist_item_text">Simple and powerful API</span>
</div>
</div>`

	obj.Checklist()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Checklist is different from expected
}

func TestTableBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<table class="table_block">
<tr class="table_tr">
<th class="table_th">Kine</th>
<th class="table_th">Pigs</th>
<th class="table_th">Chicken</th>
</tr>
<tr class="table_tr">
<td class="table_td">1 pcs</td>
<td class="table_td">3 pcs</td>
<td class="table_td">12 pcs</td>
</tr>
<tr class="table_tr">
<td class="table_td">100$</td>
<td class="table_td">200$</td>
<td class="table_td">150$</td>
</tr>
</table>`

	obj.Table()

	actual1 := obj.Result[0]

	is.Equal(expected1, actual1) // Table 1 is different from expected

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<table class="table_block">
<tr class="table_tr">
<td class="table_td">Kine</td>
<td class="table_td">1 pcs</td>
<td class="table_td">100$</td>
</tr>
<tr class="table_tr">
<td class="table_td">Pigs</td>
<td class="table_td">3 pcs</td>
<td class="table_td">200$</td>
</tr>
<tr class="table_tr">
<td class="table_td">Chickens</td>
<td class="table_td">12 pcs</td>
<td class="table_td">150$</td>
</tr>
</table>`

	obj.Table()

	actual2 := obj.Result[0]

	is.Equal(expected2, actual2) // Table 2 is different from expected
}

func TestAnyButtonBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<a class="anyButton" href="https://editorjs.io/">editorjs official</a>`

	obj.AnyButton()

	actual := obj.Result[0]

	is.Equal(expected, actual) // AnyButton is different from expected
}

func TestCodeBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<pre class="code_pre">
<code class="code_block">body {
 font-size: 14px;
 line-height: 16px;
}
</code></pre>`

	obj.Code()

	actual1 := obj.Result[0]

	is.Equal(expected1, actual1) // Code 1 is different from expected

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<pre class="code_pre">
<code class="code_block">body {
 font-size: 14px;
 line-height: 16px;
}
</code></pre>`

	obj.Code()

	actual2 := obj.Result[0]

	is.Equal(expected2, actual2) // Code 2 is different from expected
}

func TestRawBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<pre class="raw_pre">
<code class="raw_block">&lt;div style="background: #000; color: #fff; font-size: 30px; padding: 50px;"&gt;Any HTML code&lt;/div&gt;
</code></pre>`

	obj.Raw()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Raw is different from expected
}

func TestImageBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input1 := `{
    "blocks": [
        {
		  "type" : "image",
		  "data" : {
			"url" : "https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg",
			"caption" : "Mountain",
			"withBorder" : true,
			"withBackground" : true,
			"stretched" : true
		  }
		}
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<div class="image image_with_background" ><img class="image image_with_border image image_stretched" src="https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg" alt="Mountain" title="Mountain" /></div>`

	obj.Image()

	actual1 := obj.Result[0]

	is.Equal(expected1, actual1) // Image 1 is different from expected

	obj.Result = []string{}

	input2 := `{
    "blocks": [
        {
		  "type" : "image",
		  "data" : {
			"file": {
			  "url" : "https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg"
			},
			"caption" : "Mountain",
			"withBorder" : false,
			"withBackground" : true,
			"stretched" : false
		  }
		}
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<div class="image image_with_background" ><img class="" src="https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg" alt="Mountain" title="Mountain" /></div>`

	obj.Image()

	actual2 := obj.Result[0]

	is.Equal(expected2, actual2) // Image 2 is different from expected
}

func TestLinkToolBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input := `{
    "blocks": [
        {
		  "type" : "linkTool",
		  "data" : {
			"link" : "https://codex.so",
			"meta" : {
			  "title" : "CodeX Team",
			  "site_name" : "CodeX",
			  "description" : "Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.",
			  "image" : {
				"url" : "https://pbs.twimg.com/profile_images/993612654861344768/wMPEM5XW_400x400.jpg"
			  }
			}
		  }
		}
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<a href="https://codex.so" target="_Blank" rel="nofollow noindex noreferrer" class="">
<div class="linkTool_content">
<div class="">
<div class="linkTool_left_side">
<div class="linkTool_title">
CodeX Team
</div>
<div class="linkTool_description">
Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.
</div>
<div class="linkTool_anchor">
codex.so
</div>
</div>
<div class="linkTool_image_block">
<img class="linkTool_image" src="https://pbs.twimg.com/profile_images/993612654861344768/wMPEM5XW_400x400.jpg" />
</div>
</div>
</div>
</a>`

	obj.LinkTool()

	actual := obj.Result[0]

	is.Equal(expected, actual) // LinkTool is different from expected
}

func TestAttachesBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<a href="https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg" rel="noopener noreferrer" target="_blank" class="">
<div class="attaches_content">
<div class="" >
<div class="attaches_left" >
<img class="attaches_image" src="https://i.ibb.co/K7Myr2k/file-icon.png" />
</div>
<div class="attaches_center">
<div class="attaches_filename">
hero.jpg
</div>
<div class="attaches_size">
254 KiB
</div>
</div>
<div class="attaches_right" >
<img class="attaches_image" src="https://i.ibb.co/VYyHr6C/download-icon.png" />
</div>
</div>
</div>
</a>`

	obj.Attaches()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Attaches is different from expected
}

func TestEmbedBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="embed_block" style="max-width: 560px">
<div class="embed_title">Lamborghini Aventador SVJ</div>
<iframe width="560" height="315" src="https://www.youtube.com/embed/viW44cUfxCE" title="Lamborghini Aventador SVJ" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
<div class="embed_bottom">
<a class="embed_link" href="https://www.youtube.com/watch?v=viW44cUfxCE" target="_Blank">Watch on Youtube</a>
</div>
</div>`

	obj.Embed()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Embed is different from expected
}

func TestImageGalleryBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

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
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="gg-container">
<div class="gg-box" id="">
<img src="https://www.nawpic.com/media/2020/ocean-nawpic-15.jpg" id="gg-image-0" />
<img src="https://www.nawpic.com/media/2020/ocean-nawpic-18.jpg" id="gg-image-1" />
<img src="https://wallpapercave.com/wp/6L4TVMP.jpg" id="gg-image-2" />
<img src="https://wallpapercave.com/wp/wp9810772.jpg" id="gg-image-3" />
<img src="https://wallpapercave.com/wp/wp9121482.jpg" id="gg-image-4" />
<img src="https://wallpapercave.com/wp/wp9100484.jpg" id="gg-image-5" />
<img src="https://cdn.wallpapersafari.com/94/22/4H3mFp.jpg" id="gg-image-6" />
</div>
</div>`

	obj.ImageGallery()

	actual := obj.Result[0]

	is.Equal(expected, actual) // ImageGallery is different from expected
}
