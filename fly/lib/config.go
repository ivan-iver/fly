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
		panic(err)
	}
	//	fmt.Println("App | Config loaded successfully! \n")
	config.IsProduction = strings.EqualFold(config.Default("env"), production)
	return
}

// Returns configuration file path
func (self *Config) File() (file string) {
	return path.Join(self.Pwd, self.Filename)
}

// Gets config property from default section
func (self *Config) Default(property string) (result string) {
	result, _ = self.String("default", property)
	return
}
