# CSViz

Read and inspect CSVs directly from you terminal!

## Installation

### Using `go install`
If you have a [Go](https://go.dev/) environment ready to go, it's as easy as:
```
go install https://github.com/davidpalves/csviz
```

### Building from source
Since this tool is written in [Go](https://go.dev/) you need to install the Go language/compiler/etc. Full details of installation and set up can be found on the [Go language website](https://golang.org/doc/install).

Clone this repository:
```
git clone https://github.com/davidpalves/csviz.git
```

 Once installed Go and cloned this repo, you have two options:

#### Compiling
`CSViz` has external dependencies, and so they need to be pulled in first:

```
go get && go build
```

This will create a `csviz` binary for you. If you want to install it in the $GOPATH/bin folder you can run:

```
go install
```

## Usage
You can see the available flags of `csviz` using the following command:
```
./csviz
```
or
```
./csviz --help
```
The following output will be given:

```
Reads CSV directly from terminal

Usage:
  csviz [flags]

Flags:
  -f, --filepath string   Path to the CSV file to be read
  -h, --help              help for csviz
```

## Demo

```
csviz -f examples/example.csv
```

Output: