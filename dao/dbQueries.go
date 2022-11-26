package dao

import (
	"database/sql"
	"energyByDate/infraestructure/formatDate"
	"strconv"
	"strings"
	"time"
)

type DbQueries struct {
}

func (db *DbQueries) GetConn() (*sql.DB, error) {
	conn, err := ConnectPostgres()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (db *DbQueries) GetReportDaily(date string) (Report, error) {
	report := Report{}
	row := Row{}
	
	conn, err := db.GetConn()
	if err != nil {
		return report, err
	}
	defer conn.Close()

	yesterday, err := db.GetYesterday(date)
	if err != nil {
		return report, err
	}
	yesterday = yesterday + ",%"

	today, err := db.GetToday(date)
	if err != nil {
		return report, err
	}
	today = today + ",%"

	rows, err := conn.Query("select active_energy, meter_date from report where meter_date between $1 and $2", yesterday, today)
	if err != nil {
		return report, err
	}

	for rows.Next() {
		err = rows.Scan(&row.Value, &row.MeterDate)
		if err != nil {
			return report, err
		}
		report.Rows = append(report.Rows, row)
	}

	report, err = db.ProccesReportDaily(report, date)
	if err != nil {
		return report, err
	}

	return report, nil
}

func (db *DbQueries) GetReportWeekly(date string) (Report, error) {
	var reportWeekly Report

	day, err := db.GetDateMonday(date)
	if err != nil {
		return Report{}, err
	}

	for i := 0; i < 7; i++ {
		var spent float64
		var row Row
		report, err := db.GetReportDaily(day.Format("2006-01-02"))
		if err != nil {
			return report, err
		}
		for _, row := range report.Rows {
			valueParsed, err := strconv.ParseFloat(row.Value, 64)
			if err != nil {
				return reportWeekly, err
			}
			spent = spent + valueParsed
		}
		row.Value = strconv.FormatFloat(spent, 'f', 2, 64)
		row.MeterDate = day.Format("2006-01-02")+" 00:00:00"
		reportWeekly.Rows = append(reportWeekly.Rows, row)
		day = day.AddDate(0, 0, 1)
	}

	return reportWeekly, nil
}

func (db *DbQueries) GetReportMonthly(date string) (Report, error) {
	var reportMonthly Report

	day, err := db.GetFirtsDayMonth(date)
	if err != nil {
		return reportMonthly, err
	}

	daysMonth := day.AddDate(0, 1, -1).Day()

	for i := 0; i < daysMonth; i++ {
		var spent float64
		var row Row
		report, err := db.GetReportDaily(day.Format("2006-01-02"))
		if err != nil {
			return reportMonthly, err
		}
		for _, row := range report.Rows {
			valueParsed, err := strconv.ParseFloat(row.Value, 64)
			if err != nil {
				return reportMonthly, err
			}
			spent = spent + valueParsed
		}
		row.Value = strconv.FormatFloat(spent, 'f', 2, 64)
		row.MeterDate = day.Format("2006-01-02")+" 00:00:00"
		reportMonthly.Rows = append(reportMonthly.Rows, row)
		day = day.AddDate(0, 0, 1)
	}
	return reportMonthly, nil
}

func (db *DbQueries) ProccesReportDaily(report Report, date string) (Report, error) {
	var reportFinal Report

	dateParsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return reportFinal, err
	}
	today := dateParsed.Format("2006-01-02")
	today = today + " 00:00:00"

	todayParsed, err := time.Parse("2006-01-02 15:04:05", today)
	if err != nil {
		return reportFinal, err
	}

	for i := -1; i < 23; i++ {
		var spent float64
		row := Row{}
		start := todayParsed.Add(time.Hour * time.Duration(i))
		end := todayParsed.Add(time.Hour * time.Duration(i+1))
		for _, rowReport := range report.Rows {
			period, err := formatDate.FormatToEn(rowReport.MeterDate)
			if err != nil {
				return reportFinal, err
			}
			if period.After(start) && period.Before(end) {
				value := strings.ReplaceAll(rowReport.Value, ".", "")
				value = strings.ReplaceAll(value, ",", ".")
				valueParsed, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return reportFinal, err
				}
				spent = spent + valueParsed
			}
		}
		row.Value = strconv.FormatFloat(spent, 'f', 2, 64)
		row.MeterDate = end.Format("2006-01-02 15:04:05")
		reportFinal.Rows = append(reportFinal.Rows, row)
	}
	return reportFinal, nil
}

func (db *DbQueries) GetYesterday(date string) (string, error) {
	dateParsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}
	date = dateParsed.AddDate(0, 0, -2).Format("2006-01-02")
	date, err = formatDate.Format(date)
	if err != nil {
		return "", err
	}
	return date, nil
}

func (db *DbQueries) GetTomorrow(date string) (string, error) {
	dateParsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}
	date = dateParsed.AddDate(0, 0, 1).Format("2006-01-02")
	date, err = formatDate.Format(date)
	if err != nil {
		return "", err
	}
	return date, nil
}

func (db *DbQueries) GetToday(date string) (string, error) {
	dateParsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}
	date = dateParsed.Format("2006-01-02")
	date, err = formatDate.Format(date)
	if err != nil {
		return "", err
	}
	return date, nil
}

func (db *DbQueries) GetDateMonday(date string) (time.Time, error) {
	var dateResult time.Time

	dateParsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return dateResult, err
	}
	date = dateParsed.Format(time.RFC850)
	for !strings.Contains(date, "Monday") {
		dateParsed = dateParsed.AddDate(0, 0, -1)
		date = dateParsed.Format(time.RFC850)
	}
	dateResult, err = time.Parse(time.RFC850, date)
	if err != nil {
		return dateResult, err
	}

	return dateResult, nil
}

func (db *DbQueries) GetFirtsDayMonth(date string) (time.Time, error) {
	var firstDayMonth time.Time

	dateSplited := strings.Split(date, "-")
	day := dateSplited[2]
	
	date = strings.ReplaceAll(date, day, "01")
	firstDayMonth, err := time.Parse("2006-01-02", date)
	if err != nil {
		return firstDayMonth, err
	}

	return firstDayMonth, nil
}
