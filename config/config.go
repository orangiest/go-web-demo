package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func InitConfig(filename string) error {

	fmt.Println("文件名： ",filename)
	splits := strings.Split(filepath.Base(filename), ".")
	viper.SetConfigName(splits[0])
	viper.AddConfigPath(filepath.Dir(filename))
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func GetString(key string) string {

	if !viper.IsSet(key) {
		fmt.Printf("Configuration key %s not found.\n", key)
		os.Exit(1)
	}

	return viper.GetString(key)

}
