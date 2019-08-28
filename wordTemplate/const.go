package main

//SelfTags are Self Closing Tags
type SelfTags byte

//Const fields of type SelfTags
const (
	FieldCurrentPage   SelfTags = 1
	FieldNumberofPages SelfTags = 2
	PageBreak          SelfTags = 3
	LineBreak          SelfTags = 4
)

//Tags are the tags with childrens
type Tags byte

//Const field of type Tags
const (
	Document    Tags = 1
	Body        Tags = 2
	Center      Tags = 3
	Left        Tags = 4
	Right       Tags = 5
	PageHeader  Tags = 6
	PageFooter  Tags = 7
	DocProps    Tags = 8
	Title       Tags = 9
	Author      Tags = 10
	Description Tags = 11
	Category    Tags = 12
	Version     Tags = 13
	Application Tags = 14
	Company     Tags = 15
)

type StyleTags int

const (
	TextShading         StyleTags = -6 //style,color,fill
	TextBorder          StyleTags = -5 //style,color,frame,shadow,size,space
	TextEffect          StyleTags = -4 //style
	TextHighlight       StyleTags = -3 //style
	Underline           StyleTags = -2 //style,color
	Emphasis            StyleTags = -1 //style
	Font                StyleTags = 0  //family,size,kern,charSpacing,color,scale,csize
	Bold                StyleTags = 1 << 1
	Italic              StyleTags = 1 << 2
	Caps                StyleTags = 1 << 3
	SmallCaps           StyleTags = 1 << 4
	StrikeThrough       StyleTags = 1 << 5
	DoubleStrikeThrough StyleTags = 1 << 6
	Outline             StyleTags = 1 << 7
	Shadow              StyleTags = 1 << 8
	Emboss              StyleTags = 1 << 9
	Imprint             StyleTags = 1 << 10
	NoProof             StyleTags = 1 << 11
	SnapToGrid          StyleTags = 1 << 12
	Vanish              StyleTags = 1 << 13
	WebHidden           StyleTags = 1 << 14
	RightToLeft         StyleTags = 1 << 15
	SubScript           StyleTags = 1 << 16
	SuperScript         StyleTags = 1 << 17
)

type Attrib byte

var WhiteListSelfTags = map[string]SelfTags{}

func initWhiteListSelfTags() {
	WhiteListSelfTags["fieldcurrentpage"] = FieldCurrentPage
	WhiteListSelfTags["fieldnumberofpages"] = FieldNumberofPages
	WhiteListSelfTags["pagebreak"] = PageBreak
	WhiteListSelfTags["linebreak"] = LineBreak
}

var WhiteListStyleTags = map[string]StyleTags{}

func initWhiteListStyleMap() {
	WhiteListStyleTags["b"] = Bold
	WhiteListStyleTags["i"] = Italic
	WhiteListStyleTags["caps"] = Caps
	WhiteListStyleTags["smallcaps"] = SmallCaps
	WhiteListStyleTags["strike"] = StrikeThrough
	WhiteListStyleTags["dstrike"] = DoubleStrikeThrough
	WhiteListStyleTags["outline"] = Outline
	WhiteListStyleTags["shadow"] = Shadow
	WhiteListStyleTags["emboss"] = Emboss
	WhiteListStyleTags["imprint"] = Imprint
	WhiteListStyleTags["noproof"] = NoProof
	WhiteListStyleTags["snaptogrid"] = SnapToGrid
	WhiteListStyleTags["vanish"] = Vanish
	WhiteListStyleTags["webhidden"] = WebHidden
	WhiteListStyleTags["subscript"] = SubScript
	WhiteListStyleTags["superscript"] = SuperScript
	WhiteListStyleTags["rtl"] = RightToLeft
	WhiteListStyleTags["u"] = Underline
	WhiteListStyleTags["em"] = Emphasis
	WhiteListStyleTags["font"] = Font
	WhiteListStyleTags["highlight"] = TextHighlight
	WhiteListStyleTags["texteffect"] = TextEffect
	WhiteListStyleTags["textborder"] = TextBorder
	WhiteListStyleTags["textshading"] = TextShading
}

var WhiteListTags = map[string]Tags{}

func initWhileListTags() {
	WhiteListTags["document"] = Document
	WhiteListTags["body"] = Body
	WhiteListTags["center"] = Center
	WhiteListTags["left"] = Left
	WhiteListTags["right"] = Right
	WhiteListTags["pageheader"] = PageHeader
	WhiteListTags["pagefooter"] = PageFooter
	WhiteListTags["docprops"] = DocProps
	WhiteListTags["title"] = Title
	WhiteListTags["author"] = Author
	WhiteListTags["description"] = Description
	WhiteListTags["category"] = Category
	WhiteListTags["version"] = Version
	WhiteListTags["application"] = Application
	WhiteListTags["company"] = Company
}

var WhiteListAttribs = map[string]Attrib{}

func initConstMap() {
	initWhileListTags()
	initWhiteListSelfTags()
	initWhiteListStyleMap()
}
