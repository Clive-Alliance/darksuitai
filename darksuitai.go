package darksuitai

import (
	"github.com/clive-alliance/darksuitai/internal/prompts"
	"github.com/clive-alliance/darksuitai/pkg/ai"
	"github.com/clive-alliance/darksuitai/types"
)

// DarkSuitAI is the main struct that users will interact with
type ChatLLMArgs types.ChatLLMArgs

// NewChatLLMArgs creates a new ChatLLMArgs with default values
func NewChatLLMArgs() *ChatLLMArgs {

	return &ChatLLMArgs{
		ChatInstruction: prompts.PromptTemplate,
		PromptKeys:      make(map[string][]byte),
		ModelType:       make(map[string]string),
		ModelKwargs: []struct {
			MaxTokens     int    `json:"max_tokens"`
			Temperature   float64  `json:"temperature"`
			Stream        bool     `json:"stream"`
			StopSequences []string `json:"stop_sequences"`
		}{
			{
				MaxTokens:     500,
				Temperature:   0.2,
				Stream:        false,
				StopSequences: []string{},
			},
		},
	}
}

/*
	SetChatInstruction sets the chat instruction in ChatLLMArgs.

This method allows you to define the main instruction or prompt that will guide the chat model's responses.

Example:

args := darksuitAI.NewChatLLMArgs()

args.SetChatInstruction([]byte("Your chat instruction goes here"))

In this example, the byte slice containing the chat instruction is set, which will be used by the chat model to generate responses.
*/
func (args *ChatLLMArgs) SetChatInstruction(prompt []byte) {
	args.ChatInstruction = prompt
}

/*
	AddPromptKey adds a key-value pair to the PromptKeys map in ChatLLMArgs.

This method allows you to dynamically insert or update prompt-specific variables that can be used within the chat instruction template.

Example:

args := darksuitAI.NewChatLLMArgs()

args.AddPromptKey("year", []byte(`2024`))

args.AddPromptKey("month", []byte(`June`))

In this example, the keys "year" and "month" with their respective values "2024" and "June" are added to the PromptKeys map, which can then be referenced in the chat instruction template.
*/
func (args *ChatLLMArgs) AddPromptKey(key string, value []byte) {
	args.PromptKeys[key] = value
}

/*
	SetModelType sets a key-value pair in the ModelType map in ChatLLMArgs.

This method allows you to specify the type of model to be used for the chat.

Example:

args := darksuitAI.NewChatLLMArgs()

args.SetModelType("openai", "gpt-4o")

In this example, the key "openai" with the value "gpt-4o" is added to the ModelType map, indicating the model type to be used.
*/
func (args *ChatLLMArgs) SetModelType(key, value string) {
	args.ModelType[key] = value
}

/*
	AddModelKwargs adds a new set of model arguments to the ModelKwargs slice in ChatLLMArgs.

This method allows you to specify various parameters for the model's behavior.

Example:

args := darksuitAI.NewChatLLMArgs()

args.AddModelKwargs(500, 0.8, true, []string{"Human:"})

In this example, the model arguments are set with a maximum of 1500 tokens, a temperature of 0.8, streaming enabled, and a stop sequence of "Human:".
*/
func (args *ChatLLMArgs) AddModelKwargs(maxTokens int, temperature float64, stream bool, stopSequences []string) {
	args.ModelKwargs = append(args.ModelKwargs, struct {
		MaxTokens     int    `json:"max_tokens"`
		Temperature   float64  `json:"temperature"`
		Stream        bool     `json:"stream"`
		StopSequences []string `json:"stop_sequences"`
	}{
		MaxTokens:     maxTokens,
		Temperature:   temperature,
		Stream:        stream,
		StopSequences: stopSequences,
	})
}

type LLM struct {
	ai ai.AI
}

// NewLLM creates a new instance of DarkSuitAI LLM
func (cargs *ChatLLMArgs) NewLLM() (*LLM, error) {

	return &LLM{
		ai: ai.AI{
			ChatInstruction: cargs.ChatInstruction,
			PromptKeys:      cargs.PromptKeys,
			ModelType:       cargs.ModelType,
			ModelKwargs:     cargs.ModelKwargs,
		},
	}, nil
}

// Chat LLM exposes the LLM method for chat
func (d *LLM) Chat(prompt string) (string, error) {
	return d.ai.Chat(prompt)
}

// Stream LLM exposes the LLM method for chat stream
func (d *LLM) Stream(prompt string) (string, error) {
	return d.ai.Chat(prompt)
}
