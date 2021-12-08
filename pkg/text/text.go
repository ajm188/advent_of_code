package text

import "sort"

// Permutations returns the list of all permutations of s.
func Permutations(s string) []string {
	return permutations("", s)
}

func permutations(prefix string, s string) (result []string) {
	if len(s) == 0 {
		return []string{prefix}
	}

	for i, r := range s {
		result = append(result, permutations(prefix+string([]rune{r}), s[:i]+s[i+1:])...)
	}
	return result
}

// SortString returns a string containing the characters in s sorted
// lexicographically.
func SortString(s string) string {
	rs := []rune(s)
	sort.Slice(rs, func(i, j int) bool { return rs[i] < rs[j] })
	return string(rs)
}
