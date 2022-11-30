package domain

import (
	"energyByDate/dao"
	"errors"
)

var Db dao.DbQueries

func (r *Report) GetReport(date, period string) (Report, error) {
	var reportParsed Report
	var err error

	switch period {
	case "daily":
		report, err := Db.GetDailyReport(date)
		if err != nil {
			return Report{}, err
		}
		reportParsed, err = ParseReport(report)
		if err != nil {
			return reportParsed, err
		}
	case "weekly":
		report, err := Db.GetWeeklyReport(date)
		if err != nil {
			return Report{}, err
		}
		reportParsed, err = ParseReport(report)
		if err != nil {
			return reportParsed, err
		}
	case "monthly":
		report, err := Db.GetMonthlyReport(date)
		if err != nil {
			return Report{}, err
		}
		reportParsed, err = ParseReport(report)
		if err != nil {
			return reportParsed, err
		}
	default:
		return Report{}, errors.New("period not found")
	}

	return reportParsed, err
}

func ParseReport(report dao.Report) (Report, error) {
	var reportParsed Report
	for _, row := range report.Rows {
		reportParsed.Rows = append(reportParsed.Rows, Row{
			MeterDate: row.MeterDate,
			Value:     row.Value,
		})
	}
	return reportParsed, nil
}
