# OpenPSD PSD2 Server

Reference implementation for the Berlin Group PSD2 specification.

Note: Work in progress!

## requirements

* Go 1.9.x

Working GO environment is expected as described [here](https://golang.org/doc/code.html#GOPATH) 
## install

`make install`

Will install `psd2-server` to you default `go` binary path.

## run

`make run`

Will start a http server at localhost:8000

## use

`GET http://localhost:8000/accounts`

## test

`make test` will run all unit tests