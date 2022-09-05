package fibbonachi

import (
	"os"
	"testing"
)

//Не понял для чего и как использовать TestMain
func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func Test_fibbonachi(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{1},
			want: 1,
		},
		{
			name: "3",
			args: args{3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibbonachi(tt.args.n); got != tt.want {
				t.Errorf("fibonachi() = %v, want %v", got, tt.want)
			}
		})
	}
}
