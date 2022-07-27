package variables

import "path"

// return the default templates directory
func TemplatesDirectory() string {
	return path.Join(HOME, ".raven/templates")
}

func TmpTemplatesDirectory() string {
	return "/tmp/raven"
}
