/*
 *  Copyright (c) 2018, https://github.com/airwide-code/airwide.datacenter
 *  All rights reserved.
 *
 *
 *
 */

package grpc_recovery2

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

// Initialization shows an initialization sequence with a custom recovery handler func.
func NewRecoveryServer(unaryCustomFunc UnaryRecoveryHandlerFunc, streamCustomFunc StreamRecoveryHandlerFunc) *grpc.Server {
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []Option{
		// grpc_recovery2.WithUnaryRecoveryHandler(unaryCustomFunc),
		WithUnaryRecoveryHandler(unaryCustomFunc),
		WithStreamRecoveryHandler(streamCustomFunc),
	}
	// Create a server. Recovery handlers should typically be last in the chain so that other middleware
	// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			UnaryServerInterceptor(opts...),
		),
		grpc_middleware.WithStreamServerChain(
			StreamServerInterceptor(opts...),
		),
	)
	return server
}

// Initialization shows an initialization sequence with a custom recovery handler func.
func NewRecoveryServer2(unaryCustomFunc UnaryRecoveryHandlerFunc, unaryCustomFunc2 UnaryRecoveryHandlerFunc, streamCustomFunc StreamRecoveryHandlerFunc) *grpc.Server {
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []Option{
		// grpc_recovery2.WithUnaryRecoveryHandler(unaryCustomFunc),
		WithUnaryRecoveryHandler(unaryCustomFunc),
		WithUnaryRecoveryHandler2(unaryCustomFunc2),
		WithStreamRecoveryHandler(streamCustomFunc),
	}
	// Create a server. Recovery handlers should typically be last in the chain so that other middleware
	// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			UnaryServerInterceptor(opts...),
		),
		grpc_middleware.WithStreamServerChain(
			StreamServerInterceptor(opts...),
		),
	)
	return server
}
