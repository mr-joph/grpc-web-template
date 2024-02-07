package config

import (
	"core/pkg/helper/tomlParser"
	"strings"

	"flag"
)

const (
	WILDCARD_CHAR = "*"
)

// exports
var Env EnvConfig

var defaultConfigPath = flag.String(
	"config",
	"./config.toml",
	"path for toml config file",
)

// load config file when the module is loaded
func init() {
	// load the toml config file
	tomlParser.Decode(*defaultConfigPath, &Env)
	// parse data needed
	parseData()
}

func parseData() {
	// parse cors addresses allowed
	allowAddressesInOneLine(&Env)
}

// Converts the array in only one line separate by ","
// in order to check later easy
func allowAddressesInOneLine(envConfig *EnvConfig) {
	arr := envConfig.Cors.AllowAddresses

	// if it finds a wildcard '*' in the array
	// allow all addresses
	for _, address := range arr {
		if address == WILDCARD_CHAR {
			envConfig.Cors.AllowAddressesOneLine = WILDCARD_CHAR
			return
		}
	}

	envConfig.Cors.AllowAddressesOneLine = strings.Join(arr, ",")

}
