package main

import (
	ENGINE "JARVS/cmd/packages/engines"
	"fmt"
)

func main() {
	mainEngine := ENGINE.SearchEngine{
		Client: ENGINE.Wikipedia,
	}

	fmt.Println(mainEngine.GetProvider())
}
