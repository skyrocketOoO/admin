package internal

import (
	restapi "admin/api/rest"
	"admin/internal/boot"
	"admin/internal/controller/rest"
	"admin/internal/controller/rest/middleware"
	"admin/internal/service/dao"
	"admin/internal/service/orm"
	"admin/internal/usecase"

	errors "github.com/rotisserie/eris"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func RunServer(cmd *cobra.Command, args []string) {
	if err := boot.InitAll(); err != nil {
		panic(err.Error())
	}

	dbConf, _ := cmd.Flags().GetString("database")
	db, err := orm.NewDB(dbConf)
	if err != nil {
		log.Fatal().Msg(errors.ToString(err, true))
	}
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	dao := dao.NewDao(db)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	usecase := usecase.NewBasicUsecase(dao)
	restController := rest.NewRestController(usecase)

	router := gin.Default()
	router.Use(middleware.CORS())
	restapi.Binding(router, restController)

	port, _ := cmd.Flags().GetString("port")
	router.Run(":" + port)
}
