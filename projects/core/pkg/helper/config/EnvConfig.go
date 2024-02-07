package config

type InfoConfig struct {
	Name    string
	Version string
	Mode    string
}

type ServerConfig struct {
	Port int
}

type WebserverConfig struct {
	Port int
}

type CorsConfig struct {
	AllowAddresses        []string `toml:"allow_addresses"`
	AllowAddressesOneLine string
}

type EnvConfig struct {
	Info      InfoConfig
	Server    ServerConfig
	Webserver WebserverConfig
	Cors      CorsConfig
}
