package libexample

import "testing"

func TestSum(t *testing.T) {
	t.Parallel()

	type args struct {
		a int64
		b int64
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "sum",
			args: args{
				a: 2,
				b: 2,
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
