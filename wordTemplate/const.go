package main

//SelfTags are Self Closing Tags
type SelfTags byte

//Const fields of type SelfTags
const (
	FieldCurrentPage   SelfTags = 1
	FieldNumberofPages SelfTags = 2
	PageBreak          SelfTags = 3
	LineBreak          SelfTags = 4
	InlineImage        SelfTags = 5
	AnchorImage        SelfTags = 6
	ParaTop            SelfTags = 7
	ParaBottom         SelfTags = 8
	ParaLeft           SelfTags = 9
	ParaRight          SelfTags = 10
)

//Tags are the tags with childrens
type Tags byte

//Const field of type Tags
const (
	Document        Tags = 1
	Body            Tags = 2
	Center          Tags = 3
	Left            Tags = 4
	Right           Tags = 5
	PageHeader      Tags = 6
	PageFooter      Tags = 7
	DocProps        Tags = 8
	Title           Tags = 9
	Author          Tags = 10
	Description     Tags = 11
	Category        Tags = 12
	Version         Tags = 13
	Application     Tags = 14
	Company         Tags = 15
	Paragraph       Tags = 16
	ParagraphBorder Tags = 17
)

type StyleTags int

const (
	TextShading         StyleTags = -6 //style={0-38},color={hex},fill={hex}
	TextBorder          StyleTags = -5 //style={0-193},color={hex},frame={true/false},shadow={true/false},size={INT},space={INT}
	TextEffect          StyleTags = -4 //style={0-7} //not working
	TextHighlight       StyleTags = -3 //style={0-17}
	Underline           StyleTags = -2 //style={0-18},color={hex}
	Emphasis            StyleTags = -1 //style={0-5}
	Font                StyleTags = 0  //family={string},size={float},kern={float},charSpacing={float},color={hex},scale={string},csize={string}
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

var WhiteListSelfTags = map[string]SelfTags{}

func initWhiteListSelfTags() {
	WhiteListSelfTags["fieldcurrentpage"] = FieldCurrentPage
	WhiteListSelfTags["fieldnumberofpages"] = FieldNumberofPages
	WhiteListSelfTags["pagebreak"] = PageBreak
	WhiteListSelfTags["linebreak"] = LineBreak
	WhiteListSelfTags["inlineImg"] = InlineImage
	WhiteListSelfTags["anchorImg"] = AnchorImage
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
	WhiteListStyleTags["texthighlight"] = TextHighlight
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
	WhiteListTags["paragraph"] = Paragraph
	WhiteListTags["paragraphborder"] = ParagraphBorder
}

func initConstMap() {
	initWhileListTags()
	initWhiteListSelfTags()
	initWhiteListStyleMap()
}
