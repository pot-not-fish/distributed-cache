// Code generated by Kitex v0.7.1. DO NOT EDIT.

package cacheservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	node "kv-cache/pkg/kitex_gen/node"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetKey(ctx context.Context, request *node.GetKeyRequest, callOptions ...callopt.Option) (r *node.GetKeyResponse, err error)
	SetKey(ctx context.Context, request *node.SetKeyRequest, callOptions ...callopt.Option) (r *node.SetKeyResponse, err error)
	DelKey(ctx context.Context, request *node.DelKeyRequest, callOptions ...callopt.Option) (r *node.DelKeyResponse, err error)
	SetGroup(ctx context.Context, request *node.SetGroupRequest, callOptions ...callopt.Option) (r *node.SetGroupResponse, err error)
	GetAll(ctx context.Context, request *node.GetAllKeysRequest, callOptions ...callopt.Option) (r *node.GetAllKeysResponse, err error)
	GetAllGroup(ctx context.Context, request *node.GetAllGroupsRequest, callOptions ...callopt.Option) (r *node.GetAllGroupsResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCacheServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCacheServiceClient struct {
	*kClient
}

func (p *kCacheServiceClient) GetKey(ctx context.Context, request *node.GetKeyRequest, callOptions ...callopt.Option) (r *node.GetKeyResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetKey(ctx, request)
}

func (p *kCacheServiceClient) SetKey(ctx context.Context, request *node.SetKeyRequest, callOptions ...callopt.Option) (r *node.SetKeyResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetKey(ctx, request)
}

func (p *kCacheServiceClient) DelKey(ctx context.Context, request *node.DelKeyRequest, callOptions ...callopt.Option) (r *node.DelKeyResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DelKey(ctx, request)
}

func (p *kCacheServiceClient) SetGroup(ctx context.Context, request *node.SetGroupRequest, callOptions ...callopt.Option) (r *node.SetGroupResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetGroup(ctx, request)
}

func (p *kCacheServiceClient) GetAll(ctx context.Context, request *node.GetAllKeysRequest, callOptions ...callopt.Option) (r *node.GetAllKeysResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAll(ctx, request)
}

func (p *kCacheServiceClient) GetAllGroup(ctx context.Context, request *node.GetAllGroupsRequest, callOptions ...callopt.Option) (r *node.GetAllGroupsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllGroup(ctx, request)
}
