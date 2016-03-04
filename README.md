# A simple Golang templates example

During the process of learning Golang templates, certain aspects were
confusing to me. The goal of this little repo is to document what I
eventually wound up doing. Hopefully with feedback this repo could
serve as an example of at least one way to use templates effectively.

By **no** means is this intended to be a proper (*or even correct*)
howto on Golang templates, rather it's just what I've learned so far.
Here's what I was trying to accomplish:

1. Have a directory of templates (`header.html`, `foobar.html`, etc).
1. Have a directory of static files (css, images, etc).
1. Use some templates as full pages (`about.hmtl`, `hello.html`).
1. Use some templates as partials (`header.hmtl`, `footer.html`).
1. Serve static content in a manner similar to
   [http.FileServer](https://golang.org/pkg/net/http/#example_FileServer).
1. Exclude templates from the static files being served.
1. Support custom template functions.
1. Compile everything into a single static binary (including templates
   and static files).

## Installation

```
go get github.com/jteeuwen/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
go get github.com/julienschmidt/httprouter
go get -d github.com/jmcfarlane/golang-templates-example
```

## Run

```
cd $GOPATH/src/github.com/jmcfarlane/golang-templates-example
go generate
go build && ./golang-templates-example
curl http://localhost:8080/hello/jack
xdg-open http://localhost:8080/static/golang.png
```

**NOTE:**

To my knowledge there is no way to run `go generate` automatically, so
by not having the resulting `bindata_assetfs.go` file in source
code... the resulting project is not "*go gettable*".

## Notes

I have not yet sorted out how to exclude the templates from being
served over http (hence the strike through).
