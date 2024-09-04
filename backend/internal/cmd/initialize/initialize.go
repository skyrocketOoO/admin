package initialize

import (
	"admin/internal/global"
	"admin/internal/model"
	"admin/internal/service/orm"

	"github.com/rs/zerolog/log"
	"github.com/skyrocketOoO/GoUtils/Struct"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Init the data, config...etc",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dbConf, _ := cmd.Flags().GetString("database")
		if err := orm.InitDb(dbConf); err != nil {
			log.Fatal().Msg(err.Error())
		}
		defer func() {
			db, _ := orm.GetDb().DB()
			db.Close()
		}()

		db := orm.GetDb()
		masterRole := model.Role{}
		masterRole.ID = global.MasterRoleID
		if err := Struct.SetBoolFieldsTrue(masterRole); err != nil {
			panic(err)
		}
		if err := db.Create(&masterRole).Error; err != nil {
			panic(err)
		}
		defaultRole := model.Role{}
		defaultRole.ID = global.DefaultRoleID
		if err := db.Create(&defaultRole).Error; err != nil {
			panic(err)
		}
	},
}

func init() {
}
