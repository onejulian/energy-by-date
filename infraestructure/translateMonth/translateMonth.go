package translatemonth

import (
	"bytes"
	"encoding/json"
	"energyByDate/env"
	"os"
)

func TranslateEnToSp(month string) (string, error) {
	translateMonth := TranslateMonth{}
	months, err := os.ReadFile(env.RootDir()+"/infraestructure/assets/monthsEnToSp.json")
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
	months, err := os.ReadFile(env.RootDir()+"/infraestructure/assets/monthsSpToEn.json")
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
