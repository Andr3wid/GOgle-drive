package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// initialize viper env reader
func initViper() {
	viper.AddConfigPath(".")
	viper.SetConfigName("env")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error while trying to read env config")
	}
}

func main() {
	initViper()

	fmt.Println()
}
