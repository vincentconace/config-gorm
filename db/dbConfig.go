package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/spf13/viper"
)

var DB *gorm.DB

type DbConfig struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	TimeZone     string
	print_log    bool
}

func Configure(ConfigPath string, ConfigName string) DbConfig {
	config.AddConfigPath(ConfigPath)
	config.SetConfigFile(ConfigName)
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if config.GetString("host") == "" {
		fmt.Println("falta host en el archivo " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("database") == "" {
		fmt.Println("falta database en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("user") == "" {
		fmt.Println("falta user en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("password") == "" {
		fmt.Println("falta password en el archivo de config " + ConfigName)
		os.Exit(1)
	}

	response := DbConfig{
		config.GetString("host"),
		config.GetString("port"),
		config.GetString("database"),
		config.GetString("user"),
		config.GetString("password"),
		config.GetString("charset"),
		config.GetInt("MaxIdleConns"),
		config.GetInt("MaxOpenConns"),
		"",
		config.GetBool("sql_log"),
	}

	return response
}
