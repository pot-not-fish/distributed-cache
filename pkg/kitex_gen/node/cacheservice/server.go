// Code generated by Kitex v0.7.1. DO NOT EDIT.
package cacheservice

import (
	server "github.com/cloudwego/kitex/server"
	node "kv-cache/pkg/kitex_gen/node"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler node.CacheService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}