package gradio

import (
	"context"
)

type GradioFlow struct {
	ctx    context.Context
	client *Gradio
	url    string
	token  string
}

func (g *GradioFlow) Process(prompt string, input string) (output string, err error) {
	g.client.SetSystemPrompt(prompt)
	return g.client.ChatCompletion(input), nil
}

func (g *GradioFlow) BatchProcess(prompt string, inputs []string) (outputs []string, err error) {
	for _, i := range inputs {
		output, _ := g.Process(prompt, i)
		outputs = append(outputs, output)
	}
	return outputs, nil
}

func New(ctx context.Context, opts ...Option) *GradioFlow {
	g := &GradioFlow{
		ctx: ctx,
	}
	for _, o := range opts {
		o(g)
	}
	g.client = NewClient(g.url, WithHfToken(g.token))
	return g
}

type Option = func(g *GradioFlow)

func WithGradioUrl(url string) Option {
	return func(g *GradioFlow) {
		g.url = url
	}
}

func WithGradioToken(token string) Option {
	return func(g *GradioFlow) {
		g.token = token
	}
}
