package tomlParser

import (
	toml "github.com/pelletier/go-toml/v2"

	"core/pkg/helper/errorhandler"

	"os"
)

func Decode(fileToParser string, ref interface{}) {
	file := loadTomlFile(fileToParser)
	err := toml.Unmarshal([]byte(file), ref)
	errorhandler.Check(err)
}

func loadTomlFile(filePath string) []byte {
	fileData, err := os.ReadFile(filePath)
	errorhandler.Check(err)

	return fileData
}
