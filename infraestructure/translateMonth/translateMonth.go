package translatemonth

import (
	"bytes"
	"encoding/json"
	"energyByDate/env" // for local
	"os"
)

func TranslateEnToSp(month string) (string, error) {
	translateMonth := TranslateMonth{}
	months, err := os.ReadFile(env.RootDir()+"/infraestructure/assets/monthsEnToSp.json") // for local
	// months, err := os.ReadFile("monthsEnToSp.json") // for docker
	if err != nil {
		return "", err
	}
	decoder := json.NewDecoder(bytes.NewReader(months))

	err = decoder.Decode(&translateMonth)
	if err != nil {
		return "", err
	}

	monthEs := translateMonth.Months.(map[string]interface{})[month].(string)
	return monthEs, nil
}

func TranslateSpToEn(month string) (string, error) {
	translateMonth := TranslateMonth{}
	months, err := os.ReadFile(env.RootDir()+"/infraestructure/assets/monthsSpToEn.json") // for local
	// months, err := os.ReadFile("monthsSpToEn.json") // for docker
	if err != nil {
		return "", err
	}
	decoder := json.NewDecoder(bytes.NewReader(months))

	err = decoder.Decode(&translateMonth)
	if err != nil {
		return "", err
	}

	monthEs := translateMonth.Months.(map[string]interface{})[month].(string)
	return monthEs, nil
}
