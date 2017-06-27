This repo has two go source files. Compile them.

`gen.go`

Creates random number generator. `./gen 10 1000` generates 10 random numbers from 0 to 1000.

`cli-bloom.go`

Creats a program that reads input from stdin and outputs any duplicates. `./cli-bloom 1000`, will create a bloom filter with an `m` of 1000.
