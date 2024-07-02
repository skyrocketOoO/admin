package orm

import (
	"fmt"

	"web-server-template/internal/service/orm/model"

	errors "github.com/rotisserie/eris"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(database string) (db *gorm.DB, err error) {
	switch database {
	case "postgres":
		log.Info().Msg("Connecting to Postgres")
		connStr := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			viper.GetString("postgres.host"),
			viper.GetString("postgres.port"),
			viper.GetString("postgres.user"),
			viper.GetString("postgres.password"),
			viper.GetString("postgres.db"),
			viper.GetString("postgres.sslmode"),
			viper.GetString("postgres.timezone"),
		)
		if db, err = gorm.Open(
			postgres.Open(connStr), &gorm.Config{
				Logger: nil,
			},
		); err != nil {
			return
		}
	case "sqlite":
		log.Info().Msg("Connecting to Sqlite")
		if db, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{}); err != nil {
			return
		}
	default:
		return nil, errors.New("database not supported")
	}
	if err = db.AutoMigrate(
		&model.Account{},
		&model.Role{},
		&model.Permission{},
	); err != nil {
		return
	}

	return db, nil
}
