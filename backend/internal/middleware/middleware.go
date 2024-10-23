package middleware

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryServerInterceptor logs the API calls
func UnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("Unary API called: %s, Request: %v", info.FullMethod, req)

	// Call the handler to execute the actual gRPC method
	resp, err := handler(ctx, req)

	// Log the response and error if any
	if err != nil {
		log.Printf("Unary API Error: %s, Error: %v", info.FullMethod, err)
	} else {
		log.Printf("Unary API Response: %s, Response: %v", info.FullMethod, resp)
	}

	return resp, err
}

// StreamServerInterceptor logs the API calls for streaming RPCs
func StreamServerInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Printf("Streaming API called: %s", info.FullMethod)

	// Retrieve and log metadata (optional)
	md, ok := metadata.FromIncomingContext(ss.Context())
	if ok {
		log.Printf("Metadata: %v", md)
	}

	// Call the handler to process the stream
	err := handler(srv, ss)

	if err != nil {
		log.Printf("Streaming API Error: %s, Error: %v", info.FullMethod, err)
	} else {
		log.Printf("Streaming API completed: %s", info.FullMethod)
	}

	return err
}
