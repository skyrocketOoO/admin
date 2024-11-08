package orm

import (
	"fmt"

	"admin/internal/model"

	errors "github.com/rotisserie/eris"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db    *gorm.DB
	Debug bool
)

func GetDb() *gorm.DB {
	if Debug {
		return db.Debug()
	}
	return db
}

func InitDb(database string) (err error) {
	config := gorm.Config{
		Logger: nil,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	}
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
			postgres.Open(connStr), &config,
		); err != nil {
			return
		}
	case "sqlite", "":
		log.Info().Msg("Connecting to Sqlite")
		if db, err = gorm.Open(sqlite.Open("sqlite.db"), &config); err != nil {
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
