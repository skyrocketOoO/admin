package internal

import (
	"net"

	"admin/api"
	"admin/internal/boot"
	"admin/internal/controller"
	"admin/internal/service/Session"
	"admin/internal/service/orm"

	"google.golang.org/grpc"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Msgf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sessionSvc := Session.NewSessionSvc()
	api.RegisterMainServer(s, controller.NewServer(sessionSvc))
	log.Info().Msgf("Listen to %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("Failed to serve: %v", err)
	}
}
