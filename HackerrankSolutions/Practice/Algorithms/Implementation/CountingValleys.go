package implementation

import (
	"strings"
)

// CountingValleys returns completely passed valleys.
func CountingValleys(n int32, s string) int32 {
	var level int = 0
	var ar = strings.Split(s, "")
	var res int32 = 0

	for count := 0; count < int(n); count++ {
		if ar[count] == "D" {
			level--
		} else {
			level++
			if level == 0 {
				res++
			}
		}
	}

	return res
}
