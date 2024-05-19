package functools_test

import (
	"fmt"
	"log"
	"strconv"

	"github.com/serverhorror/functools"
)

func Example_stringToIntMap() {
	mapper := func(s string) (int, error) {
		return strconv.Atoi(s)
	}
	out, err := functools.Map(mapper, []string{"1", "2", "3"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+#v", out)
	// Output: []int{1, 2, 3}
}

func Example_intToStringMap() {
	mapper := func(i int) (string, error) {
		return fmt.Sprintf("%v", i), nil
	}

	out, err := functools.Map(mapper, []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+#v", out)
	// Output: []string{"1", "2", "3"}

}
