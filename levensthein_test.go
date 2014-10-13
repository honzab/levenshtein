package levensthein

import "testing"

func TestDistance(t *testing.T) {
	s1 := "kitten"
	s2 := "sitting"
	if d := Distance(s1, s2); d != 3 {
		t.Fatalf("Distance %d is not expected %d", d, 3)
	}
}

func TestDistanceUnicode(t *testing.T) {
	s1 := "Smörgås"
	s2 := "smorgas"
	if d := Distance(s1, s2); d != 3 {
		t.Fatalf("Distance %d is not expected %d", d, 3)
	}
}

func TestDistanceEmptyStrings(t *testing.T) {
	var s1, s2 string
	if d := Distance(s1, s2); d != 0 {
		t.Fatalf("Distance %d is not expected %d", d, 0)
	}
}

func TestDistanceEmptyString(t *testing.T) {
	s1 := "Smörgås"
	s2 := ""
	if d := Distance(s1, s2); d != len(s1) {
		t.Fatalf("Distance %d is not expected %d", d, len(s2))
	}
	s1 = ""
	s2 = "Smörgås"
	if d := Distance(s1, s2); d != len(s2) {
		t.Fatalf("Distance %d is not expected %d", d, len(s2))
	}
}

func TestMinimum(t *testing.T) {
	if m := min(1, 2, 3); m != 1 {
		t.Fatalf("Expected minimum %s, not %s", m, 1)
	}
	if m := min(3, 2, 1); m != 1 {
		t.Fatalf("Expected minimum %s, not %s", m, 1)
	}
	if m := min(-3, 2, 1); m != -3 {
		t.Fatalf("Expected minimum %s, not %s", m, -3)
	}
	if m := min(1); m != 1 {
		t.Fatalf("Expected minimum %s, not %s", m, 1)
	}
}

func TestMinimumPanic(t *testing.T) {
	defer func() {
		if d := recover(); d == nil {
			t.Fatalf("Expected the function to panic, no panic found")
		}
	}()
	min()
}
