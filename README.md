# 🕵️ DarkSuitAI

⚡ Blazing production-ready library for building scalable reasoning AI systems ✨

[![Release Notes](https://img.shields.io/github/release/clive-alliance/darksuitAI?style=flat-square)](https://github.com/clive-alliance/darksuitAI/releases)
[![CI](https://github.com/clive-alliance/darksuitAI/actions/workflows/check_diffs.yml/badge.svg)](https://github.com/clive-alliance/darksuitAI/actions/workflows/check_diffs.yml)
[![GitHub star chart](https://img.shields.io/github/stars/clive-alliance/darksuitAI?style=flat-square)](https://star-history.com/#clive-alliance/darksuitAI)
[![Open Issues](https://img.shields.io/github/issues-raw/clive-alliance/darksuitAI?style=flat-square)](https://github.com/clive-alliance/darksuitAI/issues)
[![Open in Dev Containers](https://img.shields.io/static/v1?label=Dev%20Containers&message=Open&color=blue&logo=visualstudiocode&style=flat-square)](https://vscode.dev/redirect?url=vscode://ms-vscode-remote.remote-containers/cloneInVolume?url=https://github.com/clive-alliance/darksuitAI)
[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/clive-alliance/darksuitAI)



## Quick Install

```go
go get guthub.com/clive-alliance/darksuitAI@latest
```


## 🤔 What is DarkSuitAI?

**DarkSuitAI** is a framework for developing production-ready AI systems powered by large language models (LLMs).



## 🧱 What can you build with DarkSuitAI?


**🧱 Agent Powered AI System**

- [Documentation]()

**🤖 Chatbots**

- [Documentation]()

And much more!

## 🚀 How does DarkSuitAI bring you straight to production?
The main value props of the DarkSuitAI libraries are:
1. **Components**: composable building blocks, tools and integrations for working with language models. Components are modular and easy-to-use, and full scale production-ready for AI systems.
2. **Off-the-shelf chains**: built-in assemblages of components for accomplishing higher-level tasks

Off-the-shelf chains make it easy to get started. Components make it easy to customize existing chains and build new ones. 


## Components

Components fall into the following **modules**:

**📃 Model I/O**

This includes [prompt management](s), [prompt optimization](), a generic interface for [chat models](), and common utilities for working with [model outputs]().

```go

import (
	"fmt"
	"github.com/clive-alliance/darksuitAI"
)

func main() {

	os.Setenv("OPENAI_API_KEY", "YOUR_API_KEY")
    os.Setenv("ANTHROPIC_API_KEY", "Your Key") // anthropic API key

	args := darksuitAI.NewChatLLMArgs()

	args.SetChatInstruction("Your chat instruction goes here")
	args.AddPromptKey("year", "2024") // pass variables to your prompt
	args.SetModelType("openai", "gpt-4o") // set the model
	args.AddModelKwargs(1500, 0.8, true, []string{"Human:"}) // set model keyword arguments
	llm,err := args.NewLLM()
	if err != nil{
		// handle the error as you wish
	}
	resp,err:=llm.Chat("hello from earth🌍, what is your name?")
	if err != nil{
		// handle the error as you wish
	}
	fmt.Println(resp)

}

```


**🤖 Agents**

Agents allow an LLM autonomy over how a task is accomplished. Agents make decisions about which Actions to take, then take that Action, observe the result, and repeat until the task is complete. DarkSuitAI supercedes all other agentic frameworks/library through it's AI self-reflect action control.


## 💁 Contributing

As an open-source project in a rapidly developing field, we are extremely open to contributions, whether it be in the form of a new feature, improved infrastructure, or better documentation.
