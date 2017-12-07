package lib

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp()

	if app == nil {
		t.Error("NewApp must be not null")
	}

	if app.Log == nil {
		t.Error("App.Log must be not null")
	}
}
