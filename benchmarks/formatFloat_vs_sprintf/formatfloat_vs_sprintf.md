## FormatFloat vs sprintf Results
Benchmark Name|Iterations|Per-Iteration
----|----|----
BenchmarkFormatFloat-12 | 5556482| 229.9 ns/op
BenchmarkSprintf-12 | 4126485| 260.4 ns/op
BenchmarkFormatFloatMulti-12 | 1000000| 1044 ns/op
BenchmarkSprintfMulti-12 | 1000000| 1028 ns/op

Generated using go version go1.19.5 darwin/amd64

