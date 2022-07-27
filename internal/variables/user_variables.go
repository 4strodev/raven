package variables

import "os"

var HOME string

func init() {
	HOME = os.Getenv("HOME")
}
