package services

import "testing"

func TestNilai(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case return C",
			args: args{
				x: 40,
			},
			want: "C",
		},
		{
			name: "case return B",
			args: args{
				x: 70,
			},
			want: "B",
		},
		{
			name: "case return A",
			args: args{
				x: 90,
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Nilai(tt.args.x); got != tt.want {
				t.Errorf("Nilai() = %v, want %v", got, tt.want)
			}
		})
	}
}
