package functools

import (
	"reflect"
	"testing"
)

func TestFilterInt(t *testing.T) {
	t.Parallel()
	type args[T any] struct {
		f func(T) (bool, error)
		s []T
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    []T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name: "even numbers only",
			args: args[int]{
				f: func(i int) (bool, error) {
					return (i % 2) == 0, nil
				},
				s: []int{1, 2, 4, 5, 6},
			},
			want:    []int{2, 4, 6},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Filter(tt.args.f, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Filter() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
