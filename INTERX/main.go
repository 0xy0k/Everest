package main

import (
	"flag"
	"io/ioutil"
	"os"
	// Static files

	"github.com/TsukiCore/tsuki/INTERX/gateway"
	_ "github.com/TsukiCore/tsuki/INTERX/statik"
	tsukiapp "github.com/TsukiCore/tsuki/app"
	"google.golang.org/grpc/grpclog"
)

var serverGRPCAddress = flag.String(
	"server-gRPC-address",
	"dns:///0.0.0.0:9090",
	"The address to the gRPC server, in the gRPC standard naming format. ",
)

var serverRPCAddress = flag.String(
	"server-RPC-address",
	"http://0.0.0.0:26657",
	"The address to the RPC server, in the RPC standard naming format. ",
)

func main() {
	flag.Parse()

	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	tsukiapp.SetConfig()

	grpcAddr := os.Getenv("GRPC")
	if len(grpcAddr) == 0 {
		grpcAddr = *serverGRPCAddress
	}

	rpcAddr := os.Getenv("RPC")
	if len(rpcAddr) == 0 {
		rpcAddr = *serverRPCAddress
	}

	err := gateway.Run(grpcAddr, rpcAddr, log)
	log.Fatalln(err)
}
