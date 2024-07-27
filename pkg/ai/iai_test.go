package ai

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"os"
	"testing"
)

var (
	GoogleProvider = Provider{
		Name:           "googlevertexai",
		Model:          "gemini-1.5-pro",
		EmbeddingModel: "",
		Password:       os.Getenv("GOOGLE_API_KEY"),
		ProviderId:     "yuchou",
		ProviderRegion: "asia-southeast1",
	}
	OpenaiProvider = Provider{
		Name:           "openai",
		Model:          "gpt-3.5-turbo",
		EmbeddingModel: "",
		Password:       os.Getenv("OPENAI_API_KEY"),
		BaseURL:        "https://openrouter.ai/api/v1",
	}
	GroqProvider = Provider{
		Name:           "grop",
		Model:          "llama3-8b-8192",
		EmbeddingModel: "",
		Password:       os.Getenv("GROQ_API_KEY"),
		BaseURL:        "https://api.groq.com/openai/v1",
	}
	AIConfigurationData = Configuration{
		Providers:       []Provider{GoogleProvider, OpenaiProvider, GroqProvider},
		DefaultProvider: "openai",
	}
)

func Test_GetCompletionFromSinglePrompt(t *testing.T) {
	type args struct {
		provider string
		config   IAIConfig
		prompt   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "OpenAi client test",
			args: args{
				provider: OpenAIClientName,
				config:   &OpenaiProvider,
				prompt:   "Who was the second person to walk on the moon?",
			},
			want: true,
		},
		{
			name: "Groq client test",
			args: args{
				provider: GroqAIClientName,
				config:   &GroqProvider,
				prompt:   "Who was the second person to walk on the moon?",
			},
			want: true,
		}, {
			name: "Google vertex ai client test",
			args: args{
				provider: GoogleVertexAIClientName,
				config:   &GoogleProvider,
				prompt:   "世界上最高的山峰叫什么?",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.args.provider)
			_ = client.Configure(tt.args.config)
			resp, _ := client.GetCompletionFromSinglePrompt(context.Background(), tt.args.prompt)
			t.Logf(resp)
		})
	}
}

func Test_GetCompletion(t *testing.T) {
	type args struct {
		provider string
		config   IAIConfig
		messages []llms.MessageContent
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "OpenAi client test",
			args: args{
				provider: OpenAIClientName,
				config:   &OpenaiProvider,
				messages: []llms.MessageContent{
					llms.TextParts(llms.ChatMessageTypeSystem, "You are a helpful AI assistant."),
					llms.TextParts(llms.ChatMessageTypeHuman, "Who was the second person to walk on the moon?"),
				},
			},
			want: true,
		},
		{
			name: "Groq client test",
			args: args{
				provider: GroqAIClientName,
				config:   &GroqProvider,
				messages: []llms.MessageContent{
					llms.TextParts(llms.ChatMessageTypeSystem, "You are a helpful AI assistant."),
					llms.TextParts(llms.ChatMessageTypeHuman, "Who was the second person to walk on the moon?"),
				},
			},
			want: true,
		}, {
			name: "Google vertex ai client test",
			args: args{
				provider: GoogleVertexAIClientName,
				config:   &GoogleProvider,
				messages: []llms.MessageContent{
					llms.TextParts(llms.ChatMessageTypeHuman, "Write a python function?"),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.args.provider)
			_ = client.Configure(tt.args.config)
			resp, _ := client.GetCompletion(context.Background(), tt.args.messages)
			_ = resp
			result, _ := json.Marshal(&resp)
			t.Logf("result: %s", string(result))
		})
	}
}

func Test_GoogleVertexGetCompletion(t *testing.T) {
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "Write a python function?"),
	}

	client := NewClient(GoogleVertexAIClientName)
	_ = client.Configure(&GoogleProvider)
	resp, _ := client.GetCompletion(context.Background(), messages, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))
	result, _ := json.Marshal(&resp)
	t.Logf("result: %s", string(result))
}
