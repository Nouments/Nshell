package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func SearchBinary(bin string) (string, error) {
	for _, dir := range strings.Split(os.Getenv("PATH"), ":") {
		fullpath := filepath.Join(dir, bin)
		info, err := os.Stat(fullpath)
		if err != nil {
			continue
		}
		if info.IsDir() {
			continue
		}
		if info.Mode().Perm()&0111 != 0 {
			return fullpath, nil
		}
	}
	return "", fmt.Errorf("%s: not found", bin)
}
