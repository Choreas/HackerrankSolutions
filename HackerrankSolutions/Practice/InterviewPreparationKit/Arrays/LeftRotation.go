package arrays

//RotLeft rotates an array a by d iterations.
func RotLeft(a []int32, d int32) []int32 {
	var length = len(a)
	var length32 = int32(length)
	var b []int32
	if d >= length32 {
		d = d % length32
	}

	for idx := range a {
		var newidx = d + int32(idx)
		if newidx >= length32 {
			newidx -= length32
		}

		b = append(b, a[newidx])

	}

	return b
}
