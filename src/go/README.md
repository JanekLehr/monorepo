# Go source

## Environment Setup
#### Set helpful aliases
```
source env.sh
```

#### Generate go BUILD files
```
./pants buildgen.go --materialize --remote
```

#### IDE support
Run the following to open up an IDE with the go environment configured. Rerun it if new BUILD targets have been added, like a new `go_library`

```
# Open the repo in VSCode with GOROOT and GOPATH set to what pants controls
./pants go-env src/go/:: -- code .
```