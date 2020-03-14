package warmupchallenges

import (
	"math"
	"strings"
)

//RepeatedString solution.
func RepeatedString(s string, n int64) int64 {
	var res int64 = 0
	if s == "" {
		return 0
	}
	var factor = int64(math.Floor(float64(n / int64(len(s)))))
	var modulo = n % int64(len(s))

	res = int64(strings.Count(s, "a")) * factor
	res += int64(strings.Count(s[0:modulo], "a"))

	return res
}
