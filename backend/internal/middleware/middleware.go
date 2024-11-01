package middleware

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/rs/zerolog/log"
)

func NewLogRouteUnaryInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			// Call the next handler to get the response and error
			resp, err := next(ctx, req)

			// Log the method name and response status
			method := req.Spec().Procedure // Retrieve the fully qualified method name
			statusCode := connect.CodeOf(err)

			// Log with error handling and status
			if err != nil {
				log.Info().Msgf("%s %s Error: %v", method, colorizeStatus(statusCode.String()), err)
			} else {
				log.Info().Msgf("%s %s", method, colorizeStatus(statusCode.String()))
			}

			return resp, err
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
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
