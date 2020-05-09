# Introduction to Go Language

## Install Golang with Homebrew:

Fetch latest version of homebrew and formula.
```
brew update
```
Displays information about the given formulae.
```
brew info go
```
Install the given formulae.
```
brew install go
```
Clean up installation files
```
brew cleanup  
```
Check the installed Version
```
go version
```

## Setup the workspace:

For Bash users
```
vi ~/.bash_profile
```

Add these statements to your profile:
```
export GOPATH=$HOME/go-workspace # don't forget to change your path correctly!
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin:$GOROOT/bin
```

Create your workspace and primary folders
```
mkdir -p $GOPATH $GOPATH/src $GOPATH/pkg $GOPATH/bin
```

Check the GO Env variables.
```
go env
```

Override GOBIN setting if and only if you don't want default installation path
```
go env -w GOBIN=/somewhere/else/bin
go env -u GOBIN
```

## Basic Commands:

Basic run command
```
go run hello.go
```

Compiles the go file and generate executable on the same folder
```
go build hello.go
```

Installation command to compile and build the executable file on the default GOBIN path.
Choose one of these command options.
```
go install hello.go
go install example.com/user/hello
go install .
go install
```

Get the module.  Similar to git clone
```
go get -u github.com/gorilla/mux
```

Clean up the go cache files
```
go clean --modcache
```

AutoFormat your code
```
gofmt -w yourcode.go
```
OR
```
go fmt path/to/your/package
```

Documentaiton commands
```
godoc fmt                # documentation for package fmt
godoc fmt Printf         # documentation for fmt.Printf
godoc -src fmt           # fmt package interface in Go source form
```

## Reference:
https://golang.org/
https://blog.golang.org/godoc
https://tour.golang.org/list

