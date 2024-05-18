package functools_test

import (
	"fmt"
	"github.com/serverhorror/functools"
	"log"
	"strconv"
)

func ExampleFilter() {
	filter := func(i int) (bool, error) {
		return (i % 2) == 0, nil
	}
	out, err := functools.Filter(filter, []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	// Output: [2]
}

func Example_stringToInt_Map() {
	mapper := func(s string) (int, error) {
		return strconv.Atoi(s)
	}
	out, err := functools.Map(mapper, []string{"1", "2", "3"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	// Output: [1 2 3]
}

func Example_intToString_Map() {
	mapper := func(i int) (string, error) {
		return fmt.Sprintf("%v", i), nil
	}

	out, err := functools.Map(mapper, []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	// Output: [1 2 3]

}
func ExampleReduce() {
	reduce := func(out string, x int) (string, error) {
		s := fmt.Sprintf("%v%v", out, x)
		return s, nil
	}

	out, err := functools.Reduce(reduce, []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
	// Output: 123

}
