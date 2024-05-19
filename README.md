# functools

[![Go Reference](https://pkg.go.dev/badge/github.com/serverhorror/functools.svg)](https://pkg.go.dev/github.com/serverhorror/functools)

`functools` is a zero dependency generic Go library that provides utility functions for functional programming patterns.
It includes functions for mapping, reducing, and filtering data.

## Features

- `Map`: Transforms a slice of data by applying a provided function to each element.
- `Reduce`: Reduces a slice of data to a single value by applying a provided function.
- `Filter`: Filters a slice of data according to a provided function.

## Usage

Here are some examples of how to use the `functools` library:

### Map

```go
mapper := func(i int) (string, error) {
  return fmt.Sprintf("%v", i), nil
}

out, err := functools.Map(mapper, []int{1, 2, 3})
if err != nil {
  log.Fatal(err)
}
fmt.Println(fmt.Sprintf("%+#v", out))
// Output: []string{"1", "2", "3"}
```

### Filter

```go
filter := func(i int) (bool, error) {
  return (i % 2) == 0, nil
}
out, err := functools.Filter(filter, []int{1, 2, 3})
if err != nil {
  log.Fatal(err)
}
fmt.Println(fmt.Sprintf("%+#v", out))
// Output: []int{2}
```

### Reduce

```go
reduce := func(out string, x int) (string, error) {
  s := fmt.Sprintf("%v%v", out, x)
  return s, nil
}

out, err := functools.Reduce(reduce, []int{1, 2, 3})
if err != nil {
  log.Fatal(err)
}

fmt.Println(fmt.Sprintf("%+#v", out))
// Output: "123"
```
