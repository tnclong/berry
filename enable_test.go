package berry

import (
	"os"
	"testing"
)

func TestEnable(t *testing.T) {
	Enable(false)
	if enabled != false {
		t.Errorf("want enabled is false")
	}

	Enable(true)
	if enabled != true {
		t.Errorf("want enabled is true")
	}

	Enable(true)
	os.Setenv("TERM", "dump")
	ini()
	if enabled != false {
		t.Errorf("want enabled is false")
	}

	Enable(false)
	os.Setenv("CLICOLOR_FORCE", "1")
	ini()
	if enabled != true {
		t.Errorf("want enabled is true")
	}
}
