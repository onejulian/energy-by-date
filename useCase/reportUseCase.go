package usecase

import "energyByDate/domain"

type GetReportUseCase struct {}

func (g *GetReportUseCase) Execute(date, period string) (domain.Report, error) {
	report := domain.Report{}
	return report.GetReport(date, period)
}