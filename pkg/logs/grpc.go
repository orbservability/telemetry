package logs

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// Ensure functions match the gRPC interceptor types
var _ grpc.UnaryClientInterceptor = UnaryClientInterceptor
var _ grpc.StreamClientInterceptor = StreamClientInterceptor
var _ grpc.UnaryServerInterceptor = UnaryServerInterceptor
var _ grpc.StreamServerInterceptor = StreamServerInterceptor

func UnaryClientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	if err := preIntercept(ctx, method); err != nil {
		return err
	}

	// Make the RPC
	err := invoker(ctx, method, req, reply, cc, opts...)

	return postIntercept(ctx, method, err)
}

func StreamClientInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	if err := preIntercept(ctx, method); err != nil {
		return nil, err
	}

	// Make the RPC
	stream, err := streamer(ctx, desc, cc, method, opts...)

	return stream, postIntercept(ctx, method, err)
}

func UnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if err := preIntercept(ctx, info.FullMethod); err != nil {
		return nil, err
	}

	// Call the handler to proceed with the RPC
	resp, err := handler(ctx, req)

	return resp, postIntercept(ctx, info.FullMethod, err)
}

func StreamServerInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	if err := preIntercept(ss.Context(), info.FullMethod); err != nil {
		return err
	}

	// Call the handler to proceed with the RPC
	err := handler(srv, ss)

	return postIntercept(ss.Context(), info.FullMethod, err)
}

func preIntercept(ctx context.Context, method string) error {
	log.Debug().
		Str("rpc", method).
		Send()

	if ctx.Err() != nil {
		log.Error().
			Err(ctx.Err()).
			Str("rpc", method).
			Msg("Context error before RPC")
		return ctx.Err()
	}

	return nil
}

func postIntercept(ctx context.Context, method string, err error) error {
	if err != nil {
		log.Error().
			Err(err).
			Str("rpc", method).
			Msg("Error in RPC")
		return err
	} else if ctx.Err() != nil {
		log.Error().
			Err(ctx.Err()).
			Str("rpc", method).
			Msg("Context error after RPC")
		return ctx.Err()
	}

	return nil
}
