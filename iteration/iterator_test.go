package iteration


import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeated("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}