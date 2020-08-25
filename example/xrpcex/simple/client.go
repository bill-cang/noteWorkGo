package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"noteWork/example/xrpcex"
)

/*var (
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)*/

func main() {
	Peer2Peer()
}
func Peer2Peer() {
	flag.Parse()
	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &xrpcex.Args{
		A: 10,
		B: 20,
	}

	reply := &xrpcex.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
