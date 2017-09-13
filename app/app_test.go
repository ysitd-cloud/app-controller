package app

import "testing"

func TestAppInterface(t *testing.T) {
	testApp := func(a Application) bool {
		return true
	}

	if !testApp(&App{}) {
		t.Errorf("App fail to match interface")
	}
}
