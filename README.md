# gRPCblockchain
Created a grpc blockchain server and client that use protobuf for communication. Used grpcurl for testing the server. Implemented client and server side logging.
    
## Running

1. Start the server's main.go file using `go run main.go`.
2. In other window add blocks or get the Blockchain with client's main.go file using `go run main.go` with tag `-- list` or `-- add`.
3. Note there si no database all the blockchain data is stored in the server, as soon as the server file completes execution the data will be lost.
    
## References

1. https://stackoverflow.com/questions/26694271/go-install-doesnt-create-any-bin-file
2. https://www.digitalocean.com/community/tutorials/how-to-write-packages-in-go
3. https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc
4. https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
5. https://stackoverflow.com/questions/47347191/go-error-unexpected-nul-in-input?rq=1
6. https://www.quora.com/What-is-the-difference-between-build-and-install-in-Go
7. https://www.youtube.com/watch?v=gju-bml4kdw
8. https://github.com/tensor-programming/golang-blockchain/tree/part_10
