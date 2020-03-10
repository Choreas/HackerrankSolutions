package arrays

//RotLeft rotates an array a by d iterations.
func RotLeft(a []int32, d int32) []int32 {
	var last int32 = 0
	var length = len(a)
	var length32 = int32(length)
	if d >= length32 {
		d = d % length32
	}
	for rotation := int32(1); rotation <= d; rotation++ {
		last = a[0]
		for idx := range a {
			if idx == length-1 {
				a[idx] = last
			} else {
				a[idx] = a[idx+1]
			}
		}
	}

	return a
}
