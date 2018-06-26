package conf

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
)

func ReadConfig(path string) Config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var config Config
	json.Unmarshal(file, &config)
	return config
}
