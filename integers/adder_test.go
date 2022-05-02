package integers

import "testing"

func TestAdder(t *testing.T) {
	assertResut := func(t testing.TB, expected, sum int) {
		t.Helper()
		if expected != sum {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	}
	t.Run("Add numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertResut(t, expected, sum)

	})
}