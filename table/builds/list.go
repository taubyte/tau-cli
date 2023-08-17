package buildsTable

import (
	"os"
	"sort"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/taubyte/go-interfaces/services/patrick"
	authHttp "github.com/taubyte/tau/clients/http/auth"
)

func ListNoRender(authClient *authHttp.Client, jobMap map[string]map[int64]*patrick.Job, keys []int64, showCommit bool) (table.Writer, error) {
	keysIface := int64arr(keys)
	sort.Sort(keysIface)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAllowedRowLength(79)
	lastColumn := "Job ID"
	if showCommit {
		lastColumn = "Commit"
	}

	t.SetColumnConfigs([]table.ColumnConfig{
		{Align: text.AlignCenter},
		{Name: "Time"},
		{Name: "Type"},
		{Name: lastColumn},
	})

	t.AppendHeader(table.Row{"", "Time", "Type", lastColumn})

	timeZone, _ := time.LoadLocation("Local")
	for _, key := range keysIface {
		for _, timeMap := range jobMap {
			if job, ok := timeMap[key]; ok {

				row, err := row(authClient, job, timeZone, showCommit)
				if err != nil {
					return nil, err
				}

				t.AppendRow(row)
				t.AppendSeparator()
			}
		}
	}

	return t, nil
}
