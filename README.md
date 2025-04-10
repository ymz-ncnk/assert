# assert
Provides invariant checking in Golang code.

# How to use
```go
package main

import assert "github.com/ymz-ncnk/assert/panic"

// Turns on assertions.
func init() {
  assert.On = true
}

func main(){
  assert.Equal(1, 1)
}
```

There are also two packages `error` and `fatal` that can be used in tests:
```go
package main

import asserterror "github.com/ymz-ncnk/assert/error"

func TestSome(t *testing.T) {
  asserterror.Equal(1, 1, t)
}
```