package prompts

import (
	"embed"
	"strings"
)

var (
	//go:embed *.toml
	promptFiles embed.FS
)

// LoadPrompts loads the prompts from the embedded files
func LoadPrompts() (map[string]string, error) {
	files, err := promptFiles.ReadDir(".")
	if err != nil {
		return nil, err
	}

	prompts := make(map[string]string)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		prompt, err := promptFiles.ReadFile(file.Name())
		if err != nil {
			return nil, err
		}

		prompts[strings.TrimSuffix(file.Name(), ".toml")] = string(prompt)
	}

	return prompts, nil
}
