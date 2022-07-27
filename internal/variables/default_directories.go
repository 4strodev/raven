package variables

import "path"

func TemplatesDirectory() string {
	return path.Join(HOME, ".raven/templates")
}

func TmpTemplatesDirectory() string {
	return "/tmp/raven"
}
