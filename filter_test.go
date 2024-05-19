package functools_test

import (
	"reflect"
	"testing"

	"github.com/serverhorror/functools"
)

func TestFilterInt(t *testing.T) {
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
	test := testCase[int]{
		name: "even numbers only",
		args: args[int]{
			f: func(i int) (bool, error) {
				return (i % 2) == 0, nil
			},
			s: []int{1, 2, 4, 5, 6},
		},
		want:    []int{2, 4, 6},
		wantErr: false,
	}

	got, err := functools.Filter(test.args.f, test.args.s)
	if (err != nil) != test.wantErr {
		t.Errorf("Filter() error = %v, wantErr %v", err, test.wantErr)
	}
	if !reflect.DeepEqual(got, test.want) {
		t.Errorf("Filter() got = %v, want %v", got, test.want)
	}
}
