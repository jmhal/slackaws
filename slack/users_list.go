package main

import (
	"github.com/bluele/slack"
	"fmt"
)

func main() {
	use := Lista(token)
	for i := 0; i<len(use); i++ {
		fmt.Println(use[i])
	}
}
