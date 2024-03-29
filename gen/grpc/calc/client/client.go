// Code generated by goa v3.11.1, DO NOT EDIT.
//
// calc gRPC client
//
// Command:
// $ goa gen github.com/sraynitjsr/design

package client

import (
	"context"

	calcpb "github.com/sraynitjsr/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli calcpb.CalcClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the calc service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: calcpb.NewCalcClient(cc),
		opts:    opts,
	}
}

// Add calls the "Add" function in calcpb.CalcClient interface.
func (c *Client) Add() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildAddFunc(c.grpccli, c.opts...),
			EncodeAddRequest,
			DecodeAddResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Sub calls the "Sub" function in calcpb.CalcClient interface.
func (c *Client) Sub() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildSubFunc(c.grpccli, c.opts...),
			EncodeSubRequest,
			DecodeSubResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Multiply calls the "Multiply" function in calcpb.CalcClient interface.
func (c *Client) Multiply() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMultiplyFunc(c.grpccli, c.opts...),
			EncodeMultiplyRequest,
			DecodeMultiplyResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Divide calls the "Divide" function in calcpb.CalcClient interface.
func (c *Client) Divide() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildDivideFunc(c.grpccli, c.opts...),
			EncodeDivideRequest,
			DecodeDivideResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
