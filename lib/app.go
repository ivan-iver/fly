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
)

var (
	hash = "5fab2d8"
	log  *Logger
	// Version is a variable
	version = "Fly Server version v0.0.1"
	logFile = "logs/error.log"
)

// App models current application
type App struct {
	app    *kingpin.Application
	Log    *Logger
	Config *Config
	*Server
}

// ConfigData models data setting from Config
type ConfigData struct {
	IsDebug     bool
	LogFile     string
	IndexFile   string
	DefaultPath string
}

func getConfigData() (config *Config, c *ConfigData, err error) {
	if config, err = NewConfig(); err != nil {
		return
	}
	c = &ConfigData{
		IsDebug:     config.BooleanDefault("debug", true),
		LogFile:     config.StringDefault("logfile", logFile),
		IndexFile:   config.StringDefault("index", "index.md"),
		DefaultPath: config.StringDefault("path", ""),
	}
	return config, c, err
}

// NewApp provides a new App struct with its initializated fields
func NewApp() (application *App, err error) {
	var config *Config
	var data *ConfigData
	if config, data, err = getConfigData(); err != nil {
		return
	}

	if log, err = GetLogger(data.LogFile, data.IsDebug); err != nil {
		return
	}
	log.Infof("ConfigData: %#v", data)
	application = &App{
		app:    kingpin.New(appName, desc),
		Config: config,
		Log:    log,
		Server: &Server{
			Index: data.IndexFile,
			Debug: data.IsDebug,
			Path:  data.DefaultPath,
		},
	}
	log.Debugf("Reading %v - Is debug: %v", application.Index, application.Debug)
	kingpin.Version(application.Version())
	application.required()
	return
}

// Version function returns the application version
func (a *App) Version() (v string) {
	log.Debugf("Version is %v, %v", version, hash)
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
func (a *App) Run() (err error) {
	err = a.Server.Run()
	return
}
