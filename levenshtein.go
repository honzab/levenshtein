package levenshtein

// This methods counts the Levenshtein distance https://en.wikipedia.org/wiki/Levenshtein_distance
// of two strings. It allows only for single character edits (insertion,
// deletion and substitution).
//
// The implementation is using the two matrix row method described on
// http://www.codeproject.com/Articles/13525/Fast-memory-efficient-Levenshtein-algorithm
// and on the Wikipedia (see the link above).
func Distance(s, t string) int {
	if s == t {
		return 0
	}
	len_s := len(s)
	len_t := len(t)
	if len_s == 0 {
		return len_t
	}
	if len_t == 0 {
		return len_s
	}
	v0 := make([]int, len_t+1)
	v1 := make([]int, len_t+1)
	for i := range v0 {
		v0[i] = int(i)
	}
	var cost int = 0
	for i := range s {
		v1[0] = int(i) + 1
		for j := range t {
			if string(s[i]) == string(t[j]) {
				cost = 0
			} else {
				cost = 1
			}
			v1[j+1] = min(v1[j]+1, v0[j+1]+1, v0[j]+cost)
		}
		for j := range v0 {
			v0[j] = v1[j]
		}
	}
	return int(v1[len_t])
}

// Helper variadic min function that takes a variable amount of integers and
// returns the minimal value. Panics if no values are supplied.
func min(ints ...int) int {
	if len(ints) == 0 {
		panic("Cannot compute minimum of no elements")
	}
	min := ints[0]
	for _, v := range ints {
		if v < min {
			min = v
		}
	}
	return min
}
