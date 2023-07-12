### Using
- **protoc-gen-go** `v1.31.0`
- **protoc** `23.4-osx-x86_64`

### Compiler command
> --go_opt define the package name for .proto file
```sh
$PATH/protoc \
--plugin=$PATH/protoc-gen-go \
--go_opt=MHero.proto=./hero \
--go_out=./ \
./Hero.proto
```