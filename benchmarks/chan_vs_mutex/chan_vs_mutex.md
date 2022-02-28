## Chan vs mutex Results
Benchmark Name|Iterations|Per-Iteration
----|----|----
BenchmarkChan-12 | 20343| 58955 ns/op
BenchmarkMutexSlice-12 | 19040| 62695 ns/op
BenchmarkMutexArray-12 | 21007| 56911 ns/op

Generated using go version go1.17 darwin/amd64

