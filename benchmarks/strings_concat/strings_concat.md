## Strings concat Results
Benchmark Name|Iterations|Per-Iteration
----|----|----
BenchmarkSprintf-12 |16658494| 71.60 ns/op
BenchmarkConcat-12 |1000000000| 0.2540 ns/op
BenchmarkJoin-12 |38999876| 31.34 ns/op
BenchmarkBuffer-12 |29760900| 39.69 ns/op

Generated using go version go1.17 darwin/amd64

