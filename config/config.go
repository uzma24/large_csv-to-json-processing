package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	LargeCSVURI    string `mapstructure:"LARGE_CSV_PATH"`
	SmallCSVURI    string `mapstructure:"SMALL_CSV_PATH"`
	SmallestCSVURI string `mapstructure:"SMALLEST_CSV_PATH"`
	LargeJSONURI   string `mapstructure:"LARGE_CSV_TO_JSON_PATH"`
	SmallJSONURI   string `mapstructure:"SMALL_CSV_TO_JSON_PATH"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("error in reading config:::", err)
		return
	}
	err = viper.Unmarshal(&config)
	fmt.Println(viper.Unmarshal(&config))

	return
}
