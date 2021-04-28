package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Open parses the configurations file (or creates it)
func Open() (config SingroConfig, err error) {
	filename, err := Path()
	if err != nil {
		return
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return config, err
		}
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return
	}

	return
}

// Path returns the absolute path of the macros definition
func Path() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.singro.yml", homedir), nil
}
