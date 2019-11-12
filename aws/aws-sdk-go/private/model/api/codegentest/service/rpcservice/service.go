// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rpcservice

import (
	"github.com/Netflix/titus-executor/aws/aws-sdk-go/aws"
	"github.com/Netflix/titus-executor/aws/aws-sdk-go/aws/client"
	"github.com/Netflix/titus-executor/aws/aws-sdk-go/aws/client/metadata"
	"github.com/Netflix/titus-executor/aws/aws-sdk-go/aws/request"
	v4 "github.com/Netflix/titus-executor/aws/aws-sdk-go/aws/signer/v4"
	"github.com/Netflix/titus-executor/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// RPCService provides the API operation methods for making requests to
// RPC Service. See this package's package overview docs
// for details on the service.
//
// RPCService methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type RPCService struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "RPCService" // Name of service.
	EndpointsID = "rpcservice" // ID to lookup a service endpoint with.
	ServiceID   = "RPCService" // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the RPCService client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a RPCService client from just a session.
//     svc := rpcservice.New(mySession)
//
//     // Create a RPCService client with additional configuration
//     svc := rpcservice.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *RPCService {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *RPCService {
	svc := &RPCService{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "0000-00-00",
				JSONVersion:   "1.1",
				TargetPrefix:  "RPCService_00000000",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	svc.Handlers.UnmarshalStream.PushBackNamed(jsonrpc.UnmarshalHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a RPCService operation and runs any
// custom request initialization.
func (c *RPCService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
