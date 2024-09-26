package shared

import (
	"fmt"
	"os"

	viper "github.com/spf13/viper"
)

func ReadConfigFile(file string) (*viper.Viper, error) {
	appConfig := viper.New()
	fmt.Println(file)

	appConfig.SetConfigType("yaml")

	configPath := "./env"
	if _, err := os.Stat(configPath); err == nil {
		appConfig.AddConfigPath(configPath)

		appConfig.SetConfigName("config")

		if commonErr := appConfig.ReadInConfig(); commonErr != nil {
			return nil, commonErr
		}

		if len(file) > 0 {
			appConfig.SetConfigName(file)
			fmt.Println("Loading config file: ", file)
			if err := appConfig.MergeInConfig(); err != nil {
				appConfig.SetConfigName("config")
				if err := appConfig.MergeInConfig(); err != nil {
					return nil, err
				}
			}
		}
	}
	appConfig.AutomaticEnv()

	return appConfig, nil

}