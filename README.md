# SERVE [![Build Status](https://travis-ci.org/ttgint/serve.svg?branch=master)](https://travis-ci.org/ttgint/serve)

Simple server for single page applications.

## Usage

```
$ serve -h
Usage of serve:
  -d string
        The root directory to host (default ".")
  -l string
        Specify a URI endpoint on which to listen (default ":3000")
  -i string
        The index file (default "index.html")
```

## Building from source

`$ go build -ldflags="-s -w"`
