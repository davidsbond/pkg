# pkg [![PkgGoDev](https://pkg.go.dev/badge/pkg.dsb.dev)](https://pkg.go.dev/pkg.dsb.dev) [![Go Report Card](https://goreportcard.com/badge/pkg.dsb.dev)](https://goreportcard.com/report/pkg.dsb.dev) ![CI](https://github.com/davidsbond/pkg/workflows/CI/badge.svg)

A collection of go packages I always end up reusing, so now they live in their own module.

<!-- ToC start -->
# Table of Contents

   1. [Getting started](#getting-started)
   1. [Tools](#tools)
<!-- ToC end -->

## Getting started

Add the module as a dependency using `go get`:

```shell script
go get pkg.dsb.dev
```

Pick and choose the packages you want to use, there's all sorts of stuff in here!

## Tools

This project uses a few go-based tools for linting, formatting and code generation. You can install these using
`make install-tools`

* [golangci-lint](https://github.com/golangci/golangci-lint) - Used to lint go source code.
* [gofumpt](https://github.com/mvdan/gofumpt) - Used to format go code and package imports
* [markdown-toc](https://github.com/sebdah/markdown-toc) - Used to generate the table of contents in this README.

