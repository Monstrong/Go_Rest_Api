package config

type Config struct {
	Env string `yaml:"env" env:"ENV" env-default:"local"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-default:".storage/storage.db"`
	HTTPServer struct {
		Address string `yaml:"address" env:"HTTP_SERVER_ADDRESS" env-default:"localhost:8082"`
		Timeout string `yaml:"timeout" env:"HTTP_SERVER_TIMEOUT" env-default:"4s"`
		IdleTimeout string `yaml:"idle_timeout" env:"HTTP_SERVER_IDLE_TIMEOUT" env-default:"60s"`
	} 
}
