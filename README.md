# Fastrand
Fast random generator (6ns) without state from 0 to N. N must be power of 2.

# Example
```go
package main

import (
  "fmt"
  "github.com/miolini/fastrand"
)

func main() {
  for i:=0; i<10; i++ {
    fmt.Println(fastrand.Fastrand(64))
  }
}
