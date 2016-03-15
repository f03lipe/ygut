package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type config struct {
	Production bool   `json:"PROD"`
	Env        string `json:"ENV"`
	CsrfKey32  string `json:"CSRF_KEY_32"`
}

var C *config

// Perhaps the best would be for the application not to rely
// on env variables. But that's not always viable.
func UpdateEnv(c *config) {
	v := reflect.ValueOf(c).Elem()
	for i := 0; i < v.NumField(); i += 1 {
		key := v.Type().Field(i).Tag.Get("json")
		if key == "" {
			key = v.Type().Field(i).Name
		}
		os.Setenv(key, v.Field(i).String())
	}
}

func ReadFromJSON(path string, c *config) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	if err := json.Unmarshal(body, &c); err != nil {
		panic("Failed to read config JSON.\n" + err.Error())
	}
}

func Setup() {
	C = &config{
		//Production: false,
		Env:       "production",
		CsrfKey32: "writearandom32bytestringforcsrf.",
	}

	if _, err := os.Stat("./env.json"); !os.IsNotExist(err) {
		ReadFromJSON("./env.json", C)
	}

	UpdateEnv(C)
}
