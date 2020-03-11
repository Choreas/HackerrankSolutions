package strings

import (
	"strings"
)

var charlist = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func morganAndString(a string, b string) string {
	var combined = len(a) + len(b)
	var sa = strings.Split(a, "")
	var sb = strings.Split(b, "")
	var empty = false
	var res []string

	for ct := 0; ct <= combined-1 && empty == false; ct++ {
		var la = len(sa)
		var lb = len(sb)
		char, src := chooseLex(sa, sb, la, lb, 0)
		if src == "a" {
			sa = sa[1:]
			if la-1 == 0 {
				empty = true
			}
		} else {
			sb = sb[1:]
			if lb-1 == 0 {
				empty = true
			}
		}
		res = append(res, char)
		// fmt.Println(res)
	}
	if len(sa) > 0 {
		for _, val := range sa {
			res = append(res, val)
		}
	} else if len(sb) > 0 {
		for _, val := range sb {
			res = append(res, val)
		}
	}
	return strings.Join(res, "")
}

func chooseLex(a []string, b []string, la int, lb int, idx int) (char string, src string) {

	var ca = a[idx]
	var cb = b[idx]
	var ia = strings.Index(charlist, ca)
	var ib = strings.Index(charlist, cb)

	if ib == ia {
		if la > idx+1 && lb > idx+1 {
			_, src = chooseLex(a, b, la, lb, idx+1)
			if src == "a" {
				char = ca
			} else {
				char = cb
			}
		} else if la < lb {
			src = "b"
			char = cb
		} else {
			src = "a"
			char = ca
		}
	} else if ia > ib {
		src = "b"
		char = cb
	} else {
		src = "a"
		char = ca
	}

	return char, src
}
