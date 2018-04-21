package context

import "github.com/koding/multiconfig"

// Config : configuration struct mapping to config.toml
type Config struct {
	DB DB
}

// DB : connection info to the database server
type DB struct {
	Engine   string `toml:"engine"`
	Port     string `toml:"port"`
	Host     string `toml:"host"`
	Name     string `toml:"name"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

// LoadConfig : load configuration from config.toml file
func LoadConfig() *Config {
	config := &Config{}

	m := multiconfig.NewWithPath("config.toml")
	m.MustLoad(config)

	return config
}
