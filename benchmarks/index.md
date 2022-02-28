## Make markdown file

```sh
go test ./... -bench=. | sh ../md_maker.sh
```

## Benchmarks


- [Factorial](./factorial/factorial.md)
- [Strings concat](./strings_concat/strings_concat.md)
- [Json encoder vs marshal](./json_encoder_vs_marshal/json_encoder_vs_marshal.md)
- [Chan vs mutex](./chan_vs_mutex/chan_vs_mutex.md)
