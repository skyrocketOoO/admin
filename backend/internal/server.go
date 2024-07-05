package internal

import (
	"net"

	"admin/api"
	"admin/internal/boot"
	"admin/internal/controller"
	"admin/internal/service/orm"

	errors "github.com/rotisserie/eris"
	"google.golang.org/grpc"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func RunServer(cmd *cobra.Command, args []string) {
	if err := boot.InitAll(); err != nil {
		panic(err.Error())
	}

	dbConf, _ := cmd.Flags().GetString("database")
	if err := orm.InitDb(dbConf); err != nil {
		log.Fatal().Msg(errors.ToString(err, true))
	}
	defer func() {
		db, _ := orm.GetDb().DB()
		db.Close()
	}()

	port, _ := cmd.Flags().GetString("port")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Msgf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterAccountServiceServer(s, &controller.AccountServer{})
	log.Info().Msgf("Listen to %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("Failed to serve: %v", err)
	}
}
