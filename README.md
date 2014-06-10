# go-textdistance

Calculate various text distancewith golang.

## Implemented
* [Levenshtein distance](http://en.wikipedia.org/wiki/Levenshtein_distance)
* [Damerau-Levenshtein distance](http://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance)
* [Jaro distance](http://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance)
* [Jaro-Winkler distance](http://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance)

## How to Use


```bash
go install github.com/masatana/go-textdistance
```

```go
package main

import (
	"fmt"

	"github.com/masatana/go-textdistance"
)

func main() {
	s1 := "this is a test"
	s2 := "that is a test"
	fmt.Println(textdistance.LevenshteinDistance(s1, s2))
	fmt.Println(textdistance.DamerauLevenshteinDistance(s1, s2))
	fmt.Println(textdistance.JaroDistance(s1, s2))
	fmt.Println(textdistance.JaroWinklerDistance(s1, s2))
}
```

## License

This software is released under the MIT License, see LICENSE.txt.
