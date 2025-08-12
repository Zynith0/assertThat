package assertThat

import (
	"testing"
)

func IsTrue(t *testing.T, want bool, got bool) {
	if want != got {
		t.Errorf("Wanted %v, but instead got %v", want, got)
	}
}

func IsEqualTo(t *testing.T, want, got any) {
	if want != got {
		t.Errorf("Wanted %v, but instead got %v", want, got)
	}
}
