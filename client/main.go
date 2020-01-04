package main

import (
	"flag"
	"log"
	"time"
	"github.com/faith/grpcBlockchain/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Global client variable.
var client proto.BlockchainClient

func main() {
	// client can do two things, it can add a block or send request to list the blockchain.
	// to add block do "go run main.go --add"
	// to get blockchain do "go run main.go --list"
	addFlag := flag.Bool("add", false, "Add new block")
	listFlag := flag.Bool("list", false, "List all blocks")
	flag.Parse()

	// check connection.
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	client = proto.NewBlockchainClient(conn)

	if *addFlag {
		addBlock()
	}

	if *listFlag {
		getBlockchain()
	}
}

func addBlock() {
	// Except genesis block all other blocks have dates.
	block, addErr := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if addErr != nil {
		log.Fatalf("unable to add block: %v", addErr)
	}
	log.Printf("new block hash: %s\n", block.Hash)
}

func getBlockchain() {
	blockchain, getErr := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if getErr != nil {
		log.Fatalf("unable to get blockchain: %v", getErr)
	}

	log.Println("blocks:")
	for _, b := range blockchain.Blocks {
		log.Printf("hash %s, prev hash: %s, data: %s\n", b.Hash, b.PrevBlockHash, b.Data)
	}
}
