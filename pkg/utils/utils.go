package utils

import (
	"os"
	"path/filepath"
)

var cacheDir string

func CacheDir() string {
	if cacheDir == "" {
		var err error
		cacheDir, err = os.UserCacheDir()
		if err != nil {
			cacheDir = os.TempDir()
		}
	}
	dir := filepath.Join(cacheDir, "go-caniuse")
	return dir
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
