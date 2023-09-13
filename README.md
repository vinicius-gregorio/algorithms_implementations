# algorithms_implementations
Implementations of most common algorithms

## Searching Algorithms

### Linear Search
[Check here](./go/search/README.md)

## How to benchmark testing?

### In golang:


Go to go folder:
```bash
$ cd go
```
then:
```bash
$ go test -bench . ./benchmarks/
```

to run using only 1 cpu:
```bash
go test -bench . -cpu 1 ./benchmarks/
```