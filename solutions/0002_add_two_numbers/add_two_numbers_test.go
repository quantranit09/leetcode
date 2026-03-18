package addtwonumbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "Example 1",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			name: "Example 2",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			name: "Example 3",
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := ConvertArrayToList(tt.l1)
			l2 := ConvertArrayToList(tt.l2)
			wantList := ConvertArrayToList(tt.want)

			got := AddTwoNumbers(l1, l2)

			// helper to convert list back to array for easier deep equal
			gotArr := listToArray(got)
			wantArr := listToArray(wantList)

			if !reflect.DeepEqual(gotArr, wantArr) {
				t.Errorf("AddTwoNumbers() = %v, want %v", gotArr, wantArr)
			}
		})
	}
}

func listToArray(l *ListNode) []int {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}
