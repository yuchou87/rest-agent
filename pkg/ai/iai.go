package ai

import (
	"context"
	"github.com/tmc/langchaingo/llms"
)

var (
	Clients = map[string]IAI{
		OpenAIClientName:         &OpenAIClient{},
		GroqAIClientName:         &GroqAIClient{},
		GoogleVertexAIClientName: &GoogleVertexAIClient{},
	}

	Backends = []string{
		OpenAIClientName,
		GroqAIClientName,
		GoogleVertexAIClientName,
	}
)

// IAI is an interface all clients (representing backends) share.
type IAI interface {
	// Configure sets up client for given configuration
	Configure(config IAIConfig) error

	// GetCompletionFromSinglePrompt generates text based on prompt.
	GetCompletionFromSinglePrompt(ctx context.Context, prompt string) (string, error)

	// GetCompletion generates text based on complex prompt.
	GetCompletion(ctx context.Context, messages []llms.MessageContent, callOptions ...llms.CallOption) (*llms.ContentResponse, error)

	// GetName returns name of the backend/client.
	GetName() string
}

type IAIConfig interface {
	GetPassword() string
	GetModel() string
	GetEmbeddingModel() string
	GetBaseURL() string
	GetProxyEndpoint() string
	GetEndpointName() string
	GetEngine() string
	GetTemperature() float32
	GetProviderRegion() string
	GetTopP() float32
	GetTopK() int32
	GetMaxTokens() int
	GetProviderId() string
	GetCompartmentId() string
	GetOrganizationId() string
}

func NewClient(provider string) IAI {
	if client, ok := Clients[provider]; ok {
		return client
	}
	// default client
	return &OpenAIClient{}
}

type Configuration struct {
	Providers       []Provider `mapstructure:"providers" yaml:"providers" goconf:"providers"`
	DefaultProvider string     `mapstructure:"defaultprovider" yaml:"default_provider" goconf:"default_provider"`
}

type Provider struct {
	Name           string  `mapstructure:"name" yaml:"name,omitempty" goconf:"name"`
	Model          string  `mapstructure:"model" yaml:"model,omitempty" goconf:"model"`
	EmbeddingModel string  `mapstructure:"embedding_model" yaml:"embedding_model,omitempty" goconf:"embedding_model"`
	Password       string  `mapstructure:"password" yaml:"password,omitempty" goconf:"password"`
	BaseURL        string  `mapstructure:"base_url" yaml:"base_url,omitempty" goconf:"base_url"`
	ProxyEndpoint  string  `mapstructure:"proxy_endpoint" yaml:"proxy_endpoint,omitempty" goconf:"proxy_endpoint"`
	ProxyPort      string  `mapstructure:"proxy_port" yaml:"proxy_port,omitempty" goconf:"proxy_port"`
	EndpointName   string  `mapstructure:"endpoint_name" yaml:"endpoint_name,omitempty" goconf:"endpoint_name"`
	Engine         string  `mapstructure:"engine" yaml:"engine,omitempty" goconf:"engine"`
	Temperature    float32 `mapstructure:"temperature" yaml:"temperature,omitempty" goconf:"temperature"`
	ProviderRegion string  `mapstructure:"provider_region" yaml:"provider_region,omitempty" goconf:"provider_region"`
	ProviderId     string  `mapstructure:"provider_id" yaml:"provider_id,omitempty" goconf:"provider_id"`
	CompartmentId  string  `mapstructure:"compartment_id" yaml:"compartment_id,omitempty" goconf:"compartment_id"`
	TopP           float32 `mapstructure:"topp" yaml:"topp,omitempty" goconf:"topp"`
	TopK           int32   `mapstructure:"topk" yaml:"topk,omitempty" goconf:"topk"`
	MaxTokens      int     `mapstructure:"max_tokens" yaml:"max_tokens,omitempty" goconf:"max_tokens"`
	OrganizationId string  `mapstructure:"organization_id" yaml:"organization_id,omitempty" goconf:"organization_id"`
}

func (p *Provider) GetBaseURL() string {
	return p.BaseURL
}

func (p *Provider) GetProxyEndpoint() string {
	return p.ProxyEndpoint
}

func (p *Provider) GetEndpointName() string {
	return p.EndpointName
}

func (p *Provider) GetTopP() float32 {
	return p.TopP
}

func (p *Provider) GetTopK() int32 {
	return p.TopK
}

func (p *Provider) GetMaxTokens() int {
	return p.MaxTokens
}

func (p *Provider) GetPassword() string {
	return p.Password
}

func (p *Provider) GetModel() string {
	return p.Model
}

func (p *Provider) GetEmbeddingModel() string {
	return p.EmbeddingModel
}

func (p *Provider) GetEngine() string {
	return p.Engine
}
func (p *Provider) GetTemperature() float32 {
	return p.Temperature
}

func (p *Provider) GetProviderRegion() string {
	return p.ProviderRegion
}

func (p *Provider) GetProviderId() string {
	return p.ProviderId
}

func (p *Provider) GetCompartmentId() string {
	return p.CompartmentId
}

func (p *Provider) GetOrganizationId() string {
	return p.OrganizationId
}
