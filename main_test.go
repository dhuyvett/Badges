package badges

import "testing"

func TestF(t *testing.T) {
	v := f()
	if v != 2 {
		t.Error("Fails")
	}
}

func TestG(t *testing.T) {
	v := g()
	if v != 4 {
		t.Error("Fails")
	}
}
