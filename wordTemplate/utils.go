package main

import (
	"fmt"

	"golang.org/x/net/html"
)

func in(value string, strArray []string) bool {
	for _, one := range strArray {
		if one == value {
			return true
		}
	}
	return false
}

func getAttributes(tokenizer *html.Tokenizer) map[string]string {
	kvStore := make(map[string]string, 0)
	attr := true
	for attr {
		var k []byte
		var v []byte
		k, v, attr = tokenizer.TagAttr()
		kvStore[fmt.Sprintf("%s", k)] = string(v)
	}
	return kvStore
}
