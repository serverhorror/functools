package functools_test

import (
	"fmt"
	"log"

	"github.com/serverhorror/functools"
)

func ExampleFilter() {
	filter := func(i int) (bool, error) {
		return (i % 2) == 0, nil
	}
	out, err := functools.Filter(filter, []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+#v", out)
	// Output: []int{2}
}
