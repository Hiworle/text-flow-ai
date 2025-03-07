package constant

import "context"

type Flow interface {
	Init(prompt string)
	Prompt() string
	Process(ctx context.Context, input string) (output string, err error)
	BatchProcess(ctx context.Context, inputs []string) (outputs []string, err error)
}
