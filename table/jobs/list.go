package jobsTable

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/taubyte/go-interfaces/services/patrick"
	authHttp "github.com/taubyte/tau/clients/http/auth"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ListNoRender(authClient *authHttp.Client, jobMap map[int64]*patrick.Job, keys []int64) (table.Writer, error) {
	keysIface := int64arr(keys)
	sort.Sort(keysIface)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAllowedRowLength(79)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Date"},
		{Name: "Time"},
		{Name: "Type"},
		{Name: "Status"},
		{Name: "Commit"},
	})
	t.AppendHeader(table.Row{"Date", "Time", "Type", "Status", "Commit"})

	timeZone, _ := time.LoadLocation("Local")
	for _, key := range keysIface {
		job := jobMap[key]
		unixTime := time.Unix(job.Timestamp, 0).In(timeZone)
		date := unixTime.Format("01/02/2006")
		time := unixTime.Format("3:04:05 PM")

		repo, err := authClient.GetRepositoryById(fmt.Sprintf("%d", job.Meta.Repository.ID))
		if err != nil {
			return nil, err
		}

		repoType := "Unknown"
		name := repo.Get().Name()
		nameSplit := strings.SplitN(name, "_", 3)
		if nameSplit[0] == "tb" {
			switch nameSplit[1] {
			case "library", "website", "code":
				repoType = cases.Title(language.English).String(nameSplit[1])
			default:
				repoType = "Config"
			}
		}
		t.AppendRow(table.Row{
			date,
			time,
			repoType,
			job.Status.String(),
			job.Meta.HeadCommit.ID,
		})

		t.AppendSeparator()
	}

	return t, nil
}

type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] > a[j] }
func (a int64arr) String() (s string) {
	sep := "" // for printing separating commas
	for _, el := range a {
		s += sep
		sep = ", "
		s += fmt.Sprintf("%d", el)
	}
	return
}
