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

	var app *lib.App
	var err error
	if app, err = lib.NewApp(); err == nil {
		app.Parse(os.Args[1:])
	} else {
		log.Printf("Error creating app: %v", err)
		os.Exit(1)
	}
	if err = app.Run(); err != nil {
		log.Printf("Error while running app: %v", err)
		os.Exit(1)
	}
}
