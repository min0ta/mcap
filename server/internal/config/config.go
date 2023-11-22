package config

type Config struct {
	ServerPort      string
	JWT_SIGNING_KEY string
	JWT_EXPIRES     uint
	PATH_TO_JSON_DB string
}

func New() *Config {
	return &Config{
		ServerPort:      ":8080",
		JWT_SIGNING_KEY: "jwt signing key! should be random string of characters!",
		JWT_EXPIRES:     86400,
		PATH_TO_JSON_DB: "./config/db.json",
	}
}
