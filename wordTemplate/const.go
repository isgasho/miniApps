package main

//SelfTags are Self Closing Tags
type SelfTags byte

//Const fields of type SelfTags
const (
	FieldCurrentPage   SelfTags = 1
	FieldNumberofPages SelfTags = 2
	PageBreak          SelfTags = 3
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
	Underline           StyleTags = iota
	Emphasis            StyleTags = iota
	Font                StyleTags = iota
	Bold                StyleTags = 1 << iota
	Italic              StyleTags = 1 << iota
	Caps                StyleTags = 1 << iota
	SmallCaps           StyleTags = 1 << iota
	StrikeThrough       StyleTags = 1 << iota
	DoubleStrikeThrough StyleTags = 1 << iota
	Outline             StyleTags = 1 << iota
	Shadow              StyleTags = 1 << iota
	Emboss              StyleTags = 1 << iota
	Imprint             StyleTags = 1 << iota
	NoProof             StyleTags = 1 << iota
	SnapToGrid          StyleTags = 1 << iota
	Vanish              StyleTags = 1 << iota
	WebHidden           StyleTags = 1 << iota
	RightToLeft         StyleTags = 1 << iota
	SubScript           StyleTags = 1 << iota
	SuperScript         StyleTags = 1 << iota
)

var WhiteListSelfTags = map[string]SelfTags{}

func initWhiteListSelfTags() {
	WhiteListSelfTags["fieldcurrentpage"] = FieldCurrentPage
	WhiteListSelfTags["fieldnumberofpages"] = FieldNumberofPages
	WhiteListSelfTags["pagebreak"] = PageBreak
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

func initConstMap() {
	initWhileListTags()
	initWhiteListSelfTags()
	initWhiteListStyleMap()
}
