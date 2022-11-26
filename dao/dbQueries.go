package dao

import (
	"database/sql"
	"energyByDate/infraestructure/formatDate"
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

	dateFormated, err := formatDate.Format(date)
	if err != nil {
		return report, err
	}

	rows, err := conn.Query("select active_energy, meter_date from report where meter_date like $1", dateFormated+"%")
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

	report, err = db.ProccesReportDaily(report)
	if err != nil {
		return report, err
	}

	return report, nil
}

func (db *DbQueries) GetReportWeekly(date string) (Report, error) {
	report := Report{}
	return report, nil
}

func (db *DbQueries) GetReportMonthly(date string) (Report, error) {
	report := Report{}
	return report, nil
}

func (db *DbQueries) ProccesReportDaily(report Report) (Report, error) {
	//type code here
	return report, nil
}

