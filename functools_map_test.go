package functools

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestMap_T1(t *testing.T) {
	t.Parallel()
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
func TestMap_T2(t *testing.T) {
	t.Parallel()
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
	tests := []testCase[int, string]{
		{
			name: "convert int to string",
			args: args[int, string]{
				fn: func(i int) (string, error) {
					return fmt.Sprintf("%d", i), nil
				},
				s: []int{1, 2, 3},
			},
			want: []string{"1", "2", "3"},
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
