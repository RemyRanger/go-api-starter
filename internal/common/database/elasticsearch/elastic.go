package elasticsearch

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/olivere/elastic"
	"gitlab.com/cdv-projects/go-apis/internal/common/config"
	"gitlab.com/cdv-projects/go-apis/internal/common/logs"
)

// New : func new
func New() *elastic.Client {
	es, err := elastic.NewSimpleClient(
		elastic.SetURL(config.Conf.Es.Addr),
		elastic.SetInfoLog(new(logs.ElasticInfoLogger)),
		elastic.SetErrorLog(new(logs.ElasticErrorLogger)),
		elastic.SetTraceLog(new(logs.ElasticTraceLogger)),
		elastic.SetGzip(true),
		elastic.SetBasicAuth(config.Conf.Es.User, config.Conf.Es.Password),
	)

	if err != nil {
		log.Fatal().Err(err).Msg("Error creating elastic client")
	}

	info, _, err := es.Ping(config.Conf.Es.Addr).Do(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Error pinging elastic client")
	}
	log.Info().Msgf("Elasticsearch Version : %s\n", info.Version.Number)

	return es
}
