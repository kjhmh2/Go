package quickSort

import "testing"

func TestQuickSort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 5, 4}, []int{1, 2, 4, 5}},
		{[]int{8, 9, 8}, []int{8, 8, 9}},
		{[]int{1, 5, 2, 6, 3, 0, 2, 5, 6}, []int{0, 1, 2, 2, 3, 5, 5, 6, 6}},
	}
	for _, c := range cases {
		var array = c.in
		quickSort(array, 0, len(array) - 1)
		for i,_ := range array {
			if array[i] != c.want[i] {
				t.Errorf("Wrong Answer!")
				break
			}
		}
	}
}

