package orm

import (
	"fmt"

	"admin/internal/service/orm/model"

	errors "github.com/rotisserie/eris"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	return db
}

func InitDb(database string) (err error) {
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
		return errors.New("database not supported")
	}
	if err = db.AutoMigrate(
		&model.Account{},
		&model.Role{},
		&model.Store{},
	); err != nil {
		return
	}

	return nil
}
