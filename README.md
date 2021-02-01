# A simple Golang templates example

[![Go Report Card](https://goreportcard.com/badge/jmcfarlane/golang-templates-example)](https://goreportcard.com/report/jmcfarlane/golang-templates-example)
[![Build Status](https://img.shields.io/travis/jmcfarlane/golang-templates-example/main.svg)](https://github.com/jmcfarlane/golang-templates-example/tree/main)
[![codecov](https://codecov.io/gh/jmcfarlane/golang-templates-example/branch/main/graph/badge.svg)](https://codecov.io/gh/jmcfarlane/golang-templates-example)

During the process of learning Golang templates, certain aspects were
confusing to me. The goal of this little repo is to document what I
eventually wound up doing. Hopefully with feedback this repo could
serve as an example of at least one way to use templates effectively.

By **no** means is this intended to be a proper (*or even correct*)
howto on Golang templates, rather it's just what I've learned so far.
Here's what I was trying to accomplish:

1. Have a directory of templates (`header.html`, `foobar.html`, etc).
1. Have a directory of static files (css, images, etc).
1. Use some templates as full pages (`about.html`, `hello.html`).
1. Use some templates as partials (`header.html`, `footer.html`).
1. Serve static content in a manner similar to
   [http.FileServer](https://golang.org/pkg/net/http/#example_FileServer).
1. Exclude templates from the static files being served.
1. Support custom template functions.
1. Compile everything into a single static binary (including templates
   and static files).

## Installation

```
go get github.com/GeertJohan/go.rice/rice
go get -d github.com/jmcfarlane/golang-templates-example
```

## Run

```
cd $GOPATH/src/github.com/jmcfarlane/golang-templates-example
go get -t ./...
go generate
go test -v
go build
./golang-templates-example
curl http://localhost:8080
```

## Optionally use make

If you have `make` installed, you can try this repo by doing:

```
make
./golang-templates-example
```
