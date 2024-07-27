package prompts

import (
	"testing"
)

func TestLoadPrompts(t *testing.T) {
	prompts, _ := LoadPrompts()
	t.Logf("prompts: %+v", prompts)
	t.Logf("content: %s", prompts["test_case_generation_prompt"])
}
