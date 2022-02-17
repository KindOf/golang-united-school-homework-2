package square

import "testing"

func TestCalcSquare(t *testing.T) {
	type args struct {
		sideLen float64
		sideNum CustomInt
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Circle square",
			args: args{sideLen: 2, sideNum: 0},
			want: 12.566370614359172,
		}, {
			name: "Triangle square",
			args: args{sideLen: 2, sideNum: 3},
			want: 1.7320508075688772,
		}, {
			name: "Square area",
			args: args{sideLen: 2, sideNum: 4},
			want: 4,
		}, {
			name: "Undefined figure",
			args: args{sideLen: 2, sideNum: 2},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcSquare(tt.args.sideLen, tt.args.sideNum); got != tt.want {
				t.Errorf("CalcSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
