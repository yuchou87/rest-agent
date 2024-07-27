package ai

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/tmc/langchaingo/llms"
	"github.com/yuchou87/rest-agent/pkg/ai/prompts"
	"strings"
	"text/template"
)

type Prompt struct {
	System string
	User   string
}

type PromptBuilder struct {
	PromptName string
	vars       any
}

func NewPromptBuilder(promptName string, vars any) *PromptBuilder {
	return &PromptBuilder{
		PromptName: promptName,
		vars:       vars,
	}
}

func (p *PromptBuilder) GetPromptContent() (string, error) {
	promptFiles, err := prompts.LoadPrompts()
	if err != nil {
		return "", err
	}
	return promptFiles[p.PromptName], nil
}

func (p *PromptBuilder) RenderPromptContent() (string, error) {
	var sb strings.Builder
	promptContent, err := p.GetPromptContent()
	if err != nil {
		return "", nil
	}
	if promptContent == "" {
		return "", errors.New("no prompt content")
	}
	tmpl := template.New("prompt")
	template.Must(
		tmpl.Funcs(
			template.FuncMap{
				"trim": strings.TrimSpace,
			}).
			Parse(promptContent),
	)
	if err := tmpl.Execute(&sb, p.vars); err != nil {
		return "", err
	}

	return sb.String(), nil
}

func (p *PromptBuilder) BuildPrompt() ([]llms.MessageContent, error) {
	promptContent, err := p.RenderPromptContent()
	if err != nil {
		return nil, err
	}

	var prompt Prompt
	if err := toml.Unmarshal([]byte(promptContent), &prompt); err != nil {
		return nil, err

	}

	message := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, prompt.System),
		llms.TextParts(llms.ChatMessageTypeHuman, prompt.User),
	}

	return message, nil
}
