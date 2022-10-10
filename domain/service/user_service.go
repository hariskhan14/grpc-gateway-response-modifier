package service

import (
	"context"
	api "github.com/hariskhan14/grpc-gateway-response-modifier/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"strings"
)

type UserService interface {
	GetUserDetails(context.Context, *api.GetUserDetailsRequest) (*api.GetUserDetailsResponse, error)
}

type UserServiceImpl struct {
	api.UnimplementedUserServiceServer
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (srv *UserServiceImpl) GetUserDetails(ctx context.Context, in *api.GetUserDetailsRequest) (*api.GetUserDetailsResponse, error) {
	version := ""
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if val, ok := md["version"]; ok {
			version = strings.Join(val, ",")
		}
	}

	if version != "1.0" {
		_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", strconv.Itoa(http.StatusNotAcceptable)))
		return nil, status.Error(codes.InvalidArgument, "Invalid version")
	}

	return &api.GetUserDetailsResponse{
		FirstName: "grpc",
		Last_Name: "gateway",
	}, nil
}
