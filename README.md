# Papa Dew AI
This is a super intelligent AI and its a Papa Dew. 

# Usage Guide
## Pre-requisite 
### Go
Please install Go. Ensure its installed correctly. 

#### Checking Go version example
```bash
~/Project/papa-dew-ai [ go version 
go version go1.22.0 darwin/arm64
```

## How to use
This is a simple CLI tools which will accept input args when executing the program
### CLI structure
```bash
<PROGRAM> <ARGS1> <ARGS2>
```
Legends: 
* PROGRAM = The program itself
* ARGS1 = URL of the site to be scraped
* ARGS2 = Query Keywords to filter out contents

### Example Usage directly using Go command
```bash
go run main.go https://example.com example
```

### Example Usage using compile binary of the program
#### Building the application
Ensure that `make` cli is installed
```bash
make build
```
`papa-dew-ai` will be the binary file name under the `bin` directory
#### Executing the program
```bash
./bin/papa-dew-ai https://example.com example
```
