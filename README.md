# stringer

Stringer is a fuzzy string matching library for Go.

Supported algorithms:

* [Levenshtein distance](http://en.wikipedia.org/wiki/Levenshtein_distance)

## Installation

```go
go get github.com/zachlatta/stringer/stringer
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/zachlatta/stringer/stringer"
)

func main() {
	editDistance := stringer.LevenshteinDistance("foo", "bar")
	fmt.Println(editDistance)
}
```
