// a command line application for using our server.package main

package main

import (
	"context"
	"flag"
	proto "gRPCmyTry/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	ListFlag := flag.Bool("list", false, "get the blockchain")
	flag.Parse()

	// WithInsecure() method to disable the client security.
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to dial", "error", err)
	}

	client := proto.NewBlockchainClient(conn)

	if *addFlag {
		b, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
			Data: time.Now().String(),
		})

		if err != nil {
			log.Fatalf("error while adding block", "error", err)
		}

		log.Printf("the new generated block's hash is", b.Hash)
	}

	if *ListFlag {
		bc, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})

		if err != nil {
			log.Fatalf("unable to get blockchain: %v", err)
		}

		log.Println("blocks:")
		for _, b := range bc.Blocks {
			log.Printf("hash: %s, previous block hash: %s, data: %s", b.Hash, b.PrevBlockHash, b.Data)
		}
	}
}
