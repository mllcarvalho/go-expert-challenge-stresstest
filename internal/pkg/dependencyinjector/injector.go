package dependencyinjector

import (
	"github.com/mllcarvalho/go-expert-challenge-stresstest/internal/infra/cli"
	"github.com/mllcarvalho/go-expert-challenge-stresstest/internal/infra/cli/commands"
	"github.com/mllcarvalho/go-expert-challenge-stresstest/internal/pkg/httpclient"
	"github.com/mllcarvalho/go-expert-challenge-stresstest/internal/usecases/report"
	"github.com/mllcarvalho/go-expert-challenge-stresstest/internal/usecases/stress"
)

type DependencyInjectorInterface interface {
	Inject() (*Dependencies, error)
}

type DependencyInjector struct{}

type Dependencies struct {
	CLI cli.CLIInterface
}

func NewDependencyInjector() *DependencyInjector {
	return &DependencyInjector{}
}

func (d *DependencyInjector) Inject() (*Dependencies, error) {
	httpClient := httpclient.NewHttpClient()
	stressTestUseCase := stress.NewStressTestUseCase(httpClient)
	reportUseCase := report.NewReportUseCase()
	stressTestCmd := commands.NewStressTestCmd(stressTestUseCase, reportUseCase)

	cli := cli.NewCLI(stressTestCmd.Build())

	return &Dependencies{
		CLI: cli,
	}, nil
}
