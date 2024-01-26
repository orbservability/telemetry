package traces

import (
	"context"

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
	if ctx.Err() != nil {
		// TODO: error trace
		ctx.Err()
	}

	return nil
}

func postIntercept(ctx context.Context, method string, err error) error {
	if err != nil {
		// TODO: error trace
		return err
	} else if ctx.Err() != nil {
		// TODO: error trace
		return ctx.Err()
	}

	return nil
}
