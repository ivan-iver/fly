// Package lib contains all project logic.
package lib

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	appName = "Fly"
	desc    = "Lightweight server with markdown support"
	portMsg = "Port number is required 8080 is default port"
	version = "Fly Server version v0.0.1"
)

var hash string
var log *Logger

// App models current application
type App struct {
	app    *kingpin.Application
	Log    *Logger
	Config *Config
	*Server
}

// NewApp provides a new App struct with its initializated fields
func NewApp() (application *App) {
	log = GetLogger()
	config, _ := NewConfig()
	application = &App{
		app:    kingpin.New(appName, desc),
		Config: config,
		Log:    log,
		Server: &Server{
			Index: config.StringDefault("index", "index.md"),
			Debug: config.BooleanDefault("debug", true),
			Path:  config.StringDefault("path", ""),
		},
	}
	//log = application.Log
	log.Infof("Reading %v - Is debug: %v", application.Index, application.Debug)
	kingpin.Version(application.Version())
	application.required()
	return
}

// Version function returns the application version
func (a *App) Version() (v string) {
	v = fmt.Sprintf("%v %s", version, hash)
	a.app.Version(v)
	return
}

// Parse function process the arguments
func (a *App) Parse(args []string) {
	kingpin.MustParse(a.app.Parse(args))
}

func (a *App) required() {
	a.app.Flag("port", portMsg).Short('P').Default("8080").StringVar(&a.Port)
}

// Run is invoked to start the server application logic.
func (a *App) Run() {
	if err := a.Server.Run(); err != nil {
		panic(err)
	}
}
