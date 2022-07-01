package badges

import "testing"

func TestF(t *testing.T) {
	v := f()
	if v != 2 {
		t.Error("Fails")
	}
}
