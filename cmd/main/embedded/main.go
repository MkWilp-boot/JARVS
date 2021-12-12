package main

import (
	ENGINE "JARVS/cmd/packages/engines"
	"fmt"
)

var baseURLs = make(ENGINE.BaseURLs)
var conn = make(chan string)

func init() {
	baseURLs[ENGINE.Wikipedia] = "https://pt.wikipedia.org/"
}

func main() {
	Wikipedia := ENGINE.SearchEngine{
		Client: ENGINE.Wikipedia,
	}
	go Wikipedia.GoGet(baseURLs[Wikipedia.GetProvider()]+"wiki/São_Paulo", conn)

	for value := range conn {
		fmt.Print(value)
	}
}
