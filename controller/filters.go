package controller

import (
	"context"
	"envoy-test-filter/filters"
	ext_authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"github.com/gogo/googleapis/google/rpc"
	"log"
)

func ExecuteFilters(ctx context.Context, req *ext_authz.CheckRequest)  (*ext_authz.CheckResponse, error) {
	swagg, err := readApis()
	// fmt.Println(swagg)

	if swagg != nil {

	}

	resp , err, tokenData := filters.ValidateToken(ctx, req)
	//resp , err, tokenData2 := filters.ValidateThrottling(ctx, req)

	log.Println(tokenData)
	//Return if the authentication failed
	if resp.Status.Code != int32(rpc.OK) {
		return resp, nil
	}
	//Continue to next filter

	// Publish analytics to analytics server
	resp , err = filters.PublishAnalyticsFromJwt(ctx, req, tokenData)
	//resp , err = filters.PublishAnalyticsFromJwt(ctx, req, tokenData2)

	return resp, err

}
