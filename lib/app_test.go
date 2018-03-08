package lib_test

import (
	"github.com/iver/fly/lib"
	"testing"
)

func TestNewApp(t *testing.T) {
	app, err := lib.NewApp()

	if err != nil {
		t.Errorf("NewApp gets error: %v", err)
	}

	if app == nil {
		t.Error("NewApp returns nil")
	}

	if app.Log == nil {
		t.Error("app.Log is nil")
	}

	if app.Config == nil {
		t.Error("app.Config is nil")
	}

	if app.Server == nil {
		t.Error("app.Server is nil")
	}

}
