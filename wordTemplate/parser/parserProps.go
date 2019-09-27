package parser

import (
	"github.com/unidoc/unioffice/document"
)

type GlobalConsts struct {
	h1 float64
	h2 float64
	h3 float64
	h4 float64
	h5 float64
	h6 float64
}

type parserState struct {
	currentPara      *document.Paragraph
	currentRun       *document.Run
	prev             *parserState
	section          Tags
	currentTag       Tags
	currentTextStyle *TextStyles
	currentList      *ListProps
	currentTable     *TableProps
	numDef           *document.NumberingDefinition
}

func NewParserState(currentState *parserState, tagName Tags) *parserState {
	newState := &parserState{}
	newState.prev = currentState
	newState.currentPara = currentState.currentPara
	newState.currentRun = currentState.currentRun
	newState.section = currentState.section
	newState.currentTag = tagName
	newState.currentTextStyle = currentState.currentTextStyle
	newState.numDef = currentState.numDef
	newState.currentList = currentState.currentList
	newState.currentTable = currentState.currentTable
	return newState
}
