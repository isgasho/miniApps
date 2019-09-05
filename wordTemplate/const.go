package main

//SelfTags are Self Closing Tags
type SelfTags byte

//Const fields of type SelfTags
const (
	FieldCurrentPage   SelfTags = 1
	FieldNumberofPages SelfTags = 2
	PageBreak          SelfTags = 3
	LineBreak          SelfTags = 4
	//InlineImage props: src={string}
	InlineImage SelfTags = 5
	//AnchorImage props: src={string} noWrap={NONE} name={string}
	//hAlign={0-5} vAlign={0-5} xOffset={num} yOffset={num}
	//wrap={0-4} size={Num[width,height]}
	//origin={Num[Horizontal:{0-8},Vertical{0-8}]}
	AnchorImage SelfTags = 6
	//BorderTop props: style={0-193} size={int} space={int}
	//color={hex} frame={true} shadow={true}
	BorderTop SelfTags = 7
	//BorderBottom props: style={0-193} size={int} space={int}
	//color={hex} frame={true} shadow={true}
	BorderBottom SelfTags = 8
	//BorderLeft props: style={0-193} size={int} space={int}
	//color={hex} frame={true} shadow={true}
	BorderLeft SelfTags = 9
	//BorderRight props: style={0-193} size={int} space={int}
	//color={hex} frame={true} shadow={true}
	BorderRight SelfTags = 10
	//ParaShading props: style={0-38} color={hex} fill={hex}
	ParaShading SelfTags = 11
	//ParaAlignment props: style={0-12}
	ParaAlignment SelfTags = 12
	//ParaText props: align={0-5} direction={0-12}
	ParaText SelfTags = 13
	//ParaFrame props: dropCaps={0-3} lines={int} wrap={0-6}
	//hAnchor={0-3} vAnchor={0-3} xAlign={0-5} yAlign={0-6}
	//hRule={0-3} height={num} width={num} hSpace={num} vSpace={num}
	//x={num} y={num}
	ParaFrame SelfTags = 14
	//ParaIndent props: start={int} end={int} hang={int} first={int}
	ParaIndent SelfTags = 15
	//ParaTextBoxTightWrap props: style={0-5}
	ParaTextBoxTightWrap SelfTags = 16
	//ParaSpacing props: lineSpacing={num,{0-3}} after={num} before={num}
	//autoAfter={bool} autoBefore={bool}
	ParaSpacing SelfTags = 17
	//DocPageBackground props: color={hex} theme={0-17}
	//tint={0-255 HexVal without hash}
	//shade={0-255 HexVal without hash}
	DocBackground SelfTags = 18
	//DocPageSize props: height={string} width={string}
	//orientation={0-2}
	DocPageSize SelfTags = 19
	//DocPageMargin props: top={string} bottom={string}
	//left={string} right={string} header={string} footer={string}
	//gutter={string}
	DocPageMargin          SelfTags = 20
	WhiteSpace             SelfTags = 21
	BorderAll              SelfTags = 22
	BorderInsideHorizontal SelfTags = 23
	BorderInsideVertical   SelfTags = 24
	TableRowShading        SelfTags = 25
	TableRowMargin         SelfTags = 26
)

var WhiteListSelfTags = map[string]SelfTags{}

func initWhiteListSelfTags() {
	WhiteListSelfTags["fieldcurrentpage"] = FieldCurrentPage
	WhiteListSelfTags["fieldnumberofpages"] = FieldNumberofPages
	WhiteListSelfTags["pagebreak"] = PageBreak
	WhiteListSelfTags["br"] = LineBreak //alias of br
	WhiteListSelfTags["linebreak"] = LineBreak
	WhiteListSelfTags["inlineImg"] = InlineImage
	WhiteListSelfTags["anchorImg"] = AnchorImage
	WhiteListSelfTags["bordertop"] = BorderTop
	WhiteListSelfTags["borderright"] = BorderRight
	WhiteListSelfTags["borderleft"] = BorderLeft
	WhiteListSelfTags["borderbottom"] = BorderBottom
	WhiteListSelfTags["parashading"] = ParaShading
	WhiteListSelfTags["paraalignment"] = ParaAlignment
	WhiteListSelfTags["paratext"] = ParaText
	WhiteListSelfTags["paraframe"] = ParaFrame
	WhiteListSelfTags["paraindent"] = ParaIndent
	WhiteListSelfTags["paratextboxtightwrap"] = ParaTextBoxTightWrap
	WhiteListSelfTags["paraspacing"] = ParaSpacing
	WhiteListSelfTags["docbackground"] = DocBackground
	WhiteListSelfTags["docpagesize"] = DocPageSize
	WhiteListSelfTags["docpagemargin"] = DocPageMargin
	WhiteListSelfTags["space"] = WhiteSpace
	WhiteListSelfTags["borderall"] = BorderAll
	WhiteListSelfTags["borderinsidehorizontal"] = BorderInsideHorizontal
	WhiteListSelfTags["borderinsidevertical"] = BorderInsideVertical
	WhiteListSelfTags["rowshading"] = TableRowShading
	WhiteListSelfTags["rowmargin"] = TableRowMargin
}

type StyleTags int

const (
	// TextShading props: style={0-38},color={hex},fill={hex}
	TextShading StyleTags = -6
	// TextBorder props: style={0-193},color={hex},frame={true/false},
	//shadow={true/false},size={INT},space={INT}
	TextBorder StyleTags = -5
	//TextEffect props: style={0-7} //not working
	TextEffect StyleTags = -4
	//TextHighlight props: style={0-17}
	TextHighlight StyleTags = -3
	//Underline props: style={0-18},color={hex}
	Underline StyleTags = -2
	//Emphasis props: style={0-5}
	Emphasis StyleTags = -1
	//Font props: family={string},size={float},kern={float},
	//charSpacing={float},color={hex},scale={string},csize={string}
	Font                StyleTags = 0
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
	//Paragraph props:
	//All the properties should be mentioend only to enable them and valid values are "on,true,1"
	//keepNext,keepLines,pageBreakBefore,supressLineNumbers,
	//windowControl,wordWrap,overflowPunct,topLinePunct,autoSpaceDE,
	//autoSpaceDN,rtl,kinsoku,adjustRightInd,snapToGrid,contextualSpacing,
	//mirrorIndents,suppressOverlap,suppressAutoHyphens
	Paragraph       Tags = 16
	ParagraphBorder Tags = 17
	DocSettings     Tags = 18
	DocPageBorder   Tags = 21
	UnorderedList   Tags = 22
	OrderedList     Tags = 23
	ListItem        Tags = 24
	Heading1        Tags = 25
	Heading2        Tags = 26
	Heading3        Tags = 27
	Table           Tags = 28
	TableRow        Tags = 29
	TableData       Tags = 30
	TableBorder     Tags = 31
)

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
	WhiteListTags["docsettings"] = DocSettings
	WhiteListTags["docpageborder"] = DocPageBorder
	WhiteListTags["ul"] = UnorderedList
	WhiteListTags["ol"] = OrderedList
	WhiteListTags["li"] = ListItem
	WhiteListTags["h1"] = Heading1
	WhiteListTags["h2"] = Heading2
	WhiteListTags["h3"] = Heading3
	WhiteListTags["table"] = Table
	WhiteListTags["tr"] = TableRow
	WhiteListTags["td"] = TableData
	WhiteListTags["tableborder"] = TableBorder
}

func initConstMap() {
	initWhileListTags()
	initWhiteListSelfTags()
	initWhiteListStyleMap()
}
