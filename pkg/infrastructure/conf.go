package infrastructure

import "github.com/kelseyhightower/envconfig"

const appID = "balanceservice"

type config struct {
	GRPCAddress      string `envconfig:"grpc_address" default:":8001"`
	HTTPProxyAddress string `envconfig:"http_proxy_address" default:":8000"`
	DBUser           string `envconfig:"db_user"`
	DBName           string `envconfig:"db_name"`
	DBPort           string `envconfig:"db_port"`
	DBPass           string `envconfig:"db_pass"`
	DBHost           string `envconfig:"db_host"`
}

func ParseEnv() (*config, error) {
	c := new(config)
	if err := envconfig.Process(appID, c); err != nil {
		return nil, err
	}
	return c, nil
}
