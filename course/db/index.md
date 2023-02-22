
## Test using mariaDB
### Understand LastInserte flow
> This test was create to get more context of flow around get last inserted id,
to do this i created the methods `getLastInserteIdUsingConcurrency` and `getLastInserteIdUsingConcurrencyAndDBStmt`

> The result of this test was very interesting for me, in short, when you run an insert query using `Exec` method,
the flow receive from de server one response on binary, this response if is OK, contain the rows affected and if is a single
insert contains de las insert id

> With this result it is determined that inserts can be executed using concurrencies without affecting the result of the last inserted id