package filters

import (
	"context"
	"crypto/tls"
	analytics "envoy-test-filter/pb"
	ext_authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"github.com/gogo/googleapis/google/rpc"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"strings"
)

var data = &analytics.AnalyticsStreamMessage{
	MessageStreamName:      "",
	MetaClientType:         "",
	ApplicationConsumerKey: "",
	ApplicationName:        "",
	ApplicationId:          "",
	ApplicationOwner:       "",
	ApiContext:             "",
	ApiName:                "",
	ApiVersion:             "",
	ApiResourcePath:        "",
	ApiResourceTemplate:    "",
	ApiMethod:              "",
	ApiCreator:             "",
	ApiCreatorTenantDomain: "",
	ApiTier:                "",
	ApiHostname:            "",
	Username:               "",
	UserTenantDomain:       "",
	UserIp:                 "",
	UserAgent:              "",
	RequestTimestamp:       0,
	ThrottledOut:           false,
	ResponseTime:           0,
	ServiceTime:            0,
	BackendTime:            0,
	ResponseCacheHit:       false,
	ResponseSize:           0,
	Protocol:               "",
	ResponseCode:           0,
	Destination:            "",
	SecurityLatency:        0,
	ThrottlingLatency:      0,
	RequestMedLat:          0,
	ResponseMedLat:         0,
	BackendLatency:         0,
	OtherLatency:           0,
	GatewayType:            "",
	Label:                  "",
	Subscriber:             "",
	ThrottledOutReason:     "",
	ThrottledOutTimestamp:  0,
	Hostname:               "",
	ErrorCode:              "",
	ErrorMessage:           "",
}
var data2 = &analytics.AnalyticsStreamMessage{
	MessageStreamName:      "",
	MetaClientType:         "",
	ApplicationConsumerKey: "",
	ApplicationName:        "",
	ApplicationId:          "",
	ApplicationOwner:       "",
	ApiContext:             "",
	ApiName:                "",
	ApiVersion:             "",
	ApiResourcePath:        "",
	ApiResourceTemplate:    "",
	ApiMethod:              "",
	ApiCreator:             "",
	ApiCreatorTenantDomain: "",
	ApiTier:                "",
	ApiHostname:            "",
	Username:               "",
	UserTenantDomain:       "",
	UserIp:                 "",
	UserAgent:              "",
	RequestTimestamp:       0,
	ThrottledOut:           false,
	ResponseTime:           0,
	ServiceTime:            0,
	BackendTime:            0,
	ResponseCacheHit:       false,
	ResponseSize:           0,
	Protocol:               "",
	ResponseCode:           0,
	Destination:            "",
	SecurityLatency:        0,
	ThrottlingLatency:      0,
	RequestMedLat:          0,
	ResponseMedLat:         0,
	BackendLatency:         0,
	OtherLatency:           0,
	GatewayType:            "",
	Label:                  "",
	Subscriber:             "",
	ThrottledOutReason:     "",
	ThrottledOutTimestamp:  0,
	Hostname:               "",
	ErrorCode:              "",
	ErrorMessage:           "",
}

const (
	address     = "localhost:9806"
	gatewayType = "MICRO"
)

func PublishAnalyticsFromJwt(ctx context.Context, req *ext_authz.CheckRequest, tokenData TokenData) (*ext_authz.CheckResponse, error){

	data.MessageStreamName = "InComingRequestStream"
	data.MetaClientType = tokenData.meta_clientType
	data.ApplicationConsumerKey = tokenData.applicationConsumerKey
	data.ApplicationName = tokenData.applicationName
	data.ApplicationId = tokenData.applicationId
	data.ApplicationOwner = tokenData.applicationOwner
	data.ApiCreator = tokenData.apiCreator
	data.ApiCreatorTenantDomain = tokenData.apiCreatorTenantDomain
	data.ApiTier = tokenData.apiTier
	data.Username = tokenData.username
	data.UserTenantDomain = tokenData.userTenantDomain
	data.ThrottledOut = tokenData.throttledOut
	data.ServiceTime = tokenData.serviceTime
	data.ApiContext = tokenData.apiContext
	data.ApiName = tokenData.apiName
	data.ApiVersion = tokenData.apiVersion
	data.ApiResourcePath = req.Attributes.Request.Http.Path
	data.Protocol = req.Attributes.Request.Http.Protocol
	data.ApiMethod = req.Attributes.Request.Http.Method
	data.UserAgent = req.Attributes.Request.Http.Headers["user-agent"]
	data.Hostname = strings.Split(req.Attributes.Request.Http.Host, ":")[0]
	data.RequestTimestamp = 1000 * (req.Attributes.Request.Time.Seconds) //RequestTimestamp in milli sec
	data.GatewayType = gatewayType
	data.Label = gatewayType
	data.Subscriber = ""
	data.ThrottledOutReason = ""
	data.ThrottledOutTimestamp = 0
	data.ErrorCode = ""
	data.ErrorMessage = ""
	data.ResponseMedLat = 0
	data.RequestMedLat = 0
	data.OtherLatency = 0

	// hard coded values for analytics stream message
	data2.MessageStreamName = "InComingRequestStream"
	data2.MetaClientType = "PRODUCTION"
	data2.ApplicationConsumerKey = "snNa8q_njWUjk7oWQUWgTD5KfE4a"
	data2.ApplicationName = "Application01"
	data2.ApplicationId = "1"
	data2.ApplicationOwner = "admin"
	data2.ApiCreator = "admin"
	data2.ApiCreatorTenantDomain = ""
	data2.ApiTier = tokenData.apiTier
	data2.Username = "admin2"
	data2.UserTenantDomain = tokenData.userTenantDomain
	data2.ThrottledOut = tokenData.throttledOut
	data2.ServiceTime = tokenData.serviceTime
	data2.ApiContext = tokenData.apiContext
	data2.ApiName = "PizzaShackAPI"
	data2.ApiVersion = "1.0.0"
	data2.ApiResourcePath = req.Attributes.Request.Http.Path
	data2.Protocol = req.Attributes.Request.Http.Protocol
	data2.ApiMethod = "GET"
	data2.UserAgent = req.Attributes.Request.Http.Headers["user-agent"]
	data2.Hostname = strings.Split(req.Attributes.Request.Http.Host, ":")[0]
	data2.RequestTimestamp = 1000 * (req.Attributes.Request.Time.Seconds) //RequestTimestamp in milli sec
	data2.GatewayType = gatewayType
	data2.Label = gatewayType
	data2.Subscriber = ""
	data2.ThrottledOutReason = ""
	data2.ThrottledOutTimestamp = 0
	data2.ErrorCode = ""
	data2.ErrorMessage = ""

	config := &tls.Config{
		InsecureSkipVerify: true,
	}


	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(config)),grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := analytics.NewAnalyticsSendServiceClient(conn)
	c , err2 := client.SendAnalytics(context.Background())
	if err2 != nil {
		log.Fatalf("did not connect2: %v", err2)
	}

	log.Println(data2)
	err3 := c.Send(data2)

	if err3 != nil {
		log.Fatalf("did not send: %v", err3)
	}

	resp := &ext_authz.CheckResponse{}
	resp = &ext_authz.CheckResponse{
		Status: &status.Status{Code: int32(rpc.OK)},
		HttpResponse: &ext_authz.CheckResponse_OkResponse{
			OkResponse: &ext_authz.OkHttpResponse{

			},
		},
	}
	return resp, nil
}
