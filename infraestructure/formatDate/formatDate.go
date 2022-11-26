package formatDate

import (
	translatemonth "energyByDate/infraestructure/translateMonth"
	"strings"
	"time"
)

func Format(date string) (string, error) {
	dateFormated, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}
	date = dateFormated.Format("January 02, 2006")
	dateSplited := strings.Split(date, " ")
	month := dateSplited[0]
	monthTraslated, err := translatemonth.TranslateEnToSp(month)
	if err != nil {
		return "", err
	}
	date = strings.Replace(date, month, monthTraslated, 1)
	return date, nil
}

func FormatToEn(date string) (time.Time, error) {
	var result time.Time
	dateSplited := strings.Split(date, " ")
	month := dateSplited[0]
	monthTraslated, err := translatemonth.TranslateSpToEn(month)
	if err != nil {
		return result, err
	}
	date = strings.Replace(date, month, monthTraslated, 1)
	result, err = time.Parse("January 02, 2006, 3:04 PM", date)
	if err != nil {
		return result, err
	}
	return result, nil
}