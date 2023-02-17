package codefile

import "path"

func init() {
	getTemplateCommon = func(split []string) string {
		return path.Join(path.Join(split[0:len(split)-1]...), "common")
	}
}
