package main

import (
	"github.com/ivan-iver/fly/lib"
	"log"
	"os"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("| Fatal Error | %v", r)
			log.Printf("| Fatal Error | Stack:  %v", string(debug.Stack()))
		}
	}()
	var app = lib.NewApp()
	app.Parse(os.Args[1:])
	app.Run()
}
