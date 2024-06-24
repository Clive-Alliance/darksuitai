package ai

import (
	"github.com/Stosan/groqgo"
	ant "github.com/clive-alliance/anthropicgo"
	"github.com/clive-alliance/darksuitAI/internal/utilities"
	oai "github.com/clive-alliance/openaigo"
)

var llm LLM

func (ai AI) Chat(prompt string) (string, error) {
	for key := range ai.ModelType {
		switch key {
		case "anthropic":
			kwargs := []map[string]interface{}{{
				"model":          ai.ModelType["anthropic"],
				"max_tokens":     3000,
				"temperature":    0.0,
				"stream":         true,
				"stop_sequences": []string{"\nObservation"},
			}}

			llm = ant.ChatAnth(kwargs...)
		case "groq":
			kwargs := []map[string]interface{}{{
				"model":       ai.ModelType["groq"],
				"temperature": 0.2,
				"max_tokens":  3000,
				"stream":      true,
				"stop":        []string{"Observation"},
			}}

			llm = groqgo.ChatGroq(kwargs...)

		case "openai":
			kwargs := []map[string]interface{}{{
				"model":       ai.ModelType["openai"],
				"temperature": 0.2,
				"max_tokens":  3000,
				"stream":      true,
				"stop":        []string{"Observation"},
			}}

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
