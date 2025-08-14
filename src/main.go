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

func Contains(t *testing.T, want any, got []any) {
	for i := 0; i < len(got); i++ {
		if want != got[i] {
			t.Errorf("%v does not contain %v", got, want)
		}
	}
}
