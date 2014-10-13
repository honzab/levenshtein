package levensthein

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
