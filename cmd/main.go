package main

import (
	"github.com/mike-carey/tfgenvars"
	"os"
)

func main() {
	err := tfgenvars.Run(os.Stdin, os.Stdout, os.Args[1:])
	if err != nil {
		panic(err)
	}
}
