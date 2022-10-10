package utils

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"
	"strings"
)

const GRPCGatewayHTTPHeader = "x-http-code"

func HeaderMatcher(key string) (string, bool) {
	hdr := strings.ToLower(key)
	switch hdr {
	case "version":
		return hdr, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func ResponseStatusCodeModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	// set http status code from grpc if exists
	if hdr := md.HeaderMD.Get(GRPCGatewayHTTPHeader); len(hdr) > 0 {
		code, err := strconv.Atoi(hdr[0])
		if err != nil {
			return err
		}

		w.WriteHeader(code)
		delete(md.HeaderMD, "x-http-code")
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
	}

	return nil
}
