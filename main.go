package main

import (
	"github.com/iver/fly/lib"
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

	if app, err := lib.NewApp(); err == nil {
		app.Parse(os.Args[1:])
		app.Run()
	} else {
		os.Exit(1)
	}
}
