package golevensthein

import "testing"

func TestDistance(t *testing.T) {
	s1 := "kitten"
	s2 := "sitting"
	if d := LevenstheinDistance(s1, s2); d != 3 {
		t.Fatalf("Distance %d is not expected %d", d, 3)
	}
}

func TestDistanceUnicode(t *testing.T) {
	s1 := "Smörgås"
	s2 := "smorgas"
	if d := LevenstheinDistance(s1, s2); d != 3 {
		t.Fatalf("Distance %d is not expected %d", d, 3)
	}
}

func TestDistanceEmptyStrings(t *testing.T) {
	var s1, s2 string
	if d := LevenstheinDistance(s1, s2); d != 0 {
		t.Fatalf("Distance %d is not expected %d", d, 0)
	}
}

func TestDistanceEmptyString(t *testing.T) {
	s1 := "Smörgås"
	s2 := ""
	if d := LevenstheinDistance(s1, s2); d != len(s1) {
		t.Fatalf("Distance %d is not expected %d", d, len(s2))
	}
	s1 = ""
	s2 = "Smörgås"
	if d := LevenstheinDistance(s1, s2); d != len(s2) {
		t.Fatalf("Distance %d is not expected %d", d, len(s2))
	}
}
