# example-go-api

### Common go api project structure
Read about [project structure layout](https://github.com/golang-standards/project-layout)
- `api`: specs: parameters for endpoints, response types for endpoints (yaml spec file)
- `cmd/api`: contains `main.go` file
- `internal` contains most of the code for the api

### Common go commands for packages:
- `go install` is used to install a binary, not a package

- `go get` to download a package

- `go mod tidy` scans your project and updates the go.mod file and downloads all dependencies mentioned in the go.mod file 
  - similar to `npm i` for node

- `go mod download` only downloads the dependencies from go.mod, without modifications of the go.mod file