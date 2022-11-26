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
	monthTraslated, err := translatemonth.Translate(month)
	if err != nil {
		return "", err
	}
	date = strings.Replace(date, month, monthTraslated, 1)
	return date, nil
}