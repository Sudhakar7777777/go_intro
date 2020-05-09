# Introduction to Go Language

## Install Golang with Homebrew:

```
brew update                 # Fetch latest version of homebrew and formula.
brew info go                # Displays information about the given formulae.
brew install go             # Install the given formulae.
brew cleanup  
```

```
go version
```

## Setup the workspace:

```
vi .bashrc
export GOPATH=$HOME/go-workspace # don't forget to change your path correctly!
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin:$GOROOT/bin
mkdir -p $GOPATH $GOPATH/src $GOPATH/pkg $GOPATH/bin
```

```
go env
go env -w GOBIN=/somewhere/else/bin
go env -u GOBIN
```

## Basic Commands:

```
go run hello.go
```

```
go install hello.go
go install example.com/user/hello
go install .
go install
```

```
go get -u github.com/gorilla/mux
```

```
go clean --modcache
```

```
gofmt -w yourcode.go
```
OR
```
go fmt path/to/your/package
```

```
godoc fmt                # documentation for package fmt
godoc fmt Printf         # documentation for fmt.Printf
godoc -src fmt           # fmt package interface in Go source form
```

## Reference:
https://golang.org/
https://blog.golang.org/godoc
https://tour.golang.org/list

