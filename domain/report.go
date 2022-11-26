package domain

import (
	"energyByDate/dao"
)

var Db dao.DbQueries

func (r *Report) GetReport(date, period string) (Report, error) {
	var report Report

	switch period {
	case "daily":
		report, err := Db.GetReportDaily(date)
		if err != nil {
			return Report{}, err
		}
		reportParsed, err := ParseReport(report)
		if err != nil {
			return reportParsed, err
		}
	}

	return report, nil
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