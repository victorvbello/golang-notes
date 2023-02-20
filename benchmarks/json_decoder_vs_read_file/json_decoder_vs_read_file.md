## Json decoder vs read file Results
Benchmark Name|Iterations|Per-Iteration
----|----|----
BenchmarkJSONDecoder-12 | 110452| 11065 ns/op
BenchmarkReadFileWithUnmarshal-12 | 95091| 12502 ns/op

Generated using go version go1.17.8 darwin/amd64

