package functools

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReduce_T1(t *testing.T) {
	t.Parallel()
	type args[T any, Out any] struct {
		fn func(Out, T) (Out, error)
		s  []T
	}
	type testCase[T any, Out any] struct {
		name    string
		args    args[T, Out]
		want    Out
		wantErr bool
	}
	tests := []testCase[int, string]{
		{
			name: "concatenate ints to string",
			args: args[int, string]{
				fn: func(out string, x int) (string, error) {
					s := fmt.Sprintf("%v%v", out, x)
					return s, nil
				},
				s: []int{1, 2, 3},
			},
			want:    "123",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reduce(tt.args.fn, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reduce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}
