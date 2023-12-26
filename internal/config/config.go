package config

import "github.com/caarlos0/env"

type Config struct {
	ServerAddr  string `env:"SERVER_ADDR" envDefault:":8083"`
	SearchAddr  string `env:"SEARCH_ADDR" envDefault:":8082"`
	ParserAddr  string `env:"PARSER_ADDR" envDefault:":8081"`
	TLSEnabled  bool   `env:"TLS" envDefault:"false"`
	TLSCertFile string `env:"TLS_CERT_FILE" envDefault:""`
	TLSKeyFile  string `env:"TLS_KEY_FILE" envDefault:""`
}

func ReadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
