package bootstrap

import (
	"github.com/banjuanshu/go-editorjs/support"
	"github.com/banjuanshu/go-editorjs/support/domain"
	"strconv"

	"testing"

	"github.com/matryer/is"
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

	expected2 := `<p class=" text-center">I am a paragraph!</p>`

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

	expected := `<figure class=" text-center">
<blockquote class="blockquote">
The journey of a thousand miles begins with one step.
</blockquote>
<figcaption class="blockquote-footer">
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

	expected := `<div class="alert alert-warning">
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

	expected := `<div class="alert alert-light text-center">***</div>`

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

	expected := `<div class="alert alert-primary">
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

	expected1 := `<ul class="list-group">
<li class="list-group-item">This is a block-styled editor</li>
<li class="list-group-item">Clean output data</li>
<li class="list-group-item">Simple and powerful API</li>
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

	expected2 := `<ol class="list-group">
<li class="list-group-item">Cars</li>
<ol class="">
<li class="list-group-item">BMW</li>
<ol class="">
<li class="list-group-item">Z3</li>
</li>
<li class="list-group-item">Z4</li>
</li>
</ol>
</li>
<li class="list-group-item">Audi</li>
<ol class="">
<li class="list-group-item">A3</li>
</li>
<li class="list-group-item">A1</li>
</li>
</ol>
</li>
</ol>
</li>
<li class="list-group-item">Motorcycle</li>
<ol class="">
<li class="list-group-item">Ducati</li>
<ol class="">
<li class="list-group-item">916</li>
</li>
</ol>
</li>
<li class="list-group-item">Yamanha</li>
<ol class="">
<li class="list-group-item">DT 180</li>
</li>
</ol>
</li>
<li class="list-group-item">Honda</li>
<ol class="">
<li class="list-group-item">VFR 750R</li>
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

	expected := `<div class="">
<div class="">
<span class="badge rounded-pill bg-light">&#10004;</span>
<span class="">This is a block-styled editor</span>
</div>
<div class="">
<span class="badge rounded-pill bg-light">&nbsp;-&nbsp;</span>
<span class="">Clean output data</span>
</div>
<div class="">
<span class="badge rounded-pill bg-light">&#10004;</span>
<span class="">Simple and powerful API</span>
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

	expected1 := `<table class="table table-striped">
<tr class="">
<th class="">Kine</th>
<th class="">Pigs</th>
<th class="">Chicken</th>
</tr>
<tr class="">
<td class="">1 pcs</td>
<td class="">3 pcs</td>
<td class="">12 pcs</td>
</tr>
<tr class="">
<td class="">100$</td>
<td class="">200$</td>
<td class="">150$</td>
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

	expected2 := `<table class="table table-striped">
<tr class="">
<td class="">Kine</td>
<td class="">1 pcs</td>
<td class="">100$</td>
</tr>
<tr class="">
<td class="">Pigs</td>
<td class="">3 pcs</td>
<td class="">200$</td>
</tr>
<tr class="">
<td class="">Chickens</td>
<td class="">12 pcs</td>
<td class="">150$</td>
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

	expected := `<a class="btn btn-secondary" href="https://editorjs.io/">editorjs official</a>`

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

	expected1 := `<pre class="p-3 mb-2 bg-light">
<code class="text-dark">body {
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

	expected2 := `<pre class="p-3 mb-2 bg-light">
<code class="text-dark">body {
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

	expected := `<pre class="p-3 mb-2">
<code class="text-dark">&lt;div style="background: #000; color: #fff; font-size: 30px; padding: 50px;"&gt;Any HTML code&lt;/div&gt;
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

	expected1 := `<div class="bg-warning p-5" ><img class="border img-fluid" src="https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg" alt="Mountain" title="Mountain" /></div>`

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

	expected2 := `<div class="bg-warning p-5" ><img class="" src="https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg" alt="Mountain" title="Mountain" /></div>`

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

	expected := `<a href="https://codex.so" target="_Blank" rel="nofollow noindex noreferrer" class="text-decoration-none">
<div class="container m-3 bg-light border">
<div class="row">
<div class="col-9 col-sm-10 col-md-11 d-grid gap-3">
<div class="p-1 link-dark">
CodeX Team
</div>
<div class="p-1 link-success">
Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.
</div>
<div class="p-1 link-secondary">
codex.so
</div>
</div>
<div class="col-3 col-sm-2 col-md-1 p-2">
<img class="img-thumbnail" src="https://pbs.twimg.com/profile_images/993612654861344768/wMPEM5XW_400x400.jpg" />
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

	expected := `<a href="https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg" rel="noopener noreferrer" target="_blank" class="text-decoration-none">
<div class="container m-3 bg-light border">
<div class="row" >
<div class="col-3 col-sm-2 col-md-1 p-2" >
<img class="img-thumbnail bg-transparent border-0" src="https://i.ibb.co/K7Myr2k/file-icon.png" />
</div>
<div class="col-6 col-sm-8 col-md-10 d-grid">
<div class="p-1 link-dark">
hero.jpg
</div>
<div class="p-1 link-secondary">
254 KiB
</div>
</div>
<div class="col-3 col-sm-2 col-md-1 p-4" >
<img class="img-thumbnail bg-transparent border-0" src="https://i.ibb.co/VYyHr6C/download-icon.png" />
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

	expected := `<div class="d-grid m-5" style="max-width: 560px">
<div class="bg-dark p-1 text-light">Lamborghini Aventador SVJ</div>
<iframe width="560" height="315" src="https://www.youtube.com/embed/viW44cUfxCE" title="Lamborghini Aventador SVJ" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
<div class="d-grid justify-content-md-end bg-light p-1">
<a class="text-decoration-none" href="https://www.youtube.com/watch?v=viW44cUfxCE" target="_Blank">Watch on Youtube</a>
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
