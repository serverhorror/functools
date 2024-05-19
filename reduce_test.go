package functools_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/serverhorror/functools"
)

func TestReduce_T1(t *testing.T) {
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
	test := testCase[int, string]{
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
	}
	got, err := functools.Reduce(test.args.fn, test.args.s)
	if (err != nil) != test.wantErr {
		t.Errorf("Reduce() error = %v, wantErr %v", err, test.wantErr)
		return
	}
	if !reflect.DeepEqual(got, test.want) {
		t.Errorf("Reduce() got = %#v, want %#v", got, test.want)
	}
}
