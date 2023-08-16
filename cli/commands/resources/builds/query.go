package builds

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/taubyte/go-interfaces/services/patrick"
	schemaCommon "github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/tau-cli/cli/common"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	"github.com/taubyte/tau-cli/prompts"
	authClient "github.com/taubyte/tau-cli/singletons/auth_client"
	patrickClient "github.com/taubyte/tau-cli/singletons/patrick_client"
	buildsTable "github.com/taubyte/tau-cli/table/builds"
	"github.com/taubyte/tau-cli/validate"
	"github.com/urfave/cli/v2"
)

func queryOrList() common.Command {
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

func (link) Query() common.Command {
	return queryOrList()
}

func (link) List() common.Command {
	return queryOrList()
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

	t, err := buildsTable.ListNoRender(authC, jobMap, keys)
	if err != nil {
		return err
	}

	t.SetStyle(table.StyleLight)
	t.Render()

	return nil
}
