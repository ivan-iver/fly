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
	hash    = "build:(bfdd056)"
)

type App struct {
	app    *kingpin.Application
	Log    *Logger
	Config *Config
	*Server
}

func NewApp() (application *App) {
	config, _ := NewConfig()
	application = &App{
		app:    kingpin.New(appName, desc),
		Config: config,
		Log:    NewLogger(config),
		Server: &Server{
			Index: config.StringDefault("index", "index.md"),
			Debug: config.BooleanDefault("debug", true),
			Path:  config.StringDefault("path", ""),
		},
	}
	application.Log.Printf("Server: %v - Debud: %v", application.Index, application.Debug)
	kingpin.Version(application.Version())
	application.required()
	return
}

func (a *App) Version() (v string) {
	v = fmt.Sprintf("%v %s", version, hash)
	a.app.Version(v)
	return
}

func (a *App) Parse(args []string) {
	kingpin.MustParse(a.app.Parse(args))
}

func (a *App) required() {
	a.app.Flag("port", portMsg).Short('P').Default("8080").StringVar(&a.Port)
}

func (a *App) Run() {
	if err := a.Server.Run(); err != nil {
		panic(err)
	}
}
