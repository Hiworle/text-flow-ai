package constant

import (
	"fmt"
	"strings"
)

const PromptTmpl = `You are a text processing tool. Please process the input text according to the rules in the following example, and do not output any irrelevant characters.

### Example
%s

### Rules
%s

### Task
Please process the following input according to the rules above:
Input: "{{User Input}}"
Output: "{{Actual Output}}"
`
const _example = `Input: "%s"
Output: "%s"

`

type Example struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type Rule struct {
	Content string `json:"content"`
}

func MakeSystemPrompt(examples []Example, rules []Rule) string {
	exBuilder := strings.Builder{}
	for _, example := range examples {
		exBuilder.WriteString(fmt.Sprintf(_example, example.Input, example.Output))
	}
	rulesBuilder := strings.Builder{}
	for _, rule := range rules {
		rulesBuilder.WriteString(fmt.Sprintf("- %s\n", rule.Content))
	}
	return fmt.Sprintf(PromptTmpl, exBuilder.String(), rulesBuilder.String())
}
