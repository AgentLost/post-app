package config

type Config struct {
	HttpPort   string `env:"HTTP_PORT"`
	KafkaUrl   string `env:"K_URL"`
	StorageUrl string `env:"STORAGE_URL"`
}
