package implementation

//PageCount is the solution for Drawing Book
func PageCount(n int32, p int32) int32 {

	if n == 1 {
		n = 0
	} else {
		n = (n - (n % 2)) / 2
	}

	if p == 1 {
		p = 0
	} else {
		p = (p - (p % 2)) / 2
	}

	var res int32 = p

	if (p) > (n - p) {
		res = n - p
	}

	return res
}
