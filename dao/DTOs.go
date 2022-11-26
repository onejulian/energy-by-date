package dao

type Report struct {
	Rows []Row `json:"rows"`
}

type Row struct {
	MeterDate string `json:"meter_date"`
	Value     string `json:"value"`
}
