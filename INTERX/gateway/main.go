package gateway

import (
	"context"
	"fmt"
	"mime"
	"net/http"

	"github.com/TsukiCore/tsuki/INTERX/config"
	"github.com/TsukiCore/tsuki/INTERX/database"
	"github.com/TsukiCore/tsuki/INTERX/functions"
	cosmosAuth "github.com/TsukiCore/tsuki/INTERX/proto-gen/cosmos/auth"
	cosmosBank "github.com/TsukiCore/tsuki/INTERX/proto-gen/cosmos/bank"
	tsukiGov "github.com/TsukiCore/tsuki/INTERX/proto-gen/tsuki/gov"
	tsukiSlashing "github.com/TsukiCore/tsuki/INTERX/proto-gen/tsuki/slashing"
	tsukiStaking "github.com/TsukiCore/tsuki/INTERX/proto-gen/tsuki/staking"
	"github.com/TsukiCore/tsuki/INTERX/tasks"
	functionmeta "github.com/TsukiCore/tsuki/function_meta"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	grpclog "google.golang.org/grpc/grpclog"
)

// getOpenAPIHandler serves an OpenAPI UI.
func getOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")

	statikFS, err := fs.New()
	if err != nil {
		// Panic since this is a permanent error.
		panic("creating OpenAPI filesystem: " + err.Error())
	}

	return http.FileServer(statikFS)
}

// GetGrpcServeMux is a function to get ServerMux for GRPC server.
func GetGrpcServeMux(grpcAddr string) (*runtime.ServeMux, error) {
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	// WITH_TRANSPORT_CREDENTIALS: Empty parameters mean set transport security.
	security := grpc.WithInsecure()

	// With transport credentials
	// if strings.ToLower(os.Getenv("WITH_TRANSPORT_CREDENTIALS")) == "true" {
	// 	security = grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, ""))
	// }

	conn, err := grpc.DialContext(
		context.Background(),
		grpcAddr,
		security,
		grpc.WithBlock(),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to dial server: %w", err)
	}

	gwCosmosmux := runtime.NewServeMux()
	err = cosmosBank.RegisterQueryHandler(context.Background(), gwCosmosmux, conn)
	if err != nil {
		return nil, fmt.Errorf("failed to register gateway: %w", err)
	}

	err = cosmosAuth.RegisterQueryHandler(context.Background(), gwCosmosmux, conn)
	if err != nil {
		return nil, fmt.Errorf("failed to register gateway: %w", err)
	}

	err = tsukiGov.RegisterQueryHandler(context.Background(), gwCosmosmux, conn)
	if err != nil {
		return nil, fmt.Errorf("failed to register gateway: %w", err)
	}

	err = tsukiStaking.RegisterQueryHandler(context.Background(), gwCosmosmux, conn)
	if err != nil {
		return nil, fmt.Errorf("failed to register gateway: %w", err)
	}

	err = tsukiSlashing.RegisterQueryHandler(context.Background(), gwCosmosmux, conn)
	if err != nil {
		return nil, fmt.Errorf("failed to register gateway: %w", err)
	}

	return gwCosmosmux, nil
}

// Run runs the gRPC-Gateway, dialling the provided address.
func Run(configFilePath string, log grpclog.LoggerV2) error {
	config.LoadConfig(configFilePath)
	functions.RegisterInterxFunctions()
	functionmeta.RegisterStdMsgs()

	database.LoadBlockDbDriver()
	database.LoadFaucetDbDriver()
	database.LoadReferenceDbDriver()

	grpcAddr := config.Config.GRPC
	rpcAddr := config.Config.RPC
	port := config.Config.PORT

	tasks.RunTasks(rpcAddr)

	gwCosmosmux, err := GetGrpcServeMux(grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	oaHander := getOpenAPIHandler()

	router := mux.NewRouter()
	RegisterRequest(router, gwCosmosmux, rpcAddr)

	router.PathPrefix("/").Handler(oaHander)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})

	gatewayAddr := "0.0.0.0:" + port
	gwServer := &http.Server{
		Addr:    gatewayAddr,
		Handler: c.Handler(router),
	}

	// SERVE_HTTP: Empty parameters mean use the TLS Config specified with the server.
	// if strings.ToLower(os.Getenv("SERVE_HTTP")) == "false" {
	// 	gwServer.TLSConfig = &tls.Config{
	// 		Certificates: []tls.Certificate{insecure.Cert},
	// 	}

	// 	log.Info("Serving gRPC-Gateway and OpenAPI Documentation on https://", gatewayAddr)
	// 	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServeTLS("", ""))
	// }

	log.Info("Serving gRPC-Gateway and OpenAPI Documentation on http://", gatewayAddr)
	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())
}
