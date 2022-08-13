package cfg

import "github.com/spf13/viper"

type Config struct {
	PathFile       string
	OutputFilePath string
	OutputFormat   string
}

func NewConfig() *Config {
	v := viper.New()

	// set config file name
	v.SetConfigName("config")

	// set config file path
	v.AddConfigPath("./config")

	// set config file type
	v.SetConfigType("yml")

	// for use env variables
	v.AutomaticEnv()

	// set default configs
	setDefaultConfig(v)

	return &Config{
		PathFile:       v.GetString("csv_file_name"),
		OutputFilePath: v.GetString("output_file_path"),
		OutputFormat:   v.GetString("output_format"),
	}
}

func setDefaultConfig(v *viper.Viper) {
	v.SetDefault("csv_file_name", "paths.csv")
	v.SetDefault("output_format", "file")
	v.SetDefault("output_file_path", "result.txt")
}
