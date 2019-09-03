package main

import (
	"github.com/unidoc/unioffice/document"
)

type parserState struct {
	currentPara      *document.Paragraph
	currentRun       *document.Run
	prev             *parserState
	section          Tags
	currentTag       Tags
	currentTextStyle *TextStyles
	currentList      *ListProps
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
	return newState
}

type ListProps struct {
	level       int
	indentDelta int
	numDefLevel int
}
