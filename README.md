# assert
Provides invariant checking in Golang code.

# How to use
```go
package main

import "github.com/ymz-ncnk/assert"

// First of all you should turn on assertions.
func init() {
  assert.On = true
}

func main(){
  assert.Equal(1,1)
}
```