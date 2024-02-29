package foo

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Fatalf("1 + 2 == 3, but got %d", ans)
	}
}

func TestSubtract(t *testing.T) {
	if ans := Subtract(1, 2); ans != -1 {
		t.Fatalf("1 - 2 == -1, but got %d", ans)
	}
}
