package mysql

import (
	"time"

	"github.com/rs/zerolog/log"

	"gitlab.com/cdv-projects/go-apis/internal/common/config"
	"gitlab.com/cdv-projects/go-apis/internal/common/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// New : create client to connect to database
func New() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Conf.Mysql.DSN), &gorm.Config{
		Logger:                 logs.NewMysqlLogger(),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating mysql client")
	}
	PingDB(db)

	// check secondary connection
	if config.Conf.Mysql.DSN2 != "" {
		err := db.Use(dbresolver.Register(dbresolver.Config{
			Sources: []gorm.Dialector{mysql.Open(config.Conf.Mysql.DSN2)},
		}, "secondary"))
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating mysql client")
		}

		PingSecondaryDB(db)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating mysql pool")
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(30 * time.Second)

	return db
}

// PingDB Ping Database by runing SELECT VERSION()
func PingDB(db *gorm.DB) {
	var version string
	db.Raw("SELECT VERSION()").Scan(&version)
	log.Info().Msgf("Mysql Version : %v", version)
}

// PingSecondaryDB Ping Secondary Database by runing SELECT VERSION()
func PingSecondaryDB(db *gorm.DB) {
	var version string
	db.Clauses(dbresolver.Use("secondary")).Raw("SELECT VERSION()").Scan(&version)
	log.Info().Msgf("Mysql Secondary Version : %v", version)
}
