package jobs

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/taubyte/go-interfaces/services/patrick"
	schemaCommon "github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/tau-cli/cli/common"
	"github.com/taubyte/tau-cli/cli/common/options"
	"github.com/taubyte/tau-cli/env"
	"github.com/taubyte/tau-cli/i18n"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/taubyte/tau-cli/prompts"
	authClient "github.com/taubyte/tau-cli/singletons/auth_client"
	patrickClient "github.com/taubyte/tau-cli/singletons/patrick_client"
	jobsTable "github.com/taubyte/tau-cli/table/jobs"
	"github.com/taubyte/tau-cli/validate"
	"github.com/urfave/cli/v2"
)

type link struct {
	common.UnimplementedBasic
}

func New() common.Basic {
	return link{}
}

func (link) Base() (*cli.Command, []common.Option) {
	selected, err := env.GetSelectedProject()
	if err != nil {
		selected = "selected"
	}

	return common.Base(&cli.Command{
		Name:      "jobs",
		ArgsUsage: i18n.ArgsUsageName,
	}, options.NameFlagSelectedArg0(selected))
}

const defaultTimeFilter = "10m"

func (link) Query() common.Command {
	return common.Create(
		&cli.Command{
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "time",
					Aliases:     []string{"t"},
					Usage:       "filters jobs by time range",
					DefaultText: defaultTimeFilter,
				},
				&cli.BoolFlag{
					Name:    "defaulttime",
					Aliases: []string{"dt"},
					Usage:   "use default time filter (10m)",
				},
			},
			Action: query,
		},
	)
}

func query(ctx *cli.Context) error {
	prj, err := projectLib.SelectedProjectInterface()
	if err != nil {
		return err
	}

	patrickC, err := patrickClient.Load()
	if err != nil {
		return err
	}

	authC, err := authClient.Load()
	if err != nil {
		return err
	}

	jobIds, err := patrickC.Jobs(prj.Get().Id())
	if err != nil {
		// use i18n
		return err
	}

	timeRange := defaultTimeFilter
	if !ctx.Bool("defaulttime") {
		timeRange = prompts.GetOrRequireAString(ctx, "time", "Job time range", validate.Time, defaultTimeFilter)
	}

	tRange, err := schemaCommon.StringToTime(timeRange)
	if err != nil {
		return err
	}
	tRangeNano := time.Duration(tRange) * time.Nanosecond
	rangeEnd := time.Now().Unix() - int64(tRangeNano.Seconds())

	jobMap := make(map[int64]*patrick.Job, len(jobIds))
	for _, id := range jobIds {
		job, err := patrickC.Job(id)
		if err != nil {
			// use i18n
			return err
		}

		if job.Timestamp >= rangeEnd {
			jobMap[job.Timestamp] = job
		}
	}

	// separate keys from original for loop to ensure unique values
	keys := make([]int64, 0, len(jobMap))
	for k := range jobMap {
		keys = append(keys, k)
	}

	t, err := jobsTable.ListNoRender(authC, jobMap, keys)
	if err != nil {
		return err
	}

	t.SetStyle(table.StyleLight)
	t.Render()

	return nil
}
