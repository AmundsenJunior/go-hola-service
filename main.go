package main

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
}

func main() {
	a := App{}

	a.Initialize()

	a.Run(":8000")
}
