package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Wikindexer is a tool for finding different paths from one article to another.")
		fmt.Println()
		fmt.Println("(It is now implementation of WikiRace game)")
	}
}
