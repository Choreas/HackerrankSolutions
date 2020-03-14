package arrays

//HourglassSum gets the highest hourglass sum.
func HourglassSum(arr [][]int32) int32 {
	var highestSum int32 = -1000000
	var length = len(arr) - 1
	for rowid, row := range arr {

		var sum int32 = 0
		for idx := range row {
			sum = arr[rowid][idx] + arr[rowid][idx+1] + arr[rowid][idx+2] + arr[rowid+1][idx+1] + arr[rowid+2][idx] + arr[rowid+2][idx+1] + arr[rowid+2][idx+2]
			if sum > highestSum {
				highestSum = sum
			}
			if idx == length-2 {
				break
			}
		}
		if rowid == length-2 {
			break
		}
	}
	return highestSum
}
