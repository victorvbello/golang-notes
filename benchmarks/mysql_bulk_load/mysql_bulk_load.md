## Mysql bulk load Results
Benchmark Name|Iterations|Per-Iteration
----|----|----
BenchmarkInsertOneByOne-12 | 466| 2695586 ns/op
BenchmarkInsertBlock10Stmt-12 | 39| 33669146 ns/op
BenchmarkInsertBlock20Stmt-12 | 20| 56446092 ns/op
BenchmarkInsertBlock100000-12 | 2| 809176019 ns/op
BenchmarkInsertLoadFile100000-12 | 3| 426861694 ns/op
BenchmarkInsertValidateBlock100000-12 | 1|1546866137 ns/op

Generated using go version go1.19.5 darwin/amd64

