package serviceTable

import (
	"strings"

	structureSpec "github.com/taubyte/go-specs/structure"
	"github.com/taubyte/tau/prompts"
)

func Query(service *structureSpec.Service) {
	prompts.RenderTable([][]string{
		{"ID", service.Id},
		{"Name", service.Name},
		{"Description", service.Description},
		{"Tags", strings.Join(service.Tags, ", ")},
		{"Protocol", service.Protocol},
	})
}
