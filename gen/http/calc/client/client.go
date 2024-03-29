// Code generated by goa v3.11.1, DO NOT EDIT.
//
// calc client HTTP transport
//
// Command:
// $ goa gen github.com/sraynitjsr/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the calc service endpoint HTTP clients.
type Client struct {
	// Add Doer is the HTTP client used to make requests to the add endpoint.
	AddDoer goahttp.Doer

	// Sub Doer is the HTTP client used to make requests to the sub endpoint.
	SubDoer goahttp.Doer

	// Multiply Doer is the HTTP client used to make requests to the multiply
	// endpoint.
	MultiplyDoer goahttp.Doer

	// Divide Doer is the HTTP client used to make requests to the divide endpoint.
	DivideDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the calc service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AddDoer:             doer,
		SubDoer:             doer,
		MultiplyDoer:        doer,
		DivideDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Add returns an endpoint that makes HTTP requests to the calc service add
// server.
func (c *Client) Add() goa.Endpoint {
	var (
		decodeResponse = DecodeAddResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "add", err)
		}
		return decodeResponse(resp)
	}
}

// Sub returns an endpoint that makes HTTP requests to the calc service sub
// server.
func (c *Client) Sub() goa.Endpoint {
	var (
		decodeResponse = DecodeSubResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSubRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SubDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "sub", err)
		}
		return decodeResponse(resp)
	}
}

// Multiply returns an endpoint that makes HTTP requests to the calc service
// multiply server.
func (c *Client) Multiply() goa.Endpoint {
	var (
		decodeResponse = DecodeMultiplyResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildMultiplyRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.MultiplyDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "multiply", err)
		}
		return decodeResponse(resp)
	}
}

// Divide returns an endpoint that makes HTTP requests to the calc service
// divide server.
func (c *Client) Divide() goa.Endpoint {
	var (
		decodeResponse = DecodeDivideResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDivideRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DivideDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "divide", err)
		}
		return decodeResponse(resp)
	}
}
