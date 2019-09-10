# Html to Docx Parser

This parser generates word document from the given _custom tags_ Html file, the core purpose of this utility is to create Html template and generate `.docx` file, since the template provided by the word program are very much limited in the functionality that they provide, this tool is an effort to generate complex word documents by writing Html templates.

This parser uses custom tags since the Html tags doesn't extend to the full length of expressiveness of the word document functionality.

The tags are classified into two categories and the same should be strictly followed:

- Self-closing `<Tag/>`
- Container tags `<Tag></Tag>`

**Self Closing** tags would consist of a set of properties that would enable us to set the word formatting/document properties.

**Container** tags can consist of properties and children elements which would be `TEXT` or other `Container` or `Self Closing` tags.

The tags consisting of properties could be of various data types, below is a list of data types and the values they accept.

> Note: Elements are case-insensitive but all attributes are case sensitive

**Data Types**

1.  #### TwispMeasure

    values can be measurements like `mm, cm, in, pt, pc, pi` can be used e.g `12cm` `12mm` `23.7cm` etc or can be `float` values like `12.34`.

2.  #### SignedTwispMeasure

    values can be measurements like `mm, cm, in, pt, pc, pi` can be used e.g `12cm` `12mm` `23.7cm` etc or can be `float` values like `12.34` but all values should be >= 0.

3.  #### HexColor

    values can be 6 digit hex color code like `#000000` `#AB45DE` or `auto`.

4.  #### OnOff
    values could be `true, on, 1` everything else will be `false`.

# HTML

This is the starting point of the document and the document is divided into 3 parts.

1. [Document](#document)
2. [Formatting](#formatting)
3. [Fields](#fields)
4. [Body](#body)

## Document

The document tag would be the parent container consisting of tags that would set the properties of the word document

Container Tag: `<document></document>`

**Following are the immediate children of the document element:**

1.  [Document Properties](#document-properties)
2.  [Document Background](#document-background)
3.  [Document Border](#document-border)
4.  [Page Size](#page-size)
5.  [Page Margin](#page-margin)
6.  [Page Header and Page Footer](#page-header-and-page-footer)

---

1. ### Document Properties

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

2)  ### Document Background

This is responsible for setting the background of the word document pages.

Tag: `<DocBackground/>`

#### Properties

- color: [HexColor](#hex-color)
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

3. ### Document Border

This tag is responsible for setting the page border of the document.

Container Tag: `<DocPageBorder></DocPageBorder>`

#### Properties

- zorder: ST_PageBorderZOrder
- display: ST_PageBorderDisplay
- offset: ST_PageBorderOffset

```go
ST_PageBorderZOrderUnset ST_PageBorderZOrder = 0
ST_PageBorderZOrderFront ST_PageBorderZOrder = 1
ST_PageBorderZOrderBack  ST_PageBorderZOrder = 2
```

```go
ST_PageBorderDisplayUnset        ST_PageBorderDisplay = 0
ST_PageBorderDisplayAllPages     ST_PageBorderDisplay = 1
ST_PageBorderDisplayFirstPage    ST_PageBorderDisplay = 2
ST_PageBorderDisplayNotFirstPage ST_PageBorderDisplay = 3
```

```go
ST_PageBorderOffsetUnset ST_PageBorderOffset = 0
ST_PageBorderOffsetPage  ST_PageBorderOffset = 1
ST_PageBorderOffsetText  ST_PageBorderOffset = 2
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
ST_BorderUnset                  ST_Border = 0
ST_BorderNil                    ST_Border = 1
ST_BorderNone                   ST_Border = 2
ST_BorderSingle                 ST_Border = 3
ST_BorderThick                  ST_Border = 4
ST_BorderDouble                 ST_Border = 5
ST_BorderDotted                 ST_Border = 6
ST_BorderDashed                 ST_Border = 7
ST_BorderDotDash                ST_Border = 8
ST_BorderDotDotDash             ST_Border = 9
ST_BorderTriple                 ST_Border = 10
ST_BorderThinThickSmallGap      ST_Border = 11
ST_BorderThickThinSmallGap      ST_Border = 12
ST_BorderThinThickThinSmallGap  ST_Border = 13
ST_BorderThinThickMediumGap     ST_Border = 14
ST_BorderThickThinMediumGap     ST_Border = 15
ST_BorderThinThickThinMediumGap ST_Border = 16
ST_BorderThinThickLargeGap      ST_Border = 17
ST_BorderThickThinLargeGap      ST_Border = 18
ST_BorderThinThickThinLargeGap  ST_Border = 19
ST_BorderWave                   ST_Border = 20
ST_BorderDoubleWave             ST_Border = 21
ST_BorderDashSmallGap           ST_Border = 22
ST_BorderDashDotStroked         ST_Border = 23
ST_BorderThreeDEmboss           ST_Border = 24
ST_BorderThreeDEngrave          ST_Border = 25
ST_BorderOutset                 ST_Border = 26
ST_BorderInset                  ST_Border = 27
ST_BorderEclipsingSquares1      ST_Border = 87
ST_BorderEclipsingSquares2      ST_Border = 88
```

4. ### Page Margin

Tag: `<DocPageMargin/>`

#### Properties

- top: SignedTwispMeasure
- bottom: SignedTwispMeasure
- left: TwispMeasure
- right: TwispMeasure
- header: TwispMeasure
- footer: TwispMeasure
- gutter: TwispMeasure

5. ### Page Size

This tag is responsible for setting the size of the page i.e A4, Letter etc and orientation of the document.
Tag: `<DocPageSize/>`

#### Properties

- height : TwispMeasure
- width: TwispMeasure
- orientation : ST_PageOrientation

```go
ST_PageOrientationUnset     ST_PageOrientation = 0
ST_PageOrientationPortrait  ST_PageOrientation = 1
ST_PageOrientationLandscape ST_PageOrientation = 2
```

6. ### Page Header and Page Footer

These tags would add Header and Footer to the document

Container Tags:

1. `<PageHeader></PageHeader>`
2. `<PageFooter></PageFooter>`

#### Children Container Tags

Container Tags:

These tags are used to align header items, text elements

1. `<left></left>`
2. `<center></center>`
3. `<right></right>`

# Formatting

These are set of properties that can be applied through the document to set the formatting of the text, text can be wrapped around these tags and they can be composed as well, for e.g. `<b><i>Text</i></b>`
this is set both bold and italic styling on the text

**Following are set of tags that set the styling of the text**

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

**Following are additional styling that requires properties for styling**

1. [Underline](#underline)
2. [Emphasis](#emphasis)
3. [Font](#font)
4. [Text Highlight](#text-highlight)
5. [Text Border](#text-border)
6. [Text Shading](#text-shading)

1) #### Underline

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

2. #### Emphasis

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

3. #### Font

`<font></font>` this element is used to set font styling for the text

##### Properties

- family: string
- size: float64
- charSpacing: float32
- color: Hex_String
- kren: float64
- scale: string_percentage

#### Text Highlight

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

5. #### Text Border

`<TextBorder></TextBorder>` this element adds a border to the text element, it all adds border on all four sides, there is no option to add border to a particular size .i.e top or bottom.

##### Properties

- size: unsignInt
- space: unsignInt
- color: HexColor
- frame : OnOff
- shadow: OnOff
- style: ST_Border

```go
ST_BorderUnset                  ST_Border = 0
ST_BorderNil                    ST_Border = 1
ST_BorderNone                   ST_Border = 2
ST_BorderSingle                 ST_Border = 3
ST_BorderThick                  ST_Border = 4
ST_BorderDouble                 ST_Border = 5
ST_BorderDotted                 ST_Border = 6
ST_BorderDashed                 ST_Border = 7
ST_BorderDotDash                ST_Border = 8
ST_BorderDotDotDash             ST_Border = 9
ST_BorderTriple                 ST_Border = 10
ST_BorderThinThickSmallGap      ST_Border = 11
ST_BorderThickThinSmallGap      ST_Border = 12
ST_BorderThinThickThinSmallGap  ST_Border = 13
ST_BorderThinThickMediumGap     ST_Border = 14
ST_BorderThickThinMediumGap     ST_Border = 15
ST_BorderThinThickThinMediumGap ST_Border = 16
ST_BorderThinThickLargeGap      ST_Border = 17
ST_BorderThickThinLargeGap      ST_Border = 18
ST_BorderThinThickThinLargeGap  ST_Border = 19
ST_BorderWave                   ST_Border = 20
ST_BorderDoubleWave             ST_Border = 21
ST_BorderDashSmallGap           ST_Border = 22
ST_BorderDashDotStroked         ST_Border = 23
ST_BorderThreeDEmboss           ST_Border = 24
ST_BorderThreeDEngrave          ST_Border = 25
ST_BorderOutset                 ST_Border = 26
ST_BorderInset                  ST_Border = 27
ST_BorderEclipsingSquares1      ST_Border = 87
ST_BorderEclipsingSquares2      ST_Border = 88
```

6. #### Text Shading

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

# Fields

These are special tags used to insert special characters, or field values in the word document

### Following are the tags to insert

1. Current Page: `<FieldCurrentPage/>`  
   This tag is used to insert a current page number into the word document
2. No Of Pages: `<FieldNumberofPages/>`  
   The tag is used to insert total count of pages into the word document.
3. PageBreak: `<PageBreak/>`
   This tag is used to insert page break into the document and add a new page
4. LineBreak: `<br/>`  
   This tag is used to break the line within a document paragraph. and has property `count` which can be used to set multiple page breaks i.e `<br count="2"/>` will put 2 line breaks
5. WhiteSpace: `<space/>`
   This tag is used to insert space into the document, this is useful when space is needed before the start and after the end of the container tag's text.

# Body

This container element is the starting point of the word document content section which consists of paragraphs, tables, lists, images, etc.

Container Tag: `<Body></Body>`

**Following are the immediate children of the body element:**

1. [Paragraph](#paragraph)
2. [Image](#image)
3. [Table](#table)
4. [List](#list)

## Paragraph

The paragraph tag is used to add special properties to the text enclosing text paragraph.

Container Tag: `<Paragraph></Paragraph>`

### Properties

- keepNext: OnOff
  Keep Paragraph With Next Paragraph
- keepLines: onOff
  Keep All Lines On One Page
- pageBreakBefore: onOff
  Start Paragraph on Next Page
- wordWrap: onOff
  Allow Line Breaking At Character Level
- overflowPunct: onOff
  Allow Punctuation to Extend Past Text Extents
- rtl: onOff
  Right to Left Paragraph Layout
- suppressOverlap: onOff
  Prevent Text Frames From Overlapping
- widowControl: onOff
  Allow First/Last Line to Display on a Separate Page

### Child Tags

1. [Paragraph Border](#paragraph-border)
2. [Paragraph Shading](#paragraph-shading)
3. [Paragraph Alignment](#paragraph-alignment)
4. [Paragraph Spacing](#paragraph-spacing)
5. [Paragraph Indentation](#paragraph-indentation)
6. [Paragraph Frame](#paragraph-frame)

1) #### Paragraph Border

This sets the border on all or selected sides of the paragraph.

Container Tag: `<ParaBorder></ParaBorder>`

##### Children Tags

1.  `<BorderTop/>`
2.  `<BorderBottom/>`
3.  `<BorderLeft/>`
4.  `<BorderRight/>`
5.  `<BorderAll/>`

###### Properties

- size: unsignInt
- space: unsignInt
- color: HexColor
- frame : OnOff
- shadow: OnOff
- style: ST_Border

```go
ST_BorderUnset                  ST_Border = 0
ST_BorderNil                    ST_Border = 1
ST_BorderNone                   ST_Border = 2
ST_BorderSingle                 ST_Border = 3
ST_BorderThick                  ST_Border = 4
ST_BorderDouble                 ST_Border = 5
ST_BorderDotted                 ST_Border = 6
ST_BorderDashed                 ST_Border = 7
ST_BorderDotDash                ST_Border = 8
ST_BorderDotDotDash             ST_Border = 9
ST_BorderTriple                 ST_Border = 10
ST_BorderThinThickSmallGap      ST_Border = 11
ST_BorderThickThinSmallGap      ST_Border = 12
ST_BorderThinThickThinSmallGap  ST_Border = 13
ST_BorderThinThickMediumGap     ST_Border = 14
ST_BorderThickThinMediumGap     ST_Border = 15
ST_BorderThinThickThinMediumGap ST_Border = 16
ST_BorderThinThickLargeGap      ST_Border = 17
ST_BorderThickThinLargeGap      ST_Border = 18
ST_BorderThinThickThinLargeGap  ST_Border = 19
ST_BorderWave                   ST_Border = 20
ST_BorderDoubleWave             ST_Border = 21
ST_BorderDashSmallGap           ST_Border = 22
ST_BorderDashDotStroked         ST_Border = 23
ST_BorderThreeDEmboss           ST_Border = 24
ST_BorderThreeDEngrave          ST_Border = 25
ST_BorderOutset                 ST_Border = 26
ST_BorderInset                  ST_Border = 27
ST_BorderEclipsingSquares1      ST_Border = 87
ST_BorderEclipsingSquares2      ST_Border = 88
```

2. #### Paragraph Shading

Sets the shadings to the paragraph and all text enclosed in it.

Tag: `<ParaShading/>`

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

3. #### Paragraph Alignment

Aligns the text content within the paragraph

Tag: `<ParaAlign/>`

##### Properties

- style: ST_Jc

```go
ST_JcUnset          ST_Jc = 0
ST_JcStart          ST_Jc = 1
ST_JcCenter         ST_Jc = 2
ST_JcEnd            ST_Jc = 3
ST_JcBoth           ST_Jc = 4
ST_JcMediumKashida  ST_Jc = 5
ST_JcDistribute     ST_Jc = 6
ST_JcNumTab         ST_Jc = 7
ST_JcHighKashida    ST_Jc = 8
ST_JcLowKashida     ST_Jc = 9
ST_JcThaiDistribute ST_Jc = 10
ST_JcLeft           ST_Jc = 11
ST_JcRight          ST_Jc = 12
```

4. #### Paragraph Spacing

Sets the paragraph spacing

Tag: `<ParaSpacing/>`

##### Properties

- after: int
- autoAfter: bool
- before: int
- autoBefore: bool
- linespacing: int,ST_LineSpacingRule

```go
ST_LineSpacingRuleUnset   ST_LineSpacingRule = 0
ST_LineSpacingRuleExact   ST_LineSpacingRule = 2
ST_LineSpacingRuleAuto    ST_LineSpacingRule = 1
ST_LineSpacingRuleAtLeast ST_LineSpacingRule = 3
```

5. #### Paragraph Indentation

Set the Paragraph Indentation

Tag: `<ParaIndent/>`

##### Properties

- start: int
- end: int
- hang: int
- first: int

6. #### Paragraph Frame

Sets the Paragraph Frame

Tag: `<ParaFrame>`

##### Properties

- dropCap: ST_DropCap
- lines: int
- wrap: ST_Wrap
- hAnchor: ST_HAnchor
- vAnchor: ST_VAnchor
- xAlign: ST_XAlign
- yAlign: ST_YAlign
- hRule: ST_HeightRule
- height: TwipsMeasure
- width: TwipsMeasure
- hpad: TwipsMeasure
- vpad: TwipsMeasure

```go
ST_DropCapUnset  ST_DropCap = 0
ST_DropCapNone   ST_DropCap = 1
ST_DropCapDrop   ST_DropCap = 2
ST_DropCapMargin ST_DropCap = 3
```

```go
ST_WrapUnset     ST_Wrap = 0
ST_WrapAuto      ST_Wrap = 1
ST_WrapNotBeside ST_Wrap = 2
ST_WrapAround    ST_Wrap = 3
ST_WrapTight     ST_Wrap = 4
ST_WrapThrough   ST_Wrap = 5
ST_WrapNone      ST_Wrap = 6
```

```go
ST_HAnchorUnset  ST_HAnchor = 0
ST_HAnchorText   ST_HAnchor = 1
ST_HAnchorMargin ST_HAnchor = 2
ST_HAnchorPage   ST_HAnchor = 3
```

```go
ST_VAnchorUnset  ST_VAnchor = 0
ST_VAnchorText   ST_VAnchor = 1
ST_VAnchorMargin ST_VAnchor = 2
ST_VAnchorPage   ST_VAnchor = 3
```

```go
ST_XAlignUnset   ST_XAlign = 0
ST_XAlignLeft    ST_XAlign = 1
ST_XAlignCenter  ST_XAlign = 2
ST_XAlignRight   ST_XAlign = 3
ST_XAlignInside  ST_XAlign = 4
ST_XAlignOutside ST_XAlign = 5
```

```go
ST_YAlignUnset   ST_YAlign = 0
ST_YAlignInline  ST_YAlign = 1
ST_YAlignTop     ST_YAlign = 2
ST_YAlignCenter  ST_YAlign = 3
ST_YAlignBottom  ST_YAlign = 4
ST_YAlignInside  ST_YAlign = 5
ST_YAlignOutside ST_YAlign = 6
```

```go
ST_HeightRuleUnset   ST_HeightRule = 0
ST_HeightRuleAuto    ST_HeightRule = 1
ST_HeightRuleExact   ST_HeightRule = 2
ST_HeightRuleAtLeast ST_HeightRule = 3
```

## Image

Images can be included into the word document

**Following are the ways images can be incorporated**

1. [Inline Image](#inline-image)
2. [Anchored Image](#anchored-image)

1) ### Inline Image

   This image can be inline along with the lines of the text
   Tag: `<InlineImg/>`

#### Properties

- src: URL
- size: width, height // (int (Inches))

2. ### Anchored Image

   This image will be anchored at a specific location in the document
   Tag: `<AnchorImg/>`

#### Properties

- src: URL
- noWrap: OnOff
- name: string
- yOffset: int //(Inches)
- xOffset: int //(Inches)
- size: width, height //(int (Inches))
- origin: WdST_RelFromH, WdST_RelFromV
- hAlign: WdST_AlignH
- vAlign: WdST_AlignV
- wrap: WdST_WrapText

```go
WdST_RelFromHUnset         WdST_RelFromH = 0
WdST_RelFromHMargin        WdST_RelFromH = 1
WdST_RelFromHPage          WdST_RelFromH = 2
WdST_RelFromHColumn        WdST_RelFromH = 3
WdST_RelFromHCharacter     WdST_RelFromH = 4
WdST_RelFromHLeftMargin    WdST_RelFromH = 5
WdST_RelFromHRightMargin   WdST_RelFromH = 6
WdST_RelFromHInsideMargin  WdST_RelFromH = 7
WdST_RelFromHOutsideMargin WdST_RelFromH = 8
```

```go
WdST_RelFromVUnset         WdST_RelFromV = 0
WdST_RelFromVMargin        WdST_RelFromV = 1
WdST_RelFromVPage          WdST_RelFromV = 2
WdST_RelFromVParagraph     WdST_RelFromV = 3
WdST_RelFromVLine          WdST_RelFromV = 4
WdST_RelFromVTopMargin     WdST_RelFromV = 5
WdST_RelFromVBottomMargin  WdST_RelFromV = 6
WdST_RelFromVInsideMargin  WdST_RelFromV = 7
WdST_RelFromVOutsideMargin WdST_RelFromV = 8
```

```go
WdST_AlignHUnset   WdST_AlignH = 0
WdST_AlignHLeft    WdST_AlignH = 1
WdST_AlignHRight   WdST_AlignH = 2
WdST_AlignHCenter  WdST_AlignH = 3
WdST_AlignHInside  WdST_AlignH = 4
WdST_AlignHOutside WdST_AlignH = 5
```

```go
WdST_AlignVUnset   WdST_AlignV = 0
WdST_AlignVTop     WdST_AlignV = 1
WdST_AlignVBottom  WdST_AlignV = 2
WdST_AlignVCenter  WdST_AlignV = 3
WdST_AlignVInside  WdST_AlignV = 4
WdST_AlignVOutside WdST_AlignV = 5
```

```go
WdST_WrapTextUnset     WdST_WrapText = 0
WdST_WrapTextBothSides WdST_WrapText = 1
WdST_WrapTextLeft      WdST_WrapText = 2
WdST_WrapTextRight     WdST_WrapText = 3
WdST_WrapTextLargest   WdST_WrapText = 4
```

## Table

This tag will mark the start of the table in the word document

Container Tag: `<table>`

### Properties

- align: ST_JcTable
- layout: ST_TblLayoutType
- width: Percentage Int

```go
ST_JcTableUnset  ST_JcTable = 0
ST_JcTableCenter ST_JcTable = 1
ST_JcTableEnd    ST_JcTable = 2
ST_JcTableLeft   ST_JcTable = 3
ST_JcTableRight  ST_JcTable = 4
ST_JcTableStart  ST_JcTable = 5
```

```go
ST_TblLayoutTypeUnset   ST_TblLayoutType = 0
ST_TblLayoutTypeFixed   ST_TblLayoutType = 1
ST_TblLayoutTypeAutofit  ST_TblLayoutType = 2
```

**Following are the immediate children of the Table Tag**

1. [Table Border](#table-border)
2. [Table Row](#table-row)

1) ### Table Border
   Container Tag: `<TableBorder>`

##### Properties

- size: unsignInt
- space: unsignInt
- color: HexColor
- frame : OnOff
- shadow: OnOff
- style: ST_Border

##### Children Tags

1.  `<BorderTop/>`
2.  `<BorderBottom/>`
3.  `<BorderLeft/>`
4.  `<BorderRight/>`
5.  `<BorderAll/>`
6.  `<BorderInsideHorizontal/>`
7.  `<BorderInsideVertical/>`

2) ### Table Row

Container Tag: `<tr></tr>`

#### Properties

- height: int,rule: ST_HeightRule

```go
ST_HeightRuleUnset   ST_HeightRule = 0
ST_HeightRuleAuto    ST_HeightRule = 1
ST_HeightRuleExact   ST_HeightRule = 2
ST_HeightRuleAtLeast ST_HeightRule = 3
```

#### Children Tags

1. [Table Row Shading](#table-row-shading)
2. [Table Row Margin](#table-row-margin)
3. [Table Data](#table-data)

1) ##### Table Row Shading

Tag: `<RowShading/>`

###### Properties

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

2. ##### Table Row Margin

Tag: `<RowMargin/>`

###### Properties

- top: int
- bottom: int
- left: int
- right: int

3. ##### Table Data

Container Tag: `<td></td>`

###### Properties

- colspan: int
- rowspan: int
- align: ST_Jc
- valign: ST_VerticalJc
- width: int Percentage

```go
ST_VerticalJcUnset  ST_VerticalJc = 0
ST_VerticalJcTop    ST_VerticalJc = 1
ST_VerticalJcCenter ST_VerticalJc = 2
ST_VerticalJcBoth   ST_VerticalJc = 3
ST_VerticalJcBottom ST_VerticalJc = 4
```

```go
ST_JcUnset          ST_Jc = 0
ST_JcStart          ST_Jc = 1
ST_JcCenter         ST_Jc = 2
ST_JcEnd            ST_Jc = 3
ST_JcBoth           ST_Jc = 4
ST_JcMediumKashida  ST_Jc = 5
ST_JcDistribute     ST_Jc = 6
ST_JcNumTab         ST_Jc = 7
ST_JcHighKashida    ST_Jc = 8
ST_JcLowKashida     ST_Jc = 9
ST_JcThaiDistribute ST_Jc = 10
ST_JcLeft           ST_Jc = 11
ST_JcRight          ST_Jc = 12
```

## List

**Following are the types of list**

1. Ordered List
2. Unordered List

Container Tags:

1. Ordered List - `<ol></ol>`
2. Unordered List - `<ul></ul>`

#### Properties

- indent: int
  Indent determines the indentation of the list item
- hangIndent: int
  hangIndent determines the space between list item and list number
- align: ST_Jc
- style:
  for OrderedList: ST_NumberFormat, for UnOrderedList: Text `` or any other text, it uses font family Symbol.

```go
ST_JcUnset          ST_Jc = 0
ST_JcStart          ST_Jc = 1
ST_JcCenter         ST_Jc = 2
ST_JcEnd            ST_Jc = 3
ST_JcBoth           ST_Jc = 4
ST_JcMediumKashida  ST_Jc = 5
ST_JcDistribute     ST_Jc = 6
ST_JcNumTab         ST_Jc = 7
ST_JcHighKashida    ST_Jc = 8
ST_JcLowKashida     ST_Jc = 9
ST_JcThaiDistribute ST_Jc = 10
ST_JcLeft           ST_Jc = 11
ST_JcRight          ST_Jc = 12
```

```go
	ST_NumberFormatUnset                        ST_NumberFormat = 0
	ST_NumberFormatDecimal                      ST_NumberFormat = 1
	ST_NumberFormatUpperRoman                   ST_NumberFormat = 2
	ST_NumberFormatLowerRoman                   ST_NumberFormat = 3
	ST_NumberFormatUpperLetter                  ST_NumberFormat = 4
	ST_NumberFormatLowerLetter                  ST_NumberFormat = 5
	ST_NumberFormatOrdinal                      ST_NumberFormat = 6
	ST_NumberFormatCardinalText                 ST_NumberFormat = 7
	ST_NumberFormatOrdinalText                  ST_NumberFormat = 8
	ST_NumberFormatHex                          ST_NumberFormat = 9
	ST_NumberFormatChicago                      ST_NumberFormat = 10
	ST_NumberFormatIdeographDigital             ST_NumberFormat = 11
	ST_NumberFormatJapaneseCounting             ST_NumberFormat = 12
	ST_NumberFormatAiueo                        ST_NumberFormat = 13
	ST_NumberFormatIroha                        ST_NumberFormat = 14
	ST_NumberFormatDecimalFullWidth             ST_NumberFormat = 15
	ST_NumberFormatDecimalHalfWidth             ST_NumberFormat = 16
	ST_NumberFormatJapaneseLegal                ST_NumberFormat = 17
	ST_NumberFormatJapaneseDigitalTenThousand   ST_NumberFormat = 18
	ST_NumberFormatDecimalEnclosedCircle        ST_NumberFormat = 19
	ST_NumberFormatDecimalFullWidth2            ST_NumberFormat = 20
	ST_NumberFormatAiueoFullWidth               ST_NumberFormat = 21
	ST_NumberFormatIrohaFullWidth               ST_NumberFormat = 22
	ST_NumberFormatDecimalZero                  ST_NumberFormat = 23
	ST_NumberFormatBullet                       ST_NumberFormat = 24
	ST_NumberFormatGanada                       ST_NumberFormat = 25
	ST_NumberFormatChosung                      ST_NumberFormat = 26
	ST_NumberFormatDecimalEnclosedFullstop      ST_NumberFormat = 27
	ST_NumberFormatDecimalEnclosedParen         ST_NumberFormat = 28
	ST_NumberFormatDecimalEnclosedCircleChinese ST_NumberFormat = 29
	ST_NumberFormatIdeographEnclosedCircle      ST_NumberFormat = 30
	ST_NumberFormatIdeographTraditional         ST_NumberFormat = 31
	ST_NumberFormatIdeographZodiac              ST_NumberFormat = 32
	ST_NumberFormatIdeographZodiacTraditional   ST_NumberFormat = 33
	ST_NumberFormatTaiwaneseCounting            ST_NumberFormat = 34
	ST_NumberFormatIdeographLegalTraditional    ST_NumberFormat = 35
	ST_NumberFormatTaiwaneseCountingThousand    ST_NumberFormat = 36
	ST_NumberFormatTaiwaneseDigital             ST_NumberFormat = 37
	ST_NumberFormatChineseCounting              ST_NumberFormat = 38
	ST_NumberFormatChineseLegalSimplified       ST_NumberFormat = 39
	ST_NumberFormatChineseCountingThousand      ST_NumberFormat = 40
	ST_NumberFormatKoreanDigital                ST_NumberFormat = 41
	ST_NumberFormatKoreanCounting               ST_NumberFormat = 42
	ST_NumberFormatKoreanLegal                  ST_NumberFormat = 43
	ST_NumberFormatKoreanDigital2               ST_NumberFormat = 44
	ST_NumberFormatVietnameseCounting           ST_NumberFormat = 45
	ST_NumberFormatRussianLower                 ST_NumberFormat = 46
	ST_NumberFormatRussianUpper                 ST_NumberFormat = 47
	ST_NumberFormatNone                         ST_NumberFormat = 48
	ST_NumberFormatNumberInDash                 ST_NumberFormat = 49
	ST_NumberFormatHebrew1                      ST_NumberFormat = 50
	ST_NumberFormatHebrew2                      ST_NumberFormat = 51
	ST_NumberFormatArabicAlpha                  ST_NumberFormat = 52
	ST_NumberFormatArabicAbjad                  ST_NumberFormat = 53
	ST_NumberFormatHindiVowels                  ST_NumberFormat = 54
	ST_NumberFormatHindiConsonants              ST_NumberFormat = 55
	ST_NumberFormatHindiNumbers                 ST_NumberFormat = 56
	ST_NumberFormatHindiCounting                ST_NumberFormat = 57
	ST_NumberFormatThaiLetters                  ST_NumberFormat = 58
	ST_NumberFormatThaiNumbers                  ST_NumberFormat = 59
	ST_NumberFormatThaiCounting                 ST_NumberFormat = 60
	ST_NumberFormatBahtText                     ST_NumberFormat = 61
	ST_NumberFormatDollarText                   ST_NumberFormat = 62
	ST_NumberFormatCustom                       ST_NumberFormat = 63

```
