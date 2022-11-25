package usecase

import "energyByDate/domain"

func GetReportUseCase(date, period string) (domain.Report, error) {
	report := domain.Report{}
	return report.GetReport(date, period)
}