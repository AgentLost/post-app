package config

type Config struct {
	KafkaAddr string `env:"K_ADDR"`
	HttpPort  string `env:"HTTP_PORT"`
}
