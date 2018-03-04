package lib_test

import (
	"github.com/iver/fly/lib"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config, err := lib.NewConfig()
	if err != nil {
		t.Error("On NewConfig: ", err)
	}

	if len(config.Pwd) == 0 {
		t.Error("Config rooted path is empty")
	}

	if len(config.Filename) == 0 {
		t.Error("Config.Filename is empty")
	}

}
