package internal

import (
	"fmt"
	"net/http"

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

func RunServer(cmd *cobra.Command, args []string) {
	if err := boot.InitAll(); err != nil {
		panic(err)
	}

	dbConf, _ := cmd.Flags().GetString("database")
	if err := orm.InitDb(dbConf); err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer func() {
		db, _ := orm.GetDb().DB()
		db.Close()
	}()

	if err := boot.Check(); err != nil {
		panic(err)
	}

	port, _ := cmd.Flags().GetString("port")

	mux := http.NewServeMux()
	sessionSvc := Session.NewSessionSvc()
	interceptors := connect.WithInterceptors(middleware.NewLogRouteUnaryInterceptor())
	mux.Handle(
		protoconnect.NewMainServiceHandler(
			controller.NewServer(sessionSvc),
			interceptors,
		),
	)

	// Setup reflection with the specified services
	reflector := grpcreflect.NewStaticReflector(
		"proto.MainService", // Replace with actual service names
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// Set up CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	// Start the server with CORS and reflection support
	err := http.ListenAndServe(
		fmt.Sprintf("localhost:%s", port),
		c.Handler(h2c.NewHandler(mux, &http2.Server{})),
	)
	if err != nil {
		log.Fatal().Msgf("listen failed: %v", err)
	}

	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatal().Msgf("Failed to listen: %v", err)
	// }

	// s := grpc.NewServer(
	// 	grpc.UnaryInterceptor(middleware.UnaryServerInterceptor),
	// 	grpc.StreamInterceptor(middleware.StreamServerInterceptor),
	// )
	// reflection.Register(s)
	// sessionSvc := Session.NewSessionSvc()
	// proto.RegisterMainServer(s, controller.NewServer(sessionSvc))
	// log.Info().Msgf("Listen to %s", port)
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatal().Msgf("Failed to serve: %v", err)
	// }
}
