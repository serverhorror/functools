package functools_test

import (
	"fmt"
	"log"

	"github.com/serverhorror/functools"
)

func ExampleReduce() {
	reduce := func(out string, x int) (string, error) {
		s := fmt.Sprintf("%v%v", out, x)
		return s, nil
	}

	out, err := functools.Reduce(reduce, []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+#v", out)
	// Output: "123"

}
