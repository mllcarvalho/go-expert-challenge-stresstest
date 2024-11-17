package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/mllcarvalho/go-expert-challenge-stresstest/internal/usecases/report/dto"
)

type ReportUseCaseMock struct {
	mock.Mock
}

func (m *ReportUseCaseMock) Execute(input dto.ReportInput) *dto.ReportOutput {
	args := m.Called(input)
	return args.Get(0).(*dto.ReportOutput)
}
