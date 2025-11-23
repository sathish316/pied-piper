package config

type LLMConfig struct {
	Provider ModelProvider
	Model string
}

type ModelProvider string

const (
	ModelProviderOpenAI ModelProvider = "openai"
	ModelProviderAnthropic ModelProvider = "anthropic"
)
