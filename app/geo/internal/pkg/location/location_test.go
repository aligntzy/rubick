package location

import "testing"

func TestLocation(t *testing.T) {
	if _, err := Location("127.0.0.1"); err != nil {
		t.Error(err)
	}
}
