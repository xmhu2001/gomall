// Code generated by Kitex v0.9.1. DO NOT EDIT.

package echo

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	api "github.com/xmhu2001/gomall/demo/demo_thrift/kitex_gen/api"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"echo": kitex.NewMethodInfo(
		echoHandler,
		newEchoEchoArgs,
		newEchoEchoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	echoServiceInfo                = NewServiceInfo()
	echoServiceInfoForClient       = NewServiceInfoForClient()
	echoServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return echoServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return echoServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return echoServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "Echo"
	handlerType := (*api.Echo)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.EchoEchoArgs)
	realResult := result.(*api.EchoEchoResult)
	success, err := handler.(api.Echo).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEchoEchoArgs() interface{} {
	return api.NewEchoEchoArgs()
}

func newEchoEchoResult() interface{} {
	return api.NewEchoEchoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, req *api.Request) (r *api.Response, err error) {
	var _args api.EchoEchoArgs
	_args.Req = req
	var _result api.EchoEchoResult
	if err = p.c.Call(ctx, "echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
