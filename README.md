# GoGo

Project Level Go version manager, it's works by complimenting the existing
version manager [`golang.org/dl`](https://golang.org/doc/install#extra_versions) 🙂

## Requirement

[Google Go](https://golang.org/)  
[Git](https://git-scm.com/)

## Installation

Run the following (outside of go modules)

```sh
$ go get github.com/cjtoolkit/gogo
```

## Using GoGo

Create `.gogo` in the root of the project, with the example (it can be any
version you desire, as long as it in
[`golang.org/dl`](https://pkg.go.dev/golang.org/dl)).
```
go1.14.2
```

You should be able to run a go application with the specific version stated
in `.gogo` by running `$ gogo run main.go`.

If the specified version of go is it not installed it's will install it
automatically.