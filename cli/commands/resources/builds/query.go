package builds

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/taubyte/go-interfaces/services/patrick"
	"github.com/taubyte/tau-cli/cli/common"
	projectLib "github.com/taubyte/tau-cli/lib/project"
	authClient "github.com/taubyte/tau-cli/singletons/auth_client"
	patrickClient "github.com/taubyte/tau-cli/singletons/patrick_client"
	buildsTable "github.com/taubyte/tau-cli/table/builds"
	"github.com/urfave/cli/v2"
)

func queryOrList() common.Command {
	return common.Create(
		&cli.Command{
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "since",
					Aliases:     []string{"t", "s"},
					Usage:       "(optional) filters jobs by time range",
					DefaultText: defaultTimeFilter,
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

	since := defaultTimeFilter
	if _since := ctx.String("since"); len(_since) > 0 {
		since = _since
	}

	sinceParsed, err := time.ParseDuration(since)
	if err != nil {
		return err
	}

	rangeEnd := time.Now().Add(-sinceParsed).Unix()

	// index string for unique jobs, int64 to order by time
	jobMap := make(map[string]map[int64]*patrick.Job, len(jobIds))
	for _, id := range jobIds {
		job, err := patrickC.Job(id)
		if err != nil {
			// use i18n
			return err
		}

		if job.Timestamp >= rangeEnd {
			jobMap[id] = make(map[int64]*patrick.Job, 1)
			jobMap[id][job.Timestamp] = job
		}
	}

	// separate keys from original for loop to ensure unique values
	keys := make([]int64, 0, len(jobIds))
	for _, v := range jobMap {
		for key := range v {
			keys = append(keys, key)
		}
	}

	t, err := buildsTable.ListNoRender(authC, jobMap, keys, false)
	if err != nil {
		return err
	}

	t.SetStyle(table.StyleLight)
	t.Render()

	return nil
}
