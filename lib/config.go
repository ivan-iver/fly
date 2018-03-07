package lib

import (
	// l "log"
	"os"
	"path"
	"strings"

	conf "bitbucket.org/ivan-iver/config"
)

var (
	production = `production`
	filename   = `app.conf`
)

// Config models an entity config reader
type Config struct {
	Pwd          string
	Filename     string
	IsProduction bool
	*conf.Config
}

// NewConfig creates a config struct.
func NewConfig() (config *Config, err error) {
	config = &Config{Filename: filename}
	if config.Pwd, err = os.Getwd(); err != nil {
		config.setDefault(err)
		return
	}

	var file = config.File()
	if config.Config, err = conf.ReadDefault(file); err != nil {
		config.setDefault(err)
		return
	}
	config.IsProduction = strings.EqualFold(config.Default("env"), production)
	return
}

// Set default values when config file does not exists
func (c *Config) setDefault(err error) {
	// l.Println("Config file missing. Using default settings. ", err)
	c.Config = &conf.Config{}
	c.IsProduction = false
}

// File function  returns configuration file path
func (c *Config) File() (file string) {
	return path.Join(c.Pwd, c.Filename)
}

// Default function gets config property from default section
func (c *Config) Default(property string) (result string) {
	var err error
	if result, err = c.String("default", property); err != nil {
		// l.Println(err)
		return ""
	}
	return
}

// StringDefault gets config property from default section or use default value
func (c *Config) StringDefault(property string, strDefault string) (result string) {
	var err error
	if result, err = c.String("default", property); err != nil {
		// l.Println(err)
		return strDefault
	}
	return
}

// BooleanDefault function gets config property from default section or use default boolean value
func (c *Config) BooleanDefault(property string, boolDefault bool) (result bool) {
	var err error
	if result, err = c.Bool("default", property); err != nil {
		// l.Println(err)
		return boolDefault
	}
	return
}
