package initialize

import (
	"admin/internal/global"
	"admin/internal/model"
	"admin/internal/service/orm"
	"admin/internal/utils/password"

	"github.com/rs/zerolog/log"
	"github.com/skyrocketOoO/GoUtils/Struct"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var Cmd = &cobra.Command{
	Use:   "initialize",
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

		masterRole := model.Role{
			Name: global.MasterRoleName,
		}
		masterRole.ID = global.MasterRoleID
		if err := Struct.SetBoolFieldsTrue(&masterRole); err != nil {
			panic(err)
		}
		masterRole.DeletedAt = gorm.DeletedAt{}
		if err := db.Create(&masterRole).Error; err != nil {
			panic(err)
		}
		defaultRole := model.Role{
			Name: global.DefaultRoleName,
		}
		defaultRole.ID = global.DefaultRoleID
		if err := db.Create(&defaultRole).Error; err != nil {
			panic(err)
		}

		salt, err := password.GenSalt()
		if err != nil {
			panic(err)
		}
		masterAccount := &model.Account{
			UserName:    "admin",
			HashPass:    password.Hash("admin", salt),
			Salt:        salt,
			DisplayName: "admin",
			RoleID:      global.MasterRoleID,
		}
		if err = orm.GetDb().Create(masterAccount).Error; err != nil {
			panic(err)
		}
	},
}
