# stringer

![Go](https://github.com/hemantjadon/stringer/workflows/stringer/badge.svg?branch=master&event=push)

Package stringer provides `fmt.Stringer` wrappers to `go` builtin types.

```go
var v fmt.Stringer

v = stringer.Int64(1729)

fmt.Println(v)
// Output: 1729
```
