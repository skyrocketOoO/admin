package main

import (
	"github.com/rs/zerolog/log"

	"admin/internal/boot"
	"admin/internal/service/orm"

	"github.com/skyrocketOoO/gorm-enhance-plugin/columnname"
	"github.com/skyrocketOoO/gorm-enhance-plugin/tablename"
)

func main() {
	if err := boot.InitAll(); err != nil {
		log.Fatal().Msgf("Initialization failed: %v", err)
	}

	if err := orm.InitDb("sqlite"); err != nil {
		log.Fatal().Msgf("Database initialization failed: %v", err)
	}
	db := orm.GetDb()
	sqlDb, _ := db.DB()
	defer sqlDb.Close()

	if err := tablename.GenTableNamesCode(db, "gen/table/table.go"); err != nil {
		log.Fatal().Msgf("%v", err)
	}

	tablenames, err := tablename.GetTableNames(db)
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}
	if err := columnname.GenTableColumnNamesCode(db, tablenames, "gen/column/column.go"); err != nil {
		log.Fatal().Msgf("%v", err)
	}
}
