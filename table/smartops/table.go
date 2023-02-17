package smartopsTable

import (
	"strings"

	commonSchema "github.com/taubyte/go-project-schema/common"
	structureSpec "github.com/taubyte/go-specs/structure"
)

func getTableData(smartops *structureSpec.SmartOp, showId bool) (toRender [][]string) {
	if showId == true {
		toRender = [][]string{
			{"ID", smartops.Id},
		}
	}

	toRender = append(toRender, [][]string{
		{"Name", smartops.Name},
		{"Description", smartops.Description},
		{"Tags", strings.Join(smartops.Tags, ", ")},
		{"Timeout", commonSchema.TimeToString(smartops.Timeout)},
		{"Memory", commonSchema.UnitsToString(smartops.Memory)},
		{"Source", smartops.Source},
		{"Call", smartops.Call},
	}...)

	return toRender
}
