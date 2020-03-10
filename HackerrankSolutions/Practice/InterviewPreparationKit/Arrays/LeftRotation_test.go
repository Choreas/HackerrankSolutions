package arrays

import (
	"reflect"
	"testing"
)

func Test_rotLeft(t *testing.T) {
	type args struct {
		a []int32
		d int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{"Case 1", args{[]int32{1, 2, 3, 4, 5}, 4}, []int32{5, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotLeft(tt.args.a, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
