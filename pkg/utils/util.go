package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
)

// CovertToJson converts any type to json string
func CovertToJson(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// CalculateMd5 calculates the md5 hash of a string
func CalculateMd5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

// FileExists checks if the file exists
func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}

// EnsureDirExists ensures that the directory exists
func EnsureDirExists(dir string) error {
	err := os.MkdirAll(dir, 0o755)

	if errors.Is(err, os.ErrExist) {
		return nil
	}

	return err
}
