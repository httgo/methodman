# methodman

[![Build Status](https://travis-ci.org/httgo/methodman.svg?branch=master)](https://travis-ci.org/httgo/methodman)
[![GoDoc](https://godoc.org/gopkg.in/httgo/methodman.v0?status.svg)](http://godoc.org/gopkg.in/httgo/methodman.v0)

Method override handler

## Install

    go get gopkg.in/httgo/methodman.v0

## Example

In your form

    <input type="hidden" name="_method" value="PUT">
    <input type="hidden" name="_method" value="PATCH">
    <input type="hidden" name="_method" value="DELETE">

Wrap your handler

    mux.Put(...)

    h := methodman.MethodOverride(mux)

    err := http.ListenAndServe(":3000", h)
    if err != nil {
      log.Fatalf("fatal: listen: %s", err)
    }

If you use [Alice](https://github.com/justinas/alice).

    mux.Put(...)

    chain := alice.New(methodman.MethodOverride)
    h := chain.Then(mux)

## License

MIT

