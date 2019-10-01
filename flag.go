package main

import (
	"fmt"
	"flag"
)

func flagUsage() {
	fmt.Println("Usage: armada [-help] [args]")
	fmt.Println()
	fmt.Println("All available arguments are listed below:")
	fmt.Println()
	flag.PrintDefaults()
}