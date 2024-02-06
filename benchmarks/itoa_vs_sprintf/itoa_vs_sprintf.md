## Itoa vs sprintf Results
Benchmark Name|Iterations|Per-Iteration
----|----|----
BenchmarkItoa-12 |37211126| 29.83 ns/op
BenchmarkSprintf-12 |12495006| 98.96 ns/op
BenchmarkItoaMulti-12 | 7679972| 146.3 ns/op
BenchmarkSprintfMulti-12 | 3771542| 300.8 ns/op

Generated using go version go1.19.5 darwin/amd64

