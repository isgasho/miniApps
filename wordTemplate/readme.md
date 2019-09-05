#A Template based html to docx generator

The document tags are classified as self closing `<Tag/>` and container tags `<Tag></Tag>` and the same should be strictly followed.

###Data Types

1. **TwispMeasure:**
   values in `mm,cm,in,pt,pc,pi` like `12cm` `12mm` `23.7cm` etc or can be **float** values like `12.34`

2. **SignedTwispMeasure:**
   values in `mm,cm,in,pt,pc,pi` like `12cm` `12mm` `23.7cm` etc or can be **float** values like `12.34`
   but has to `>=0`

3. **HexColor**
   values can be 6 digit hex color code like `#000000` `#AB45DE` or `auto`

4. **OnOff**
   values could be `true,on,1` everything else will be `false`

## Document

The document would be the parent container for all the below properties
Container Tag: `<document></document>`

1. [Document Properites](#document-properties)
2. [Page Size](#page-size)
3. Page Margin
4. Document Background
5. Document Border

### Document Properties

Container Tag: `<DocProps></DocProps>`

#### Container Tags

These tags would contain text that would be inserted in the following properties i.e `<Title>Demo Doc</Title>`

1. `Company`
2. `Title`
3. `Author`
4. `Description`
5. `Category`
6. `Version`
7. `Application`

### Page Size

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

### Document Background

Tag: `<DocBackground/>`

#### Properties

- color: HexColor
- theme: ST_ThemeColor

```go
ST_ThemeColorUnset             ST_ThemeColor = 0
ST_ThemeColorDark1             ST_ThemeColor = 1
ST_ThemeColorLight1            ST_ThemeColor = 2
ST_ThemeColorDark2             ST_ThemeColor = 3
ST_ThemeColorLight2            ST_ThemeColor = 4
ST_ThemeColorAccent1           ST_ThemeColor = 5
ST_ThemeColorAccent2           ST_ThemeColor = 6
ST_ThemeColorAccent3           ST_ThemeColor = 7
ST_ThemeColorAccent4           ST_ThemeColor = 8
ST_ThemeColorAccent5           ST_ThemeColor = 9
ST_ThemeColorAccent6           ST_ThemeColor = 10
ST_ThemeColorHyperlink         ST_ThemeColor = 11
ST_ThemeColorFollowedHyperlink ST_ThemeColor = 12
ST_ThemeColorNone              ST_ThemeColor = 13
ST_ThemeColorBackground1       ST_ThemeColor = 14
ST_ThemeColorText1             ST_ThemeColor = 15
ST_ThemeColorBackground2       ST_ThemeColor = 16
ST_ThemeColorText2             ST_ThemeColor = 17
```

### Document Border

The container tag is `<DocPageBorder></DocPageBorder>`

#### Properties

- zorder: ST_PageBorderZOrder

```go
ST_PageBorderZOrderUnset ST_PageBorderZOrder = 0
ST_PageBorderZOrderFront ST_PageBorderZOrder = 1
ST_PageBorderZOrderBack  ST_PageBorderZOrder = 2
```

- display: ST_PageBorderDisplay

```go
ST_PageBorderDisplayUnset        ST_PageBorderDisplay = 0
ST_PageBorderDisplayAllPages     ST_PageBorderDisplay = 1
ST_PageBorderDisplayFirstPage    ST_PageBorderDisplay = 2
ST_PageBorderDisplayNotFirstPage ST_PageBorderDisplay = 3
```

- offset: ST_PageBorderOffset

```go
ST_PageBorderOffsetUnset ST_PageBorderOffset = 0
	ST_PageBorderOffsetPage  ST_PageBorderOffset = 1
	ST_PageBorderOffsetText  ST_PageBorderOffset = 2
```

#### Document Child Tags

1. `<BorderTop/>`
2. `<BorderBottom/>`
3. `<BorderLeft/>`
4. `<BorderRight/>`

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
	ST_BorderApples                 ST_Border = 28
	ST_BorderArchedScallops         ST_Border = 29
	ST_BorderBabyPacifier           ST_Border = 30
	ST_BorderBabyRattle             ST_Border = 31
	ST_BorderBalloons3Colors        ST_Border = 32
	ST_BorderBalloonsHotAir         ST_Border = 33
	ST_BorderBasicBlackDashes       ST_Border = 34
	ST_BorderBasicBlackDots         ST_Border = 35
	ST_BorderBasicBlackSquares      ST_Border = 36
	ST_BorderBasicThinLines         ST_Border = 37
	ST_BorderBasicWhiteDashes       ST_Border = 38
	ST_BorderBasicWhiteDots         ST_Border = 39
	ST_BorderBasicWhiteSquares      ST_Border = 40
	ST_BorderBasicWideInline        ST_Border = 41
	ST_BorderBasicWideMidline       ST_Border = 42
	ST_BorderBasicWideOutline       ST_Border = 43
	ST_BorderBats                   ST_Border = 44
	ST_BorderBirds                  ST_Border = 45
	ST_BorderBirdsFlight            ST_Border = 46
	ST_BorderCabins                 ST_Border = 47
	ST_BorderCakeSlice              ST_Border = 48
	ST_BorderCandyCorn              ST_Border = 49
	ST_BorderCelticKnotwork         ST_Border = 50
	ST_BorderCertificateBanner      ST_Border = 51
	ST_BorderChainLink              ST_Border = 52
	ST_BorderChampagneBottle        ST_Border = 53
	ST_BorderCheckedBarBlack        ST_Border = 54
	ST_BorderCheckedBarColor        ST_Border = 55
	ST_BorderCheckered              ST_Border = 56
	ST_BorderChristmasTree          ST_Border = 57
	ST_BorderCirclesLines           ST_Border = 58
	ST_BorderCirclesRectangles      ST_Border = 59
	ST_BorderClassicalWave          ST_Border = 60
	ST_BorderClocks                 ST_Border = 61
	ST_BorderCompass                ST_Border = 62
	ST_BorderConfetti               ST_Border = 63
	ST_BorderConfettiGrays          ST_Border = 64
	ST_BorderConfettiOutline        ST_Border = 65
	ST_BorderConfettiStreamers      ST_Border = 66
	ST_BorderConfettiWhite          ST_Border = 67
	ST_BorderCornerTriangles        ST_Border = 68
	ST_BorderCouponCutoutDashes     ST_Border = 69
	ST_BorderCouponCutoutDots       ST_Border = 70
	ST_BorderCrazyMaze              ST_Border = 71
	ST_BorderCreaturesButterfly     ST_Border = 72
	ST_BorderCreaturesFish          ST_Border = 73
	ST_BorderCreaturesInsects       ST_Border = 74
	ST_BorderCreaturesLadyBug       ST_Border = 75
	ST_BorderCrossStitch            ST_Border = 76
	ST_BorderCup                    ST_Border = 77
	ST_BorderDecoArch               ST_Border = 78
	ST_BorderDecoArchColor          ST_Border = 79
	ST_BorderDecoBlocks             ST_Border = 80
	ST_BorderDiamondsGray           ST_Border = 81
	ST_BorderDoubleD                ST_Border = 82
	ST_BorderDoubleDiamonds         ST_Border = 83
	ST_BorderEarth1                 ST_Border = 84
	ST_BorderEarth2                 ST_Border = 85
	ST_BorderEarth3                 ST_Border = 86
	ST_BorderEclipsingSquares1      ST_Border = 87
	ST_BorderEclipsingSquares2      ST_Border = 88
	ST_BorderEggsBlack              ST_Border = 89
	ST_BorderFans                   ST_Border = 90
	ST_BorderFilm                   ST_Border = 91
	ST_BorderFirecrackers           ST_Border = 92
	ST_BorderFlowersBlockPrint      ST_Border = 93
	ST_BorderFlowersDaisies         ST_Border = 94
	ST_BorderFlowersModern1         ST_Border = 95
	ST_BorderFlowersModern2         ST_Border = 96
	ST_BorderFlowersPansy           ST_Border = 97
	ST_BorderFlowersRedRose         ST_Border = 98
	ST_BorderFlowersRoses           ST_Border = 99
	ST_BorderFlowersTeacup          ST_Border = 100
	ST_BorderFlowersTiny            ST_Border = 101
	ST_BorderGems                   ST_Border = 102
	ST_BorderGingerbreadMan         ST_Border = 103
	ST_BorderGradient               ST_Border = 104
	ST_BorderHandmade1              ST_Border = 105
	ST_BorderHandmade2              ST_Border = 106
	ST_BorderHeartBalloon           ST_Border = 107
	ST_BorderHeartGray              ST_Border = 108
	ST_BorderHearts                 ST_Border = 109
	ST_BorderHeebieJeebies          ST_Border = 110
	ST_BorderHolly                  ST_Border = 111
	ST_BorderHouseFunky             ST_Border = 112
	ST_BorderHypnotic               ST_Border = 113
	ST_BorderIceCreamCones          ST_Border = 114
	ST_BorderLightBulb              ST_Border = 115
	ST_BorderLightning1             ST_Border = 116
	ST_BorderLightning2             ST_Border = 117
	ST_BorderMapPins                ST_Border = 118
	ST_BorderMapleLeaf              ST_Border = 119
	ST_BorderMapleMuffins           ST_Border = 120
	ST_BorderMarquee                ST_Border = 121
	ST_BorderMarqueeToothed         ST_Border = 122
	ST_BorderMoons                  ST_Border = 123
	ST_BorderMosaic                 ST_Border = 124
	ST_BorderMusicNotes             ST_Border = 125
	ST_BorderNorthwest              ST_Border = 126
	ST_BorderOvals                  ST_Border = 127
	ST_BorderPackages               ST_Border = 128
	ST_BorderPalmsBlack             ST_Border = 129
	ST_BorderPalmsColor             ST_Border = 130
	ST_BorderPaperClips             ST_Border = 131
	ST_BorderPapyrus                ST_Border = 132
	ST_BorderPartyFavor             ST_Border = 133
	ST_BorderPartyGlass             ST_Border = 134
	ST_BorderPencils                ST_Border = 135
	ST_BorderPeople                 ST_Border = 136
	ST_BorderPeopleWaving           ST_Border = 137
	ST_BorderPeopleHats             ST_Border = 138
	ST_BorderPoinsettias            ST_Border = 139
	ST_BorderPostageStamp           ST_Border = 140
	ST_BorderPumpkin1               ST_Border = 141
	ST_BorderPushPinNote2           ST_Border = 142
	ST_BorderPushPinNote1           ST_Border = 143
	ST_BorderPyramids               ST_Border = 144
	ST_BorderPyramidsAbove          ST_Border = 145
	ST_BorderQuadrants              ST_Border = 146
	ST_BorderRings                  ST_Border = 147
	ST_BorderSafari                 ST_Border = 148
	ST_BorderSawtooth               ST_Border = 149
	ST_BorderSawtoothGray           ST_Border = 150
	ST_BorderScaredCat              ST_Border = 151
	ST_BorderSeattle                ST_Border = 152
	ST_BorderShadowedSquares        ST_Border = 153
	ST_BorderSharksTeeth            ST_Border = 154
	ST_BorderShorebirdTracks        ST_Border = 155
	ST_BorderSkyrocket              ST_Border = 156
	ST_BorderSnowflakeFancy         ST_Border = 157
	ST_BorderSnowflakes             ST_Border = 158
	ST_BorderSombrero               ST_Border = 159
	ST_BorderSouthwest              ST_Border = 160
	ST_BorderStars                  ST_Border = 161
	ST_BorderStarsTop               ST_Border = 162
	ST_BorderStars3d                ST_Border = 163
	ST_BorderStarsBlack             ST_Border = 164
	ST_BorderStarsShadowed          ST_Border = 165
	ST_BorderSun                    ST_Border = 166
	ST_BorderSwirligig              ST_Border = 167
	ST_BorderTornPaper              ST_Border = 168
	ST_BorderTornPaperBlack         ST_Border = 169
	ST_BorderTrees                  ST_Border = 170
	ST_BorderTriangleParty          ST_Border = 171
	ST_BorderTriangles              ST_Border = 172
	ST_BorderTriangle1              ST_Border = 173
	ST_BorderTriangle2              ST_Border = 174
	ST_BorderTriangleCircle1        ST_Border = 175
	ST_BorderTriangleCircle2        ST_Border = 176
	ST_BorderShapes1                ST_Border = 177
	ST_BorderShapes2                ST_Border = 178
	ST_BorderTwistedLines1          ST_Border = 179
	ST_BorderTwistedLines2          ST_Border = 180
	ST_BorderVine                   ST_Border = 181
	ST_BorderWaveline               ST_Border = 182
	ST_BorderWeavingAngles          ST_Border = 183
	ST_BorderWeavingBraid           ST_Border = 184
	ST_BorderWeavingRibbon          ST_Border = 185
	ST_BorderWeavingStrips          ST_Border = 186
	ST_BorderWhiteFlowers           ST_Border = 187
	ST_BorderWoodwork               ST_Border = 188
	ST_BorderXIllusions             ST_Border = 189
	ST_BorderZanyTriangles          ST_Border = 190
	ST_BorderZigZag                 ST_Border = 191
	ST_BorderZigZagStitch           ST_Border = 192
	ST_BorderCustom                 ST_Border = 193
```
