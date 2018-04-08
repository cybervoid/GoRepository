# Setup
1. Atom


## Setting up Atom: Atom
### Windows
1. Install Atom
  - change atom to bash command C:\Program Files\Git\bin
2. Install Go
3. Set environment variables
4. run commands in (Atom terminal):
```
go get -u golang.org/x/tools/cmd/...
go get -u github.com/golang/lint
go get -u github.com/nsf/gocode
```
5. Ensure packages.txt file is in working directory and run:
```
apm install --packages-file packages.txt
or:
C:\Users\[username]\AppData\Local\atom\bin\apm install --packages-file packages.txt
or;

apm install --packages-file ./src/github.com/GoRepository/Setup/packages.txt
```



## Go Bash commands:
1. install a package from github: Example  `go get github.com/ethereum/go-ethereum/common`

## Environment Variables
1. get all go environment variables in bash: `go env`  
2. Set gopath to GoWorkspace
  - Powershell: https://stackoverflow.com/questions/12587253/how-to-modify-gopath-in-powershell  `example powershell command: [Environment]::SetEnvironmentVariable("GOPATH","C:\Users\Go\GoWorkspace")`
  - Unix terminal `https://github.com/golang/go/wiki/SettingGOPATH`
  -- `export gopath=E:\\my\\go\\workspace`
## Resources
1. [GoDocs.org](https://godoc.org/)

## Setup Atom.
Notes: https://atom.io/packages/go-plus
`go get -u github.com/joefitzgerald/go-plus`

Go certificate issue:
https://github.com/trueos/trueos-core/issues/181

### Mac
