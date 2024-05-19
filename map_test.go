package functools

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestMap_StringToInt(t *testing.T) {
	type args[TIn any, TOut any] struct {
		fn func(TIn) (TOut, error)
		s  []TIn
	}
	type testCase[TIn any, TOut any] struct {
		name    string
		args    args[TIn, TOut]
		want    []TOut
		wantErr bool
	}
	tests := []testCase[string, int]{
		{
			name: "convert string to int",
			args: args[string, int]{
				fn: func(s string) (int, error) {
					return strconv.Atoi(s)
				},
				s: []string{"1", "2", "3"},
			},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Map(tt.args.fn, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Map() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMap_IntToString(t *testing.T) {
	type args[TIn any, TOut any] struct {
		fn func(TIn) (TOut, error)
		s  []TIn
	}
	type testCase[TIn any, TOut any] struct {
		name    string
		args    args[TIn, TOut]
		want    []TOut
		wantErr bool
	}
	test := testCase[int, string]{
		name: "convert int to string",
		args: args[int, string]{
			fn: func(i int) (string, error) {
				return fmt.Sprintf("%d", i), nil
			},
			s: []int{1, 2, 3},
		},
		want: []string{"1", "2", "3"},
	}
	got, err := Map(test.args.fn, test.args.s)
	if (err != nil) != test.wantErr {
		t.Errorf("Map() error = %v, wantErr %v", err, test.wantErr)
	}
	if !reflect.DeepEqual(got, test.want) {
		t.Errorf("Map() = %v, want %v", got, test.want)
	}
}
