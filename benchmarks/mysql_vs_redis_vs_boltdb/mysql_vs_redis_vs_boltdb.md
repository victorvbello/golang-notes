## Mysql vs redis vs boltdb Results
Benchmark Name|Iterations|Per-Iteration
----|----|----
BenchmarkInsert/mysql-12 | 894| 1378139 ns/op
BenchmarkInsert/redis-12 | 2653| 456566 ns/op
BenchmarkInsert/boltDB-12 | 28| 41491470 ns/op
BenchmarkGet/mysql-12 | 2056| 497171 ns/op
BenchmarkGet/redis-12 | 2520| 419636 ns/op
BenchmarkGet/boltDB-12 | 1598455| 745.5 ns/op
BenchmarkUpdate/mysql-12 | 268| 4774764 ns/op
BenchmarkUpdate/redis-12 | 2704| 534616 ns/op
BenchmarkUpdate/boltDB-12 | 30| 39360318 ns/op
BenchmarkDelete/mysql-12 | 1761| 574233 ns/op
BenchmarkDelete/redis-12 | 2869| 516632 ns/op
BenchmarkDelete/boltDB-12 | 30| 38770721 ns/op

Generated using go version go1.19.5 darwin/amd64

