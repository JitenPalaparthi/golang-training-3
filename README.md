
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


### Builtin functions (18)
append,cap,clear,copy,complx, len, make,print, println

### Keywords (25)


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

### Variable declaration

- Code/Text Segment --> binary is loaded into this segment, consts are loaded into code segment
- Data Segment      --> all global(package) variables(static) are loaded into data segment

- Stack Memory      --> stack memory, generally all local variables, complier writes the code based on blocks
- Heap Memory       --> variables on heap, by GC