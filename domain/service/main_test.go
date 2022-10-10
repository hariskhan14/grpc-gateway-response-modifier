package service_test

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hariskhan14/grpc-gateway-response-modifier/domain/service"
	api "github.com/hariskhan14/grpc-gateway-response-modifier/proto"
	"github.com/hariskhan14/grpc-gateway-response-modifier/utils"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

type TestContext struct {
	userService *service.UserServiceImpl
	grpcServer  *httptest.Server
}

var testCtx TestContext

func TestMain(m *testing.M) {
	userService := service.NewUserService()

	//GRPC-Gateway
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(utils.HeaderMatcher),
		runtime.WithForwardResponseOption(utils.ResponseStatusCodeModifier),
	)
	if err := api.RegisterUserServiceHandlerServer(ctx, mux, userService); err != nil {
		log.Fatalf("Failed to register gRPC gateway service endpoint: %v", err)
	}

	server := httptest.NewServer(mux)
	defer server.Close()

	testCtx = TestContext{
		userService: userService,
		grpcServer:  server,
	}

	os.Exit(m.Run())
}
