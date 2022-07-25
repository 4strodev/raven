package fs

import (
	"os"

	"github.com/spf13/afero"
)

var fs = afero.NewOsFs()

const (
	_DIR_MODE = os.FileMode(0755)
)

func MkdirAll(path string) error {
	return fs.MkdirAll(path, _DIR_MODE)
}
