package ai

import (
	"testing"
)

func TestPromptBuilder_GetPromptContent(t *testing.T) {
	p := &PromptBuilder{
		PromptName: "test_case_generation_prompt",
		vars:       "",
	}
	content, err := p.GetPromptContent()
	if err != nil {
		t.Errorf("got a error: %+v", err)
	}
	t.Logf("content: %s", content)
}

func TestPromptBuilder_RenderPromptContent(t *testing.T) {
	p := &PromptBuilder{
		PromptName: "test_case_generation_prompt",
		vars: struct {
			SwaggerFile string
		}{
			SwaggerFile: "test",
		},
	}
	content, err := p.RenderPromptContent()
	if err != nil {
		t.Errorf("got a error: %+v", err)
	}
	t.Logf("content: %s", content)
}

func TestPromptBuilder_BuildPrompt(t *testing.T) {
	p := &PromptBuilder{
		PromptName: "test_case_generation_prompt",
		vars: struct {
			SwaggerFile string
		}{
			SwaggerFile: `{"hello": "world"}`,
		},
	}
	message, err := p.BuildPrompt()
	if err != nil {
		t.Errorf("got a error: %+v", err)
	}
	t.Logf("content: %s", message)
}
