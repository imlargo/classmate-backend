package config

type AppConfig struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host string
	Port string
}

var Config AppConfig

func LoadConfig() {
	err := loadEnv()
	if err != nil {
		panic("Error loading environment variables: " + err.Error())
	}

	Config = AppConfig{
		Server: ServerConfig{
			Host: GetEnv(API_URL, "localhost"),
			Port: GetEnv(PORT, "8000"),
		},
	}
}
