## Mysql bulk load Results
Benchmark Name|Iterations|Per-Iteration
----|----|----
BenchmarkInsertOneByOne-12 | 444| 2756142 ns/op
BenchmarkInsertBlock10Stmt-12 | 42| 29446363 ns/op
BenchmarkInsertBlock20Stmt-12 | 21| 54889352 ns/op
BenchmarkInsertBlock100000-12 | 2| 743058785 ns/op
BenchmarkInsertValidateBlock100000-12 | 1|1740333863 ns/op

Generated using go version go1.19.5 darwin/amd64

