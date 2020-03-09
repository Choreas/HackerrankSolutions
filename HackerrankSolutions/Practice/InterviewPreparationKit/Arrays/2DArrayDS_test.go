package arrays

import "testing"

func Test_hourglassSum(t *testing.T) {
	type args struct {
		arr [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"Case 1", args{getArr(1)}, 19},
		{"Case 2", args{getArr(2)}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HourglassSum(tt.args.arr); got != tt.want {
				t.Errorf("hourglassSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getArr(testcase int) (out [][]int32) {
	switch testcase {
	case 1:
		{
			out = [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			}
		}
	case 2:
		{
			out = [][]int32{
				{-9, -9, -9, 1, 1, 1},
				{0, -9, 0, 4, 3, 2},
				{-9, -9, -9, 1, 2, 3},
				{0, 0, 8, 6, 6, 0},
				{0, 0, 0, -2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			}
		}
	}
	return out
}
