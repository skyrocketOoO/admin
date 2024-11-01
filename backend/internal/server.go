package internal

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"admin/internal/boot"
	"admin/internal/controller"
	"admin/internal/middleware"
	"admin/internal/service/Session"
	"admin/internal/service/orm"
	"admin/proto/protoconnect"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// RunServer initializes and runs the gRPC server
func RunServer(cmd *cobra.Command, args []string) {
	if err := boot.InitAll(); err != nil {
		log.Fatal().Msgf("Initialization failed: %v", err)
	}

	dbConf, err := cmd.Flags().GetString("database")
	if err != nil {
		log.Fatal().Msgf("Failed to get database config: %v", err)
	}
	if err := initDatabase(dbConf); err != nil {
		log.Fatal().Msgf("Database initialization failed: %v", err)
	}
	defer closeDatabase()

	if err := boot.Check(); err != nil {
		log.Fatal().Msgf("System check failed: %v", err)
	}

	port, err := cmd.Flags().GetString("port")
	if err != nil || port == "" {
		port = "50051"
	}

	mux := setupMux()

	corsHandler := configureCORS()

	server := &http.Server{
		Addr:    port,
		Handler: corsHandler.Handler(h2c.NewHandler(mux, &http2.Server{})),
	}
	startServer(server)
}

// initDatabase initializes the database connection
func initDatabase(config string) error {
	return orm.InitDb(config)
}

// closeDatabase closes the database connection
func closeDatabase() {
	db, _ := orm.GetDb().DB()
	db.Close()
}

// setupMux sets up the HTTP request multiplexer
func setupMux() *http.ServeMux {
	mux := http.NewServeMux()
	sessionSvc := Session.NewSessionSvc()

	interceptors := connect.WithInterceptors(middleware.NewLogRouteUnaryInterceptor())
	mux.Handle(protoconnect.NewMainServiceHandler(controller.NewServer(sessionSvc), interceptors))

	// Set up reflection with specified services
	reflector := grpcreflect.NewStaticReflector("proto.MainService")
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	return mux
}

// configureCORS sets up CORS middleware configuration
func configureCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
}

// startServer runs the server and handles graceful shutdown
func startServer(server *http.Server) {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("Server listen failed: %v", err)
		}
	}()

	log.Info().Msg("Server started, awaiting connections...")

	// Wait for interrupt signal to gracefully shut down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Server shutdown failed: %v", err)
	}
	log.Info().Msg("Server exited gracefully")
}
