package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hariskhan14/grpc-gateway-response-modifier/domain/service"
	api "github.com/hariskhan14/grpc-gateway-response-modifier/proto"
	"github.com/hariskhan14/grpc-gateway-response-modifier/utils"
	"log"
	"net/http"
)

const (
	GRPCGatewayPort = ":8090"
)

func main() {
	userService := service.NewUserService()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(utils.HeaderMatcher),
		runtime.WithForwardResponseOption(utils.ResponseStatusCodeModifier),
	)
	if err := api.RegisterUserServiceHandlerServer(ctx, mux, userService); err != nil {
		log.Fatalf("Failed to register gRPC service server: %v", err)
	}

	gwServer := &http.Server{Addr: GRPCGatewayPort, Handler: mux}
	log.Printf("Serving gRPC-Gateway on %s", GRPCGatewayPort)
	log.Fatalln(gwServer.ListenAndServe())
}
