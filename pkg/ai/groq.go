package ai

const GroqAIClientName = "groq"

type GroqAIClient struct {
	OpenAIClient
}

func (a *GroqAIClient) GetName() string {
	return GroqAIClientName
}
