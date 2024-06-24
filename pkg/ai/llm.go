package ai

import (
	"github.com/Stosan/groqgo"
	ant "github.com/clive-alliance/anthropicgo"
	"github.com/clive-alliance/darksuitai/internal/utilities"
	oai "github.com/clive-alliance/openaigo"
)

var llm LLM

func (ai AI) Chat(prompt string) (string, error) {
	kwargs := make([]map[string]interface{}, 5)
	for key := range ai.ModelType {
		switch key {
		case "anthropic":
			for k, v := range ai.ModelKwargs {
				kwargs[k] = map[string]interface{}{
					"model":          ai.ModelType["anthropic"],
					"max_tokens":     v.MaxTokens,
					"temperature":    v.Temperature,
					"stream":         v.Stream,
					"stop_sequences": v.StopSequences,
				}

			}

			llm = ant.ChatAnth(kwargs...)
		case "groq":
			for k, v := range ai.ModelKwargs {
				kwargs[k] = map[string]interface{}{
					"model":       ai.ModelType["groq"],
					"max_tokens":  v.MaxTokens,
					"temperature": v.Temperature,
					"stream":      v.Stream,
					"stop":        v.StopSequences,
				}

			}

			llm = groqgo.ChatGroq(kwargs...)

		case "openai":
			for k, v := range ai.ModelKwargs {
				kwargs[k] = map[string]interface{}{
					"model":          ai.ModelType["openai"],
					"max_tokens":     v.MaxTokens,
					"temperature":    v.Temperature,
					"stream":         v.Stream,
					"stop_sequences": v.StopSequences,
				}

			}
			llm = oai.ChatOAI(kwargs...)
		default:
			llm = nil
		}
	}
	promptMap := ai.PromptKeys
	promptMap["query"] = []byte(prompt)
	promptTemplate := utilities.CustomFormat(ai.ChatInstruction, promptMap)
	resp, err := llm.StreamCompleteChat(string(promptTemplate), "")
	return resp, err
}
