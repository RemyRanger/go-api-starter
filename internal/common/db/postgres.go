package db

import (
	"database/sql"
	"embed"
	"go-api-starter/internal/common/config"
	"go-api-starter/internal/common/logger"
	"time"

	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func NewPostgres(app_config config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(app_config.Db.Addr), &gorm.Config{
		Logger: logger.NewGormLogger(),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating postgres client")
	}

	// enable opentelemetry on ORM
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Fatal().Err(err).Msg("Error creating otelgorm plugin")
	}

	pingDB(db)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating postgres pool")
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(30 * time.Second)

	migrateSchemas(sqlDB)

	return db
}

func pingDB(db *gorm.DB) {
	var version string
	db.Raw("SELECT VERSION()").Scan(&version)
	log.Info().Msgf("Database version : %v", version)
}

func migrateSchemas(db *sql.DB) {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}
