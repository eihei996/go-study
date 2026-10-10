package mathtool

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}

	if got := Add(-2, 3); got != 1 {
		t.Fatalf("Add(-2, 3) = %d, want 1", got)
	}
}

func TestSub(t *testing.T) {
	if got := Sub(5, 2); got != 3 {
		t.Fatalf("Sub(5, 2) = %d, want 3", got)
	}

	if got := Sub(2, 5); got != -3 {
		t.Fatalf("Sub(2, 5) = %d, want -3", got)
	}
}
