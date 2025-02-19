base10 encoding of primitive types

```bash
$ go test -bench=. -benchmem . 
goos: darwin
goarch: arm64
pkg: github.com/ndx-technologies/base10quant
cpu: Apple M3 Max
BenchmarkL9/string-16           93159283                12.71 ns/op           16 B/op          1 allocs/op
BenchmarkL9/from_string-16      149476885                8.018 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/ndx-technologies/base10quant 3.436s
```
