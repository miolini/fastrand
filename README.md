# FastRand
FastRand generate pseudo random int between 0 and n. Argument n must be power of 2. Average working time of FastRand is 6ns.

# Example
```go
package main

import (
  "fmt"
  "github.com/miolini/fastrand"
)

func main() {
  for i:=0; i<10; i++ {
    fmt.Println(fastrand.FastRand(64))
  }
}
```
