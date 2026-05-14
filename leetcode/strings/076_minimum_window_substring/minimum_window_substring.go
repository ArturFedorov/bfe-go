package minimum_window_substring

import "math"

func minWindow(s string, t string) string {
	if len(s) < len(t) || len(t) == 0 {
		return ""
	}

	freqMap := [128]int{}
	count := len(t)
	start, minStart, minLen := 0, 0, math.MaxInt32

	for _, c := range t {
		freqMap[c]++
	}

	for end := 0; end < len(s); end++ {
		if freqMap[s[end]] > 0 {
			count--
		}

		freqMap[s[end]]--

		for count == 0 {
			if end-start+1 < minLen {
				minStart = start
				minLen = end - start + 1
			}

			freqMap[s[start]]++
			if freqMap[s[start]] > 0 {
				count++
			}

			start++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return s[minStart : minStart+minLen]
}
