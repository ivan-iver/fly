package main

import (
	"fly/lib"
	"os"
)

func main() {
	var app = lib.NewApp()
	app.Parse(os.Args[1:])
	app.Run()
}
