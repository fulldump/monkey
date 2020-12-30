# Monkey

This repo is just a personal exercise to put in practice *Writing an interpreter in go*
written by **Thorsten Ball**. (https://www.amazon.com/Writing-Interpreter-Go-Thorsten-Ball/dp/3982016118)

<img src="https://images-na.ssl-images-amazon.com/images/I/31C139bzhML._SX348_BO1,204,203,200_.jpg">

## License

This software is published under MIT License except in cases of intelectual property conflict
with the original author **Thorsten Ball**.

I also would like to declare that, although software architecture and syntax is quite
similar to the Author's work, it has been written 100% from scratch without copying the code.

## Disclaimer

Consider this a mere software exercise, there is no warranty of any kind at all.


## How to use

You can compile REPL and get a binary inside `./bin/` folder with:

```shell
make repl
```

or

```shell
go build -o bin/repl ./cmd/repl/*
```

## How to test

```shell
make test
```

or

```shell
go test ./...
```

