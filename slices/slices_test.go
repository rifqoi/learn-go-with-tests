package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		arrs []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "collection of 5 numbers",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "collection of any size",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: 45,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.args.arrs)
			if got != tt.want {
				t.Errorf("got: %d, want: %d", got, tt.want)
			}
		})
	}

}

func TestSumAll(t *testing.T) {
	type args struct {
		arrs [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "Sum All",
			args: args{[][]int{{1, 2}, {3, 4}, {5, 6}}},
			want: []int{3, 7, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumAll(tt.args.arrs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
