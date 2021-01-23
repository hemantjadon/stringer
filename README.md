# stringer

![Go](https://github.com/hemantjadon/stringer/workflows/stringer/badge.svg?branch=master&event=push)
[![Go Reference](https://pkg.go.dev/badge/github.com/hemantjadon/stringer.svg)](https://pkg.go.dev/github.com/hemantjadon/stringer)

Package stringer provides `fmt.Stringer` wrappers to `go` builtin types.

```go
var v fmt.Stringer

v = stringer.Int64(1729)

fmt.Println(v)
// Output: 1729
```
