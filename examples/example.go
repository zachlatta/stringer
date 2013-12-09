package main

import (
	"fmt"

	"github.com/zachlatta/stringer/stringer"
)

func main() {
	editDistance := stringer.LevenshteinDistance("foo", "bar")
	fmt.Println(editDistance)
}
