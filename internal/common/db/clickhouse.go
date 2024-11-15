package db

import (
	"context"
	"go-api-starter/internal/common/config"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
	"github.com/uptrace/go-clickhouse/chotel"
)

// New : create client to connect clickhouse database
func NewClickHouse(conf config.Config) *ch.DB {
	db := ch.Connect(
		ch.WithAddr(conf.Db.Addr),
		ch.WithUser("default"),
		ch.WithDatabase("default"),
		ch.WithTimeout(5*time.Second),
		ch.WithDialTimeout(5*time.Second),
		ch.WithReadTimeout(5*time.Second),
		ch.WithWriteTimeout(5*time.Second),
	)

	// Ping/Pong db
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal().Err(err)
	}
	log.Info().Msg("Connected to Clickhouse")

	// Log failed queries with app logger
	db.AddQueryHook(chdebug.NewQueryHook(chdebug.WithWriter(log.Logger)))
	// Trace requests with otel
	db.AddQueryHook(chotel.NewQueryHook())

	return db
}
