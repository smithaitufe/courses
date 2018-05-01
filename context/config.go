package context

import "github.com/koding/multiconfig"

type Config struct {
	TokenSigning TokenSigning
	App          App
	DB           DB
	Logger       Logger
}
type TokenSigning struct {
	Secret     string
	Expiration int
}
type App struct {
	Name    string
	Version string
	Port    string
	Host    string
}
type DB struct {
	Engine   string `toml:"engine"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Name     string `toml:"name"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type Logger struct {
	DebugMode bool
	Format    string
}

func LoadConfig() *Config {
	config := &Config{}

	m := multiconfig.NewWithPath("config.toml")
	m.MustLoad(config)

	return config
}
