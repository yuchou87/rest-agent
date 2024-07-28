package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"time"
)

// CovertToJson converts any type to json string
func CovertToJson(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// CovertToJsonWithIndent converts any type to json string with indent
func CovertToJsonWithIndent(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

// CoverToYaml converts any type to yaml string
func CoverToYaml(v any) string {
	b, _ := yaml.Marshal(v)
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

// GetConfigDir returns the config directory
func GetConfigDir() string {
	return filepath.Join(xdg.ConfigHome, "rest-agent")
}

// GetDateTime returns the current date and time
func GetDateTime() string {
	return time.Now().Format("2006-01-02_15:04:05")
}
