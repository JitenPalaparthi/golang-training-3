
1. Binary 
2. Library

### Creating a module

```bash
go mod init demo
```

- run 
```bash
go run main.go
```
-build
```bash
go build main.go
```

- build output

```bash
go build -o hello main.go
```


### Builtin functions


### Keywords

### go env

1. GOROOT --> all the standard libraries and go itself is installed here
2. GOPATH --> all third party packages are managed here
3. GOBIN  --> all binaries are installed here.If GOBIN is not available it pics from GOPATH
4. GOOS --> Operating system
5. GOARCH --> Processor architecture 

### go targets

```bash
go tool dist list
```


### Cross compilation

```bash
GOOS=windows GOARCH=amd64 go build -o build/hello_windows_amd64.exe main.go 
```

### release build 

```bash
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o build/hello_small main.go
```


### Golang Presentation

https://docs.google.com/presentation/d/1WVvsbvgHKBrNrKtnT4XWRfrsfkNlbw5u6L9O1DeVBn0/edit?usp=sharing