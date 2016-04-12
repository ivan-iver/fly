package lib

import (
	"fmt"
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
		fmt.Errorf("| Error | %v \n", err)
		panic(err)
	}

	var file = config.File()
	//	fmt.Printf("App | Config will be loaded from %v \n", file)
	if config.Config, err = c.ReadDefault(file); err != nil {
		fmt.Errorf("| Error | %v \n", err)
		config.setDefault()
		return
	}
	//	fmt.Println("App | Config loaded successfully! \n")
	config.IsProduction = strings.EqualFold(config.Default("env"), production)
	// fmt.Println("App | Config.IsProduction \n", config.IsProduction)
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
	fmt.Printf("Property: %v", property)
	fmt.Printf("En Default %v", c)
	if result, err = c.String("default", property); err != nil {
		fmt.Errorf("| Error | %v \n", err)
		return ""
	}
	return
}

// Gets config property from default section or use default value
func (c *Config) StringDefault(property string, strDefault string) (result string) {
	var err error
	fmt.Printf("Property: %v", property)
	fmt.Printf("En Default %v", c)
	if result, err = c.String("default", property); err != nil {
		fmt.Errorf("| Error | %v \n", err)
		return ""
	}
	return
}
