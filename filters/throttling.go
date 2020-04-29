package filters

import (
	"context"
	"fmt"
	ext_authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func ValidateThrottling(ctx context.Context, req *ext_authz.CheckRequest) (*ext_authz.CheckResponse, error, TokenData) {

	caCert,_ := ReadFile("./artifacts/server.pem")

	var keys []string
	auth := false
	var tokenData TokenData
	for k := range req.Attributes.Request.Http.Headers {
		if k == "authorization" {
			//h = true
			//header := req.Attributes.Request.Http.Headers["authorization"]
			auth, tokenData, _ = HandleJWT(false, caCert,req.Attributes.Request.Http.Headers )
			fmt.Println("JWT header detected" + k)
		}
		keys = append(keys, k)
	}

	resp := &ext_authz.CheckResponse{}
	if auth {
		resp = &ext_authz.CheckResponse{
			Status: &status.Status{Code: int32(rpc.OK)},
			HttpResponse: &ext_authz.CheckResponse_OkResponse{
				OkResponse: &ext_authz.OkHttpResponse{

				},
			},
		}

	} else {
		resp = &ext_authz.CheckResponse{
			Status: &status.Status{Code: int32(rpc.UNAUTHENTICATED)},
			HttpResponse: &ext_authz.CheckResponse_DeniedResponse{
				DeniedResponse: &ext_authz.DeniedHttpResponse{
					Status:  &envoy_type.HttpStatus{
						Code: envoy_type.StatusCode_Unauthorized,
					},
					Body: "Error occurred while authenticating.",

				},
			},
		}
	}

	return resp, nil, tokenData
}
