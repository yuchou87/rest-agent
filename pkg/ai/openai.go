package ai

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

const OpenAIClientName = "openai"

var _ IAI = (*OpenAIClient)(nil)

type OpenAIClient struct {
	llm            llms.Model
	model          string
	embeddingModel string
	temperature    float32
	topP           float32
	organizationId string
}

func (c *OpenAIClient) Configure(config IAIConfig) error {
	var options []openai.Option
	if config.GetPassword() != "" {
		options = append(options, openai.WithToken(config.GetPassword()))
	}
	if config.GetBaseURL() != "" {
		options = append(options, openai.WithBaseURL(config.GetBaseURL()))
	}
	if config.GetOrganizationId() != "" {
		options = append(options, openai.WithOrganization(config.GetOrganizationId()))
	}

	if config.GetModel() != "" {
		options = append(options, openai.WithModel(config.GetModel()))
	}

	if config.GetEmbeddingModel() != "" {
		options = append(options, openai.WithEmbeddingModel(config.GetEmbeddingModel()))
	}

	llm, err := openai.New(options...)
	if err != nil {
		return fmt.Errorf("creating OpenAi client: %w", err)
	}

	c.llm = llm
	c.model = config.GetModel()
	c.embeddingModel = config.GetEmbeddingModel()
	c.temperature = config.GetTemperature()
	c.topP = config.GetTopP()
	c.organizationId = config.GetOrganizationId()

	return nil
}

func (c *OpenAIClient) GetCompletion(ctx context.Context, messages []llms.MessageContent, callOptions ...llms.CallOption) (*llms.ContentResponse, error) {
	resp, err := c.llm.GenerateContent(ctx, messages, callOptions...)
	return resp, err
}

func (c *OpenAIClient) GetCompletionFromSinglePrompt(ctx context.Context, prompt string) (string, error) {
	return llms.GenerateFromSinglePrompt(ctx, c.llm, prompt)
}

func (c *OpenAIClient) GetName() string {
	return OpenAIClientName
}
