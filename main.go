package main

import (
	"context"
	proto "gRPCblockchain/proto"
	blockchain "gRPCblockchain/server/blockchain"
	"net"
	"os"

	"google.golang.org/grpc/reflection"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()

	proto.RegisterBlockchainServer(gs, &BlockchainServer{
		l:          log,
		Blockchain: blockchain.NewBlockchain(),
	})

	reflection.Register(gs)

	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(listener)
}

// implementation of blockchain server present in the .pb.go file.
type BlockchainServer struct {
	l          hclog.Logger
	Blockchain *blockchain.Blockchain // pointing to the struct.
}

/* func NewBlockchainServer(l hclog.Logger, bc *blockchain.Blockchain) *BlockchainServer {
	return &BlockchainServer{l, bc}
} */

func (bs *BlockchainServer) AddBlock(ctx context.Context, abr *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method AddBlock not implemented")
	bs.l.Info("Handle AddBlock", "data", abr.GetData())
	b := bs.Blockchain.AddBlock(abr.GetData())

	return &proto.AddBlockResponse{
		Hash: b.Hash,
	}, nil
}

func (bs *BlockchainServer) GetBlockchain(ctx context.Context, gbr *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	bs.l.Info("Handle GetBloackchain")

	resp := new(proto.GetBlockchainResponse)
	for _, b := range bs.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Data:          b.Data,
			Hash:          b.Hash,
		})
	}

	return resp, nil
}
