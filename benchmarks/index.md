## Make markdown file

```sh
go test ./... -bench=. | sh ../md_maker.sh
```

## Benchmarks


- [Factorial](./factorial/factorial.md)
- [Strings concat](./strings_concat/strings_concat.md)
- [Json encoder vs marshal](./json_encoder_vs_marshal/json_encoder_vs_marshal.md)
- [Chan vs mutex](./chan_vs_mutex/chan_vs_mutex.md)
- [For v2 range](./for_v2_range/for_v2_range.md)
- [Copy file](./copy_file/copy_file.md)
- [Json encoder vs write file](./json_encoder_vs_write_file/json_encoder_vs_write_file.md)
- [Json decoder vs read file](./json_decoder_vs_read_file/json_decoder_vs_read_file.md)
- [Switch default vs switch case](./switch_default_vs_switch_case/switch_default_vs_switch_case.md)
- [Mysql vs redis vs boltdb](./mysql_vs_redis_vs_boltdb/mysql_vs_redis_vs_boltdb.md)
- [Mysql bulk load](./mysql_bulk_load/mysql_bulk_load.md)
