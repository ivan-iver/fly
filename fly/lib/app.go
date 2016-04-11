package lib

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	desc    = "Lightweight server with markdown support"
	portMsg = "Port number is required 8080 is default port"
	version = "Fly Server version v0.0.1"
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
		app:    kingpin.New("Fly", desc),
		Config: config,
		Log:    NewLogger(config),
		Server: &Server{},
	}
	application.required()
	return
}

func (a *App) Version() (v string) {
	var conf = a.Config.Default
	v = fmt.Sprintf("%v build(%s)", version, conf("app.version"))
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
