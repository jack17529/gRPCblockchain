# gRPCblockchain
## How to know the value of Go Environment Variables?

Use `go env` command.

## Using custom made packages in Go

1. Create the package in it's own folder.
2. Install it using go install command.
3. Import it in the file using the address relative to $GOPATH

For Example

For a package file "blockchain.go" present in `/src/github.com/faith/00-grpc/server/blockchain`
Can be used inside a file "main.go" present in `/src/github.com/faith/00-grpc/server`
Using `"github.com/faith/00-grpc/server/blockchain"` in imports.

## Directory structure for Ubuntu

- ~
  - Go                              <- Where $GOPATH is set.
    - src
      - github.com
        - faith
          - grpcBlockchain                 <- Project name
            - server
              - blockchain          <- Each package needs to have it's own directory.
                - blockchain.go
              - main.go
            - client
              - main.go
            - proto
              - blockchain.proto
              - blockchain.pb.go
    - pkg
    - bin
    
## Running

1. Start the server's main.go file using `go run main.go`.
2. In other window add blocks or get the Blockchain using client's main.go file using `go run main.go` with tag `-- list` or `-- add`.
3. Note there si no database all the blockchain data is stored in the server, as soon as the server file completes execution the data will be lost.
    
## References

1. https://stackoverflow.com/questions/26694271/go-install-doesnt-create-any-bin-file
2. https://www.digitalocean.com/community/tutorials/how-to-write-packages-in-go
3. https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc
4. https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
5. https://stackoverflow.com/questions/47347191/go-error-unexpected-nul-in-input?rq=1
6. https://www.quora.com/What-is-the-difference-between-build-and-install-in-Go
7. https://www.youtube.com/watch?v=gju-bml4kdw
