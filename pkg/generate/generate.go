package generate

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/llms"
	"github.com/yuchou87/rest-agent/pkg/ai"
	"github.com/yuchou87/rest-agent/pkg/utils"
	"os"
)

type Generate struct {
	Context     context.Context
	AIClient    ai.IAI
	AIProvider  string
	PromptName  string
	PromptVars  any
	OutputType  string
	Output      any
	Results     string
	RunningTime string
}

func NewGenerate(ctx context.Context, promptName string, promptVars any, outputType string, output any) (*Generate, error) {
	var aiConfig ai.Configuration
	if err := viper.UnmarshalKey("ai", &aiConfig); err != nil {
		return nil, err
	}

	backend := aiConfig.DefaultProvider
	if backend == "" {
		return nil, errors.New("default AI provider not specified in configuration. Please check")
	}

	if len(aiConfig.Providers) == 0 {
		return nil, errors.New("AI provider not specified in configuration. Please check")
	}

	var aiProvider ai.Provider
	for _, provider := range aiConfig.Providers {
		if backend == provider.Name {
			aiProvider = provider
			break
		}
	}

	aiClient := ai.NewClient(backend)
	if err := aiClient.Configure(&aiProvider); err != nil {
		return nil, err
	}

	gen := &Generate{
		Context:     ctx,
		AIClient:    aiClient,
		AIProvider:  aiProvider.Name,
		PromptName:  promptName,
		PromptVars:  promptVars,
		OutputType:  outputType,
		Output:      output,
		RunningTime: utils.GetDateTime(),
	}

	return gen, nil
}

// Generate generates the text based on the prompt.
func (g *Generate) Generate() error {
	prompt, err := g.BuildPrompt()
	if err != nil {
		return err
	}

	response, err := g.AIClient.GetCompletion(g.Context, prompt)
	if err != nil {
		return err
	}

	g.Results = response.Choices[0].Content

	result := ai.NewStructuredResponse(g.Results, g.OutputType, g.Output)
	if err := result.Parse(); err != nil {
		return err
	}

	if err := g.SaveAndPrintResults(); err != nil {
		return err
	}
	return nil
}

// BuildPrompt builds the prompt based on the prompt name and variables.
func (g *Generate) BuildPrompt() ([]llms.MessageContent, error) {
	prompt := ai.NewPromptBuilder(g.PromptName, g.PromptVars)
	return prompt.BuildPrompt()
}

// SaveAndPrintResults saves the results and errors to the files and prints the errors.
func (g *Generate) SaveAndPrintResults() error {
	return g.SaveFile(g.GetOutputFile(), []byte(g.GetOutputResults()))
}

// SaveFile saves the data to the specified file.
func (g *Generate) SaveFile(filename string, data []byte) error {
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return err
	}
	return nil
}

// GetFileSuffix returns the file suffix based on the output format.
func (g *Generate) GetFileSuffix() string {
	switch g.OutputType {
	case ai.StructuredResponseCodeTypeYAML:
		return "yaml"
	default:
		return "json"
	}
}

// GenerateFileName generates the file name based on the prompt variables.
func (g *Generate) GenerateFileName() string {
	return utils.CalculateMd5(utils.CovertToJson(g.PromptVars))
}

// GetOutputFile returns the output file path.
func (g *Generate) GetOutputFile() string {
	return fmt.Sprintf("%s/%s_%s.%s", utils.GetConfigDir(), g.GenerateFileName(), g.RunningTime, g.GetFileSuffix())
}

// GetOutputResults returns the results.
func (g *Generate) GetOutputResults() string {
	switch g.OutputType {
	case ai.StructuredResponseCodeTypeYAML:
		return utils.CoverToYaml(g.Output)
	default:
		return utils.CovertToJsonWithIndent(g.Output)
	}
}
