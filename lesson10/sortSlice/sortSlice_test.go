package sortSlice

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestSortSlice(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name   string
		args   args
		expect []int
	}{
		{
			name:   "1,2,3",
			args:   args{[]int{1, 2, 3}},
			expect: []int{1, 2, 3},
		},
		{
			name:   "3,2,1",
			args:   args{[]int{3, 2, 1}},
			expect: []int{1, 2, 3},
		},
		{
			name:   "-100,0,-300",
			args:   args{[]int{-100, 0, -300}},
			expect: []int{-300, -100, 0},
		},
		{
			name:   "-300,-400,-500",
			args:   args{[]int{-300, -400, -500}},
			expect: []int{-500, -400, -300},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstSlice := make([]int, len(tt.args.array))
			copy(firstSlice, tt.args.array)
			SortSlice(tt.args.array)
			for i, v := range tt.args.array {
				if v != tt.expect[i] {
					t.Errorf("После сортировки слайса %v получили %v, ожидали %v", firstSlice, tt.args.array, tt.expect)
				}
			}
		})
	}
}

func BenchmarkSortSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortSort([]int{100, 500, 900, 100, -100, 0, 245, 543, 6562, 4, 341, 43, 4, 1234, 32})
	}
}

func BenchmarkSortSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortSlice([]int{100, 500, 900, 100, -100, 0, 245, 543, 6562, 4, 341, 43, 4, 1234, 32})
	}
}

func ExampleSortSlice() {
	a := []int{5, 4, 3, 2, 6}
	SortSlice(a)
	fmt.Println(a)
	//Output: [2 3 4 5 6]
}
