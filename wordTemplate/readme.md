# Html to Docx Parser

This parser generates word document from the given _custom tags_ html file, the core purpose of this utility is to create html template and generate `.docx` file, since the template provided by the word program are very much limited in the functionality that they provide, this tool is an effort to generate complex word documents by writing html templates.

This parser uses custom tags since the html tags doesnt extend to the full length of expressiveness of the word document functionality.

The tags are classified into two categories and the same should be strictly followed:

1. Self closing `<Tag/>`
2. Container tags `<Tag></Tag>`

**Self Closing** tags would consist of set of properties that would enable us to set the word formatting/document properties.

**Container** tags can consist of properties and children elements which would be `TEXT` or other `Container` or `Self Closing` tags.

The tags consisting of properites could be of various data types, below is a list of data types and the values they accept.

> Note: Elements are case-insensitive but all attributes are case sensitive

**Data Types**

1.  _TwispMeasure:_ values can be measurements like `mm,cm,in,pt,pc,pi` can be used e.g `12cm` `12mm` `23.7cm` etc or can be `float` values like `12.34`.

2.  _SignedTwispMeasure:_ values can be measurements like `mm,cm,in,pt,pc,pi` can be used e.g `12cm` `12mm` `23.7cm` etc or can be `float` values like `12.34` but all values should be >= 0.

3.  _HexColor_: values can be 6 digit hex color code like `#000000` `#AB45DE` or `auto`.

4.  _OnOff_: values could be `true,on,1` everything else will be `false`.

# HTML

This is the starting point of the document and the document is divided into 3 parts.

1. [Document](#document)
2. [Body](#body)
3. [Formatting](#formatting)
4. [Fields](#fields)

## Document

The document tag would be the parent container consisting of tags that would set the properties of the word document

Container Tag: `<document></document>`

**Following are the immediate children of the document element:**

1.  [Document Properites](#document-properties)
2.  [Document Background](#document-background)
3.  [Document Border](#document-border)
4.  [Page Size](#page-size)
5.  [Page Margin](#page-margin)
6.  [Page Header and Page Footer](#page-header-and-page-footer)

### Document Properties

This is the container element consisting of tags which would set respectively word document meta properties.

Container Tag: `<DocProps></DocProps>`

#### Children Container Tags

These tags would set respective word document properties i.e `<Title>Demo Doc</Title>`

1.  `Company`
2.  `Title`
3.  `Author`
4.  `Description`
5.  `Category`
6.  `Version`
7.  `Application`

### Document Background

This is responsible for setting the background of the word document pages.

Tag: `<DocBackground/>`

#### Properties

- color: HexColor
- theme: ST_ThemeColor

```go
ST_ThemeColorUnset ST_ThemeColor = 0
ST_ThemeColorDark1 ST_ThemeColor = 1
ST_ThemeColorLight1 ST_ThemeColor = 2
ST_ThemeColorDark2 ST_ThemeColor = 3
ST_ThemeColorLight2 ST_ThemeColor = 4
ST_ThemeColorAccent1 ST_ThemeColor = 5
ST_ThemeColorAccent2 ST_ThemeColor = 6
ST_ThemeColorAccent3 ST_ThemeColor = 7
ST_ThemeColorAccent4 ST_ThemeColor = 8
ST_ThemeColorAccent5 ST_ThemeColor = 9
ST_ThemeColorAccent6 ST_ThemeColor = 10
ST_ThemeColorHyperlink ST_ThemeColor = 11
ST_ThemeColorFollowedHyperlink ST_ThemeColor = 12
ST_ThemeColorNone ST_ThemeColor = 13
ST_ThemeColorBackground1 ST_ThemeColor = 14
ST_ThemeColorText1 ST_ThemeColor = 15
ST_ThemeColorBackground2 ST_ThemeColor = 16
ST_ThemeColorText2 ST_ThemeColor = 17
```

### Document Border

This tag is responsible for setting the page border of the document.

Container Tag: `<DocPageBorder></DocPageBorder>`

#### Properties

- zorder: ST_PageBorderZOrder
- display: ST_PageBorderDisplay
- offset: ST_PageBorderOffset

```go
ST_PageBorderZOrderUnset ST_PageBorderZOrder = 0
ST_PageBorderZOrderFront ST_PageBorderZOrder = 1
ST_PageBorderZOrderBack ST_PageBorderZOrder = 2
```

```go
ST_PageBorderDisplayUnset ST_PageBorderDisplay = 0
ST_PageBorderDisplayAllPages ST_PageBorderDisplay = 1
ST_PageBorderDisplayFirstPage ST_PageBorderDisplay = 2
ST_PageBorderDisplayNotFirstPage ST_PageBorderDisplay = 3
```

```go
ST_PageBorderOffsetUnset ST_PageBorderOffset = 0
ST_PageBorderOffsetPage ST_PageBorderOffset = 1
ST_PageBorderOffsetText ST_PageBorderOffset = 2
```

#### Children Tags

1.  `<BorderTop/>`
2.  `<BorderBottom/>`
3.  `<BorderLeft/>`
4.  `<BorderRight/>`
5.  `<BorderAll>`

##### Properties

- size: unsignInt
- space: unsignInt
- color: HexColor
- frame : OnOff
- shadow: OnOff
- style: ST_Border

```go
ST_BorderUnset ST_Border = 0
ST_BorderNil ST_Border = 1
ST_BorderNone ST_Border = 2
ST_BorderSingle ST_Border = 3
ST_BorderThick ST_Border = 4
ST_BorderDouble ST_Border = 5
ST_BorderDotted ST_Border = 6
ST_BorderDashed ST_Border = 7
ST_BorderDotDash ST_Border = 8
ST_BorderDotDotDash ST_Border = 9
ST_BorderTriple ST_Border = 10
ST_BorderThinThickSmallGap ST_Border = 11
ST_BorderThickThinSmallGap ST_Border = 12
ST_BorderThinThickThinSmallGap ST_Border = 13
ST_BorderThinThickMediumGap ST_Border = 14
ST_BorderThickThinMediumGap ST_Border = 15
ST_BorderThinThickThinMediumGap ST_Border = 16
ST_BorderThinThickLargeGap ST_Border = 17
ST_BorderThickThinLargeGap ST_Border = 18
ST_BorderThinThickThinLargeGap ST_Border = 19
ST_BorderWave ST_Border = 20
ST_BorderDoubleWave ST_Border = 21
ST_BorderDashSmallGap ST_Border = 22
ST_BorderDashDotStroked ST_Border = 23
ST_BorderThreeDEmboss ST_Border = 24
ST_BorderThreeDEngrave ST_Border = 25
ST_BorderOutset ST_Border = 26
ST_BorderInset ST_Border = 27
ST_BorderEclipsingSquares1 ST_Border = 87
ST_BorderEclipsingSquares2 ST_Border = 88
```

### Page Margin

Tag: `<DocPageMargin/>`

#### Properties

- top: SignedTwispMeasure
- bottom: SignedTwispMeasure
- left: TwispMeasure
- right: TwispMeasure
- header: TwispMeasure
- footer: TwispMeasure
- gutter: TwispMeasure

### Page Size

This tag is responsible for setting the size of the page i.e A4,Letter etc and orientation of the document.
Tag: `<DocPageSize/>`

#### Properties

- height : TwispMeasure
- width: TwispMeasure
- orientation : ST_PageOrientation

```go
ST_PageOrientationUnset ST_PageOrientation = 0
ST_PageOrientationPortrait ST_PageOrientation = 1
ST_PageOrientationLandscape ST_PageOrientation = 2
```

### Page Header and Page Footer

These tags would add Header and Footer to the document

Container Tag:

1. `<PageHeader></PageHeader>`
2. `<PageFooter></PageFooter>`

#### Children Container Tags

Container Tags:

These tags are used to align header items

1. `<left></left>`
2. `<center></center>`
3. `<right></right>`

# Body

This container element is the starting point of the word document content section which consist of paragraphs, tables, lists, images etc.

Container Tag: `<Body></Body>`

**Following are the immediate children of the body element:**

1. Paragraph
2. Image
3. Table
4. List

## Paragraph

This container element can have text content inside of it.

## Formatting

These are set of properties that can be applied throught the document to set the formatting of the text, text can be wrapped around these tags and they can be composed as well, for e.g. `<b><i>Text</i></b>`
this is set both bold and italic styling on the text

### Following are set of tags that sets the styling of the text

1. Bold - Sets the text bold e.g `<b>bold text</b>`
2. Italic - Sets the text italic e.g `<i>italic text</i>`
3. Caps - Capitialize the text e.g `<caps>CAP</caps>`
4. SmallCaps - Shrinks the size of the capital text `<small>Text</small>`
5. StrikeThrough - Striks the text `<strike>Text</strike>`
6. DoubleStrikeThrough - Double strikes the text `<dstrike>Text</dstrike>`
7. Outline - Outlines the text `<outline>Text</outline>`
8. Shadow - Shadows the text `<shadow>Text</shadow>`
9. Emboss - Embosses the text `<emboss>Text</emboss>`
10. Imprint - Imprints the text `<imprint>Text</imprint>`
11. NoProof - To stop word from showing spellcheck error the text could be wrapped in this element `<NoProof>NAME</NoProof>`
12. Vanish - In order to hide some text from displaying in the word document the text could be wrapped in `<Vanish>Hidden Text</Vanish>` tags.
13. SubScript - To make a text sub-script wrap it in `<sub>Val</sub>`
14. SuperScript - To make a text super-script wrap it in `<sup>Val</sup>`

### Following are additional styling that requires props for styling\*\*

#### 1. Underline

`<u></u>` this element is used to underline the text

##### Properties

- style: ST_Underline
- color: Hex_String

```go
ST_UnderlineUnset           ST_Underline = 0
ST_UnderlineSingle          ST_Underline = 1
ST_UnderlineWords           ST_Underline = 2
ST_UnderlineDouble          ST_Underline = 3
ST_UnderlineThick           ST_Underline = 4
ST_UnderlineDotted          ST_Underline = 5
ST_UnderlineDottedHeavy     ST_Underline = 6
ST_UnderlineDash            ST_Underline = 7
ST_UnderlineDashedHeavy     ST_Underline = 8
ST_UnderlineDashLong        ST_Underline = 9
ST_UnderlineDashLongHeavy   ST_Underline = 10
ST_UnderlineDotDash         ST_Underline = 11
ST_UnderlineDashDotHeavy    ST_Underline = 12
ST_UnderlineDotDotDash      ST_Underline = 13
ST_UnderlineDashDotDotHeavy ST_Underline = 14
ST_UnderlineWave            ST_Underline = 15
ST_UnderlineWavyHeavy       ST_Underline = 16
ST_UnderlineWavyDouble      ST_Underline = 17
ST_UnderlineNone            ST_Underline = 18
```

#### 2. Emphasis

`<em></em>` this element is used to Emphasis the text

##### Properties

- style:

```go
ST_EmUnset    ST_Em = 0
ST_EmNone     ST_Em = 1
ST_EmDot      ST_Em = 2
ST_EmComma    ST_Em = 3
ST_EmCircle   ST_Em = 4
ST_EmUnderDot ST_Em = 5
```

#### 3. Font

`<font></font>` this element is used to set font styling for the text

##### Properties

- family: string
- size: float64
- charSpacing: float32
- color: Hex_String
- kren: float64
- scale: string_percentage

#### 4. Text Highlight

`<TextHighlight></TextHighlight>` this element highlights the text by adding background color to the text

##### Properties

- style:

```go
ST_HighlightColorUnset       ST_HighlightColor = 0
ST_HighlightColorBlack       ST_HighlightColor = 1
ST_HighlightColorBlue        ST_HighlightColor = 2
ST_HighlightColorCyan        ST_HighlightColor = 3
ST_HighlightColorGreen       ST_HighlightColor = 4
ST_HighlightColorMagenta     ST_HighlightColor = 5
ST_HighlightColorRed         ST_HighlightColor = 6
ST_HighlightColorYellow      ST_HighlightColor = 7
ST_HighlightColorWhite       ST_HighlightColor = 8
ST_HighlightColorDarkBlue    ST_HighlightColor = 9
ST_HighlightColorDarkCyan    ST_HighlightColor = 10
ST_HighlightColorDarkGreen   ST_HighlightColor = 11
ST_HighlightColorDarkMagenta ST_HighlightColor = 12
ST_HighlightColorDarkRed     ST_HighlightColor = 13
ST_HighlightColorDarkYellow  ST_HighlightColor = 14
ST_HighlightColorDarkGray    ST_HighlightColor = 15
ST_HighlightColorLightGray   ST_HighlightColor = 16
ST_HighlightColorNone        ST_HighlightColor = 17
```

#### 5. Text Border

`<TextBorder></TextBorder>` this element adds border to the text element, it all adds border on all four size, there is no option to add border to a particular size .i.e top or bottom.

##### Properties

- size: unsignInt
- space: unsignInt
- color: HexColor
- frame : OnOff
- shadow: OnOff
- style: ST_Border

```go
ST_BorderUnset ST_Border = 0
ST_BorderNil ST_Border = 1
ST_BorderNone ST_Border = 2
ST_BorderSingle ST_Border = 3
ST_BorderThick ST_Border = 4
ST_BorderDouble ST_Border = 5
ST_BorderDotted ST_Border = 6
ST_BorderDashed ST_Border = 7
ST_BorderDotDash ST_Border = 8
ST_BorderDotDotDash ST_Border = 9
ST_BorderTriple ST_Border = 10
ST_BorderThinThickSmallGap ST_Border = 11
ST_BorderThickThinSmallGap ST_Border = 12
ST_BorderThinThickThinSmallGap ST_Border = 13
ST_BorderThinThickMediumGap ST_Border = 14
ST_BorderThickThinMediumGap ST_Border = 15
ST_BorderThinThickThinMediumGap ST_Border = 16
ST_BorderThinThickLargeGap ST_Border = 17
ST_BorderThickThinLargeGap ST_Border = 18
ST_BorderThinThickThinLargeGap ST_Border = 19
ST_BorderWave ST_Border = 20
ST_BorderDoubleWave ST_Border = 21
ST_BorderDashSmallGap ST_Border = 22
ST_BorderDashDotStroked ST_Border = 23
ST_BorderThreeDEmboss ST_Border = 24
ST_BorderThreeDEngrave ST_Border = 25
ST_BorderOutset ST_Border = 26
ST_BorderInset ST_Border = 27
ST_BorderEclipsingSquares1 ST_Border = 87
ST_BorderEclipsingSquares2 ST_Border = 88
```

#### 6. Text Shading

`<TextShading></TextShading>` this element adds shading to the text.

##### Properties

- style: ST_Shd
- color: HexColor
- fill: HexColor

```go
ST_ShdUnset                 ST_Shd = 0
ST_ShdNil                   ST_Shd = 1
ST_ShdClear                 ST_Shd = 2
ST_ShdSolid                 ST_Shd = 3
ST_ShdHorzStripe            ST_Shd = 4
ST_ShdVertStripe            ST_Shd = 5
ST_ShdReverseDiagStripe     ST_Shd = 6
ST_ShdDiagStripe            ST_Shd = 7
ST_ShdHorzCross             ST_Shd = 8
ST_ShdDiagCross             ST_Shd = 9
ST_ShdThinHorzStripe        ST_Shd = 10
ST_ShdThinVertStripe        ST_Shd = 11
ST_ShdThinReverseDiagStripe ST_Shd = 12
ST_ShdThinDiagStripe        ST_Shd = 13
ST_ShdThinHorzCross         ST_Shd = 14
ST_ShdThinDiagCross         ST_Shd = 15
ST_ShdPct5                  ST_Shd = 16
ST_ShdPct10                 ST_Shd = 17
ST_ShdPct12                 ST_Shd = 18
ST_ShdPct15                 ST_Shd = 19
ST_ShdPct20                 ST_Shd = 20
ST_ShdPct25                 ST_Shd = 21
ST_ShdPct30                 ST_Shd = 22
ST_ShdPct35                 ST_Shd = 23
ST_ShdPct37                 ST_Shd = 24
ST_ShdPct40                 ST_Shd = 25
ST_ShdPct45                 ST_Shd = 26
ST_ShdPct50                 ST_Shd = 27
ST_ShdPct55                 ST_Shd = 28
ST_ShdPct60                 ST_Shd = 29
ST_ShdPct62                 ST_Shd = 30
ST_ShdPct65                 ST_Shd = 31
ST_ShdPct70                 ST_Shd = 32
ST_ShdPct75                 ST_Shd = 33
ST_ShdPct80                 ST_Shd = 34
ST_ShdPct85                 ST_Shd = 35
ST_ShdPct87                 ST_Shd = 36
ST_ShdPct90                 ST_Shd = 37
ST_ShdPct95                 ST_Shd = 38
```
