package applicationTable

import (
	"strings"

	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/prompts"
)

func Query(app *structureSpec.App) {
	prompts.RenderTable([][]string{
		{"ID", app.Id},
		{"Name", app.Name},
		{"Description", app.Description},
		{"Tags", strings.Join(app.Tags, ", ")},
	})
}
