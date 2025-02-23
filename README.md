
1. Binary  
2. Library/packages 

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


### Builtin functions (18 of 18)
append,cap,clear,close,copy,complex,delete,imag,len,make,max,min,new,panic,print,println,real,recover

### Keywords (25)
break,case,chan,const,continue,default,defer,else,fallthrough,for,func,go,goto,if,import,interface,map,package,range,return,select,struct,switch,type, var (25 of 25)

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

### escape analysis

```bash
go build -gcflags="-m" main.go
```

### Golang Presentation

https://docs.google.com/presentation/d/1WVvsbvgHKBrNrKtnT4XWRfrsfkNlbw5u6L9O1DeVBn0/edit?usp=sharing

### Variable declaration

- Code/Text Segment --> binary is loaded into this segment, consts are loaded into code segment
- Data Segment      --> all global(package) variables(static) are loaded into data segment

- Stack Memory      --> stack memory, generally all local variables, complier writes the code based on blocks
- Heap Memory       --> variables on heap, by GC


### go get to fetch other packages

```bash
go get github.com/JitenPalaparthi/icici-shapes-package3
```


### To test a pacakge

```bash
go test demo/models
```
### To run the single test

```bash
go test -timeout 30s -run ^TestUserValidateSuccess$ demo/models
```

### To run all tests in a file

```bash
go test -timeout 30s -run ^(TestUserValidateSuccess|TestUserValidateNameFailure|TestUserValidateEmailFailure|TestUserToBytes|TestUserToString)$ demo/models
```

### Cover Profile

```bash
go test -coverprofile=coverage.out demo/models
```
### To see the coverage using html format


```bash
go tool cover -html coverage.out
```