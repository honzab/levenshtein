package golevensthein

import "math"

func LevenstheinDistance(s, t string) int64 {
	if s == t {
		return int64(0)
	}
	if len(s) == 0 {
		return int64(len(t))
	}
	if len(t) == 0 {
		return int64(len(s))
	}
	v0 := make([]int64, len(t)+1)
	v1 := make([]int64, len(t)+1)
	for i := range v1 {
		v1[i] = 0
	}
	for i := range v0 {
		v0[i] = int64(i)
	}
	for i := range s {
		v1[0] = int64(i) + 1
		for j := range t {
			var cost int64
			if string(s[i]) == string(t[j]) {
				cost = 0
			} else {
				cost = 1
			}
			v1[j+1] = int64(math.Min(float64(v1[j]+1), math.Min(float64(v0[j+1]+1), float64(v0[j]+cost))))
		}
		for j := range v0 {
			v0[j] = v1[j]
		}
	}
	return v1[len(t)]
}
