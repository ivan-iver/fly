package lib

import (
	"log"
	"os"
	"path"
	"strings"

	c "bitbucket.org/ivan-iver/config"
)

var (
	production = `production`
	filename   = `app.conf`
)

// Models an entity config reader
type Config struct {
	Pwd          string
	Filename     string
	IsProduction bool
	*c.Config
}

// Creates a config struct
func NewConfig() (config *Config, err error) {
	config = &Config{Filename: filename}
	if config.Pwd, err = os.Getwd(); err != nil {
		log.Fatalf("| Error | %v \n", err)
		panic(err)
	}

	var file = config.File()
	//	log.Printf("App | Config will be loaded from %v \n", file)
	if config.Config, err = c.ReadDefault(file); err != nil {
		log.Fatalf("| Error | %v \n", err)
		config.setDefault()
		return
	}
	//	log.Println("App | Config loaded successfully! \n")
	config.IsProduction = strings.EqualFold(config.Default("env"), production)
	//log.Println("App | Config.IsProduction ", config.IsProduction)
	return
}

// Set default values when config file does not exists
func (x *Config) setDefault() {
	x.Config = &c.Config{}
	x.IsProduction = false
}

// Returns configuration file path
func (c *Config) File() (file string) {
	return path.Join(c.Pwd, c.Filename)
}

// Gets config property from default section
func (c *Config) Default(property string) (result string) {
	var err error
	//log.Printf("App | Property: %v \n", property)
	if result, err = c.String("default", property); err != nil {
		log.Fatalf("| Error | %v \n", err)
		return ""
	}
	//log.Printf("App | Value: %v \n", result)
	return
}

// Gets config property from default section or use default value
func (c *Config) StringDefault(property string, strDefault string) (result string) {
	var err error
	//log.Printf("Config | Property: %v", property)
	if result, err = c.String("default", property); err != nil {
		log.Fatalf("| Error | %v \n", err)
		return strDefault
	}
	//log.Printf("Config | Value: %v \n", result)
	return
}

// Gets config property from default section or use default boolean value
func (c *Config) BooleanDefault(property string, boolDefault bool) (result bool) {
	var err error
	// log.Printf("Config | Property: %v", property)
	if result, err = c.Bool("default", property); err != nil {
		log.Fatalf("| Error | %v \n", err)
		return boolDefault
	}
	//log.Printf("Config | Value: %v \n", result)
	return
}
