# vmd-parser

`vmd-parser` is a parser for VMD files implemented in Go.

## How to use

```go
package main

import (
	"fmt"

	vmd "github.com/aki-0421/vmd-parser"
)

func main() {
	v, err := vmd.Parse("dance.vmd")
	if err != nil {
		panic(err)
	}

	fmt.Println(v.Header)
}
```

## Inspired by

| Project                                                          | License |
| ---------------------------------------------------------------- | ------- |
| [@takahirox/mmd-parser](https://github.com/takahirox/mmd-parser) | MIT     |

## Plan

[ ] Add tests
