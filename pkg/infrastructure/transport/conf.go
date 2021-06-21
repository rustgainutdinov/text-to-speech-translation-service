package transport

import "github.com/kelseyhightower/envconfig"

const appID = "translationservice"

type Config struct {
	GRPCAddress           string `envconfig:"grpc_address"`
	HTTPProxyAddress      string `envconfig:"http_proxy_address"`
	BalanceServiceAddress string `envconfig:"balance_service_address"`
	DBUser                string `envconfig:"db_user"`
	DBName                string `envconfig:"db_name"`
	DBPort                string `envconfig:"db_port"`
	DBPass                string `envconfig:"db_pass"`
	DBHost                string `envconfig:"db_host"`
	RabbitMqUser          string `envconfig:"rabbitmq_user"`
	RabbitMqPass          string `envconfig:"rabbitmq_pass"`
	RabbitMqHost          string `envconfig:"rabbitmq_host"`
}

func ParseEnv() (*Config, error) {
	c := new(Config)
	if err := envconfig.Process(appID, c); err != nil {
		return nil, err
	}
	return c, nil
}
