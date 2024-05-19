package functools_test

import (
	"fmt"
	"log"

	"github.com/serverhorror/functools"
)

func ExampleMap() {
	mapper := func(i int) (string, error) {
		lookup := map[int]string{
			1: "z",
			2: "y",
			3: "x",
		}
		return lookup[i], nil
	}

	out, err := functools.Map(mapper, []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+#v\n", out)
	// Output: []string{"z", "y", "x"}

}
