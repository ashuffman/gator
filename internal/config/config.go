package config

import (
	"os"
	"encoding/json"
)

type Config struct {
	DBurl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	// use os.UserHomedir to get home directory path
	home, err := os.UserHomeDir()
	// check for errors, and return gracefully as needed
	if err != nil {
		return Config{}, err
	}
	// create file path from home and file strings
	filePath := home + "/" + configFileName

	// read file into variable
	fileBytes, err := os.ReadFile(filePath)
	// check for errors, and return gracefully as needed
	if err != nil {
		return Config{}, err
	}

	// unmarshal json from file to Config struct
	var config Config
	// check for unmarshaling errors, and return gracefully as needed
	if err := json.Unmarshal(fileBytes, &config); err != nil {
		return Config{}, err
	}

	// if we get here without error, return the unmarshalled config struct and nil
	return config, nil
}

func (c Config) SetUser(username string) error {
	// updated struct with new username
	c.CurrentUserName = username

	// write new struct to cfg file
	err := write(&c)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error){
	// use os.UserHomedir to get home directory path
	home, err := os.UserHomeDir()
	// check for errors, and return gracefully as needed
	if err != nil {
		return "" , err
	}
	// create file path from home and file strings
	filePath := home + "/" + configFileName
	
	// return the file path and nil error
	return filePath, nil
}

func write(cfg *Config) error {
	// get the config file path
	filePath, err := getConfigFilePath()
	// gracefully return an error if needed
	if err != nil {
		return err
	}
	
	// marshal cfg to json
	cfgBytes, err := json.Marshal(cfg)
	// check for errors, and handle gracefully as needed
	if err != nil {
		return err
	}
	// write json to file
	err = os.WriteFile(filePath, cfgBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
