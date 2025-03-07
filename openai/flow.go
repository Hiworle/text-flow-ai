package openai

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"os"
)

type OpenAiFlow struct {
	token   string
	baseUrl string
	client  *openai.Client
	prompt  string
	model   string
	ready   bool
}

func (o *OpenAiFlow) Prompt() string {
	return o.prompt
}

func (o *OpenAiFlow) Init(prompt string) {
	o.prompt = prompt
	o.ready = true
	return
}

func (o *OpenAiFlow) Process(ctx context.Context, input string) (output string, err error) {
	if !o.ready {
		return "", errors.New("flow not init")
	}
	resp, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: o.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: o.Prompt(),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input,
				},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
		},
	)
	if err != nil {
		return
	}
	return resp.Choices[0].Message.Content, nil
}

func (o *OpenAiFlow) BatchProcess(ctx context.Context, inputs []string) (outputs []string, err error) {
	if !o.ready {
		return nil, errors.New("flow not init")
	}
	messages := []openai.ChatCompletionMessage{{
		Role:    openai.ChatMessageRoleSystem,
		Content: o.Prompt(),
	}}
	for _, input := range inputs {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		})
	}
	resp, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    o.model,
			Messages: messages,
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
		},
	)
	if err != nil {
		return
	}
	for _, ch := range resp.Choices {
		outputs = append(outputs, ch.Message.Content)
	}
	return
}

func New(opts ...Option) *OpenAiFlow {
	o := &OpenAiFlow{
		baseUrl: "https://api.lkeap.cloud.tencent.com/v1",
		model:   "deepseek-v3",
		token:   os.Getenv("OPENAI_API_KEY"),
	}
	for _, op := range opts {
		op(o)
	}
	config := openai.DefaultConfig(o.token)
	config.BaseURL = o.baseUrl
	o.client = openai.NewClientWithConfig(config)
	return o
}

type Option = func(o *OpenAiFlow)

func WithToken(token string) Option {
	return func(o *OpenAiFlow) {
		o.token = token
	}
}

func WithBaseUrl(url string) Option {
	return func(o *OpenAiFlow) {
		o.baseUrl = url
	}
}

func WithModel(model string) Option {
	return func(o *OpenAiFlow) {
		o.model = model
	}
}
