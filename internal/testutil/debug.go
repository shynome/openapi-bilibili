package testutil

import (
	"os"
	"path/filepath"
	"strings"
)

var Debug = false

func init() {
	name := filepath.Base(os.Args[0])
	Debug = strings.HasPrefix(name, "__debug_bin")
}
