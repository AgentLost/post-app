package config

type Config struct {
	KafkaAddr string `env:"K_ADDR"`

	HttpPort string `env:"HTTP_PORT"`

	Port       string `env:"DB_PORT"`
	Host       string `env:"DB_HOST"`
	User       string `env:"DB_USER"`
	Name       string `env:"DB_NAME"`
	SSLMode    string `env:"DB_SSLMODE"`
	Password   string `env:"DB_PASSWORD"`
	DriverName string `env:"DB_DRIVER"`

	Dir string `env:"MIGRATE_DIR"`
}
