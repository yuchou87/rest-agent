package ai

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/googleai/vertex"
)

const GoogleVertexAIClientName = "googlevertexai"

var _ IAI = (*GoogleVertexAIClient)(nil)

type GoogleVertexAIClient struct {
	llm            llms.Model
	model          string
	embeddingModel string
	temperature    float32
	topP           float32
	topK           int32
	maxTokens      int
}

func (g *GoogleVertexAIClient) Configure(config IAIConfig) error {
	var options []googleai.Option
	if config.GetProviderId() != "" {
		options = append(options, googleai.WithCloudProject(config.GetProviderId()))
	}
	if config.GetProviderRegion() != "" {
		options = append(options, googleai.WithCloudLocation(config.GetProviderRegion()))
	}
	if config.GetPassword() != "" {
		options = append(options, googleai.WithCredentialsJSON([]byte(config.GetPassword())))
	}
	if config.GetModel() != "" {
		options = append(options, googleai.WithDefaultModel(config.GetModel()))
	}

	if config.GetEmbeddingModel() != "" {
		options = append(options, googleai.WithDefaultEmbeddingModel(config.GetEmbeddingModel()))
	}

	llm, err := vertex.New(
		context.Background(),
		options...,
	)

	if err != nil {
		return fmt.Errorf("creating genai Google SDK client: %w", err)
	}

	g.llm = llm
	g.model = config.GetModel()
	g.embeddingModel = config.GetEmbeddingModel()
	g.temperature = config.GetTemperature()
	g.topP = config.GetTopP()
	g.topK = config.GetTopK()
	g.maxTokens = config.GetMaxTokens()
	return nil
}

func (g *GoogleVertexAIClient) GetCompletion(ctx context.Context, messages []llms.MessageContent, callOptions ...llms.CallOption) (*llms.ContentResponse, error) {
	resp, err := g.llm.GenerateContent(ctx, messages, callOptions...)
	return resp, err
}

func (g *GoogleVertexAIClient) GetCompletionFromSinglePrompt(ctx context.Context, prompt string) (string, error) {
	return llms.GenerateFromSinglePrompt(ctx, g.llm, prompt)
}

func (g *GoogleVertexAIClient) GetName() string {
	return GoogleVertexAIClientName
}
