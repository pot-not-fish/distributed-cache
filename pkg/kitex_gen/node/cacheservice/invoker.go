// Code generated by Kitex v0.7.1. DO NOT EDIT.

package cacheservice

import (
	server "github.com/cloudwego/kitex/server"
	node "kv-cache/pkg/kitex_gen/node"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler node.CacheService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
