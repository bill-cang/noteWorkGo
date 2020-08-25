package main

import (
	"flag"
	//"github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"

	"noteWork/example/xrpcex"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	/*	flag.Parse()
		s := server.Server{}
		//rcvr := new(xrpcex.Arith)
		s.RegisterName("Arith", new(xrpcex.Arith), "")
		go s.Serve("tcp", *addr)
		select {}*/

	s := server.NewServer()
	s.RegisterName("Arith", new(xrpcex.Arith), "")
	s.Serve("tcp", ":8972")
}
