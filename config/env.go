package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var _ = godotenv.Load(".env") // Cargar del archivo llamado ".env"
var (

	//"server=127.0.0.1;uid=root;pwd=12345;database=test"
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
)

const AllowedCORSDomain = "http://localhost"

var envkey map[string]string

// EnvSetup - load env from json

func EnvSetup() {
	var _ = godotenv.Load(".env") // Cargar del archivo llamado ".env"

	/*	jsonFile, err := os.Open("config/env.json")
		defer jsonFile.Close()
		if err != nil {
			panic("failed to open json file: " + err.Error())
		}

		rawJSON, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println(err)
			panic("failed to read json file : " + err.Error())
		}

		err = json.Unmarshal(rawJSON, &envkey)
		if err != nil {
			panic("failed to unmarhsall env.json : " + err.Error())
		}

		for key, value := range envkey {
			os.Setenv(key, value)
		} */
}
