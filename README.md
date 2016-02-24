# FastRand
FastRand generate pseudo random int between 0 and n. Argument n must be power of 2. Average working time of FastRand is 6ns.

[![Build Status](https://travis-ci.org/miolini/fastrand.svg)](https://travis-ci.org/miolini/fastrand) [![Go Report Card](http://goreportcard.com/badge/miolini/fastrand)](http://goreportcard.com/report/miolini/fastrand)

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
