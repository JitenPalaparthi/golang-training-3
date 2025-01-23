# All Go Commands with Descriptions

---

### 1. `go build`
- **Description**: Compiles and links the packages to produce a binary.
- **Options**:
  - `-o <file>`: Specifies the output binary name.
  - `-a`: Forces rebuilding of all packages.
  - `-v`: Prints the names of packages compiled.
  - `-x`: Prints the build commands.
- **Example**:
  ```bash
  go build -o app main.go
  go build -v ./...

2. go run
	•	Description: Compiles and runs the specified Go program.
	•	Options:
	•	<files>: Specify the source file(s) to run.
	•	-exec <cmd>: Specifies a custom execution wrapper.
	•	Example:

```bash
go run main.go
go run -exec "strace" .
```