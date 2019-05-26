package berry

import (
	"testing"
)

func TestEnable(t *testing.T) {
	if enabled != true {
		t.Errorf("want default enabled true")
	}

	Enable(false)
	if enabled != false {
		t.Errorf("want enabled is false")
	}

	Enable(true)
	if enabled != true {
		t.Errorf("want enabled is true")
	}
}
