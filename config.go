package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	prefixFilename = "properties"
)

// Contains all properties of ConfFile
type Config struct {
	Auth Auth
	Api  Api `mapstructure:"api_url"`
}

// Config is used to specify credential to Netatmo API
// ClientID : Client ID from netatmo app registration at http://dev.netatmo.com/
// ClientSecret : Client app secret
// Username : Your netatmo account username
// Password : Your netatmo account password
type Auth struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	Username     string
	Password     string
}

type Api struct {
	Core string
	Auth string
	Home string
}

// fileName of config file in JSON format
// prefixFileName fileName without .env and .json
// return filePath of file
func GetConfig() Config {

	fileEnv := getFileName()
	viper.SetConfigName(fileEnv)

	viper.SetConfigType("yaml")

	pathOfFile := getPathFolderOfFile()
	viper.AddConfigPath(pathOfFile)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Missing Config File, %s", err))
	}

	var conf Config

	err = viper.Unmarshal(&conf)
	if err != nil {
		fmt.Println("Cannot bind struct with file", err)
	}
	return conf
}

// Manage fileName in terms of environment (i.e : test, development, production)
// return prefix.env
// if no env defined, return prefix.development
func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := prefixFilename + "." + env
	return filename
}

// Get Path absolute of folder of file.
// input : filename, to get path of file of parent folder
// return absolute path of folder
func getPathFolderOfFile() string {
	currentFolder, _ := os.Getwd()

	if strings.Contains(currentFolder, "utils") {
		folderConfig := strings.Replace(currentFolder, "utils", "config", 1)
		return folderConfig
	} else {
		return currentFolder + "/config"
	}

}
