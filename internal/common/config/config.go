package config

import (
	"os"
	"path/filepath"
	"reflect"

	"github.com/rs/zerolog/log"

	"gopkg.in/yaml.v2"
)

var Conf *Configs

// Configs : struct Configs
type Configs struct {
	Adequation Adequation    `yaml:"adequation"`
	Telemetry  Telemetry     `yaml:"telemetry"`
	Es         ElasticSearch `yaml:"elasticsearch"`
	Srv        Server        `yaml:"srv"`
	Mysql      Mysql         `yaml:"mysql"`
	Yanport    JWTClient     `yaml:"yanport"`
	Cadredevie Cadredevie    `yaml:"cadredevie"`
	Firebase   Firebase      `yaml:"firebase"`
	Logs       Logs          `yaml:"logs"`
	Security   Security      `yaml:"security"`
	Mail       Mail          `yaml:"mail"`
}

// Server : struct Server
type Server struct {
	Addr         string `yaml:"addr"`
	Certificate  string `yaml:"certificate"`
	Key          string `yaml:"key"`
	ImageBaseURL string `yaml:"imageBaseURL"`
	ProductID    string `yaml:"product_id"`
}

// Telemetry : struct Telemetry
type Telemetry struct {
	ServiceName              string `yaml:"service_name"`
	SignozAccessToken        string `yaml:"signoz_access_token"`
	OtelExporterOTLPEndpoint string `yaml:"otel_exporter_otlp_endpoint"`
}

// Logs : struct Logs
type Logs struct {
	Level string `yaml:"level"`
}

// Security : struct Security
type Security struct {
	Secret string `yaml:"secret"`
}

// Mysql : struct Mysql
type Mysql struct {
	DSN  string `yaml:"dsn"`
	DSN2 string `yaml:"dsn2"`
}

// Firebase : struct Firebase
type Firebase struct {
	CredentialsFile string `yaml:"credentials_file"`
}

// ElasticSearch : struct ElasticSearch
type ElasticSearch struct {
	Addr            string `yaml:"addr"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	ProgramsIndex   string `yaml:"programsIndex"`
	PropertiesIndex string `yaml:"propertiesIndex"`
}

// Cadredevie : struct Cadredevie
type Cadredevie struct {
	RealEstate RealEstate `yaml:"realestate"`
}

// RealEstate : struct RealEstate
type RealEstate struct {
	BaseURL string `yaml:"base_url"`
	Apikey  string `yaml:"apikey"`
}

// Adequation : struct Adequation
type Adequation struct {
	Programs       JWTClient `yaml:"programs"`
	Attractiveness JWTClient `yaml:"attractiveness"`
	Commerce       JWTClient `yaml:"commerce"`
	Insee          JWTClient `yaml:"insee"`
	Promoters      JWTClient `yaml:"promoters"`
	DVF            DVF       `yaml:"dvf"`
}

// JWTClient : struct JWTClient
type JWTClient struct {
	BaseURL string `yaml:"base_url"`
	JWT     string `yaml:"jwt"`
}

// DVF : struct DVF
type DVF struct {
	BaseURL  string `yaml:"base_url"`
	DVFIndex string `yaml:"dvf_index"`
}

// Mail : struct Mail
type Mail struct {
	SendgridApiKey string `yaml:"sendgrid_api_key"`
	AllowedOrigin  string `yaml:"allowed_origin"`
	MailFrom       string `yaml:"mail_from"`
	MailTo         string `yaml:"mail_to"`
}

// New : func new
func New(file string) *Configs {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal().Err(err).Msgf("Homedir error")
	}

	path := filepath.Join(home, ".config", "cadredevie", file)

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal().Err(err).Msgf("Read config error")
	}

	var cfg Configs
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatal().Err(err).Msgf("Unmarshal config error")
	}

	return &cfg
}

// CheckRealEstate check if mandatory config is ok
func (c *Configs) CheckRealEstate() {
	if reflect.DeepEqual(c.Srv, Server{}) {
		log.Fatal().Msgf("Missing server config in realestate.yaml")
	}

	if reflect.DeepEqual(c.Mysql, Mysql{}) {
		log.Fatal().Msgf("Missing mysql config in realestate.yaml")
	}

	if reflect.DeepEqual(c.Es, ElasticSearch{}) {
		log.Fatal().Msgf("Missing elasticsearch config in realestate.yaml")
	}
}
