package initializer

import (
	"fmt"
	"plbooking_go_structure1/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("./config/") // look for the path
	viper.SetConfigName("local")     // look if ppath has the .env name
	viper.SetConfigType("yaml")      // read only .env file

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal configuration struct: %w", err))
	}
	fmt.Println("loaded config successully")
}
