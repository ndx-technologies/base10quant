base10 encoding of primitive types

```bash
$ go test -bench=. -benchmem . 
goos: darwin
goarch: arm64
pkg: github.com/ndx-technologies/base10quant
cpu: Apple M3 Max
BenchmarkL9/string-16           96114614                12.32 ns/op           16 B/op          1 allocs/op
BenchmarkL9/from_string-16      160158512                7.478 ns/op           0 B/op          0 allocs/op
```
