# unid.go

**Highly unique, timestamp based identifiers generator for Golang.**

[![wercker status](https://app.wercker.com/status/9ed51fe7081392603a3835b9f732ce51/s/ "wercker status")](https://app.wercker.com/project/bykey/9ed51fe7081392603a3835b9f732ce51)

Unid is a small and efficient package that generates timestamp based
identifiers with extra uniqueness ensure. The algorithm is pretty simple:

1. Generate a unix nano timestamp,
2. Multiply this timestamp by 100,
3. Add sequential atomically generated number from range 0-99.

This simple way we have time based and highly unique identfier in our hands.

## Getting started

Add `unid.go` to your imports...

    import "github.com/kkvlk/unid.go"

... and run go get from the project directory:

    $ go get

## Usage

You can find full Go documentation [here](http://godoc.org/github.com/kkvlk/unid.go)

## Hacking

If you wanna hack on the repo for some reason, first clone it, then run tests:

    $ go test

You can also run benchmarks with:

    $ go test -bench .

That's all.

## Contributing

1. Work on your changes in a feature branch.
2. Make sure that tests are passing.
3. Send a pull request.
4. Wait for feedback.

## Copyrights

Copyright (c) by Kris Kovalik <<hi@kkvlk.me>>

Released under BSD License. Check [LICENSE.txt](LICENSE.txt) for details.
