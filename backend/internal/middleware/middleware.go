package middleware

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor logs the API calls and status code
func UnaryServerInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	// Call the handler to execute the actual gRPC method
	resp, err := handler(ctx, req)

	// Extract the gRPC status code
	st, _ := status.FromError(err)

	// Log the response, error, and status code
	if err != nil {
		log.Info().Msgf("%s %s Error: %v", info.FullMethod, colorizeStatus(st.Code().String()), err)
	} else {
		log.Info().Msgf("%s %s", info.FullMethod, colorizeStatus(st.Code().String()))
	}

	return resp, err
}

// StreamServerInterceptor logs the API calls, metadata, and status code for streaming RPCs
func StreamServerInterceptor(
	srv any,
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	// Retrieve and log metadata (optional)
	md, ok := metadata.FromIncomingContext(ss.Context())
	if ok {
		log.Info().Msgf("Metadata: %v", md)
	}

	// Call the handler to process the stream
	err := handler(srv, ss)

	// Extract the gRPC status code
	st, _ := status.FromError(err)

	// Log the error and status code if any
	if err != nil {
		log.Info().Msgf("%s, Status: %s, Error: %v", info.FullMethod, colorizeStatus(st.Code().String()), err)
	} else {
		log.Info().Msgf("%s, Status: %s", info.FullMethod, colorizeStatus(st.Code().String()))
	}

	return err
}

func colorizeStatus(statusCode string) string {
	switch statusCode {
	case "OK":
		return fmt.Sprintf("\033[32m%s\033[0m", statusCode) // Green for OK
	case "INVALID_ARGUMENT":
		return fmt.Sprintf("\033[33m%s\033[0m", statusCode) // Yellow for warnings
	case "INTERNAL":
		return fmt.Sprintf("\033[31m%s\033[0m", statusCode) // Red for errors
	default:
		return fmt.Sprintf("\033[34m%s\033[0m", statusCode) // Blue for other status codes
	}
}
