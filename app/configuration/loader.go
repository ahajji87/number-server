package configuration

import (
	"fmt"
	"io/ioutil"
	"number-server/app"

	"gopkg.in/yaml.v2"
)

func Load(file string) app.Config {
	config := app.Config{}

	if file == "" {
		panic("ERROR: Missing mandatory config.yml file, use --config option")
	}

	config = loadFromFile(file, config)

	return config
}

func loadFromFile(file string, config app.Config) app.Config {
	content, err := ioutil.ReadFile(file)

	if err != nil {
		panic(fmt.Sprintf("Error reading config file: %v", err))
	}

	if err = yaml.Unmarshal(content, &config); err != nil {
		panic(fmt.Sprintf("Error composing configuration: %v", err))
	}

	return config
}
