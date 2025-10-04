package library

type Request struct {
	Model    string    `json:"model"`
	Temperature float32	`json:"temperature"`
	Messages []Message `json:"messages"`
}

type Response struct {
	ID       string   `json:"id"`
	Object   string   `json:"object"`
	Created  int64    `json:"created"`
	Model    string   `json:"model"`
	Metadata Metadata `json:"metadata"`
	Prompt   []any    `json:"prompt"`
	Choices  []Choice `json:"choices"`
	Usage    Usage    `json:"usage"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Metadata struct {
	WeightVersion string `json:"weight_version"`
}

type Choice struct {
	Index        int       `json:"index"`
	Message      MessageEx `json:"message"`
	Logprobs     any       `json:"logprobs"`
	FinishReason string    `json:"finish_reason"`
	Seed         any       `json:"seed"`
}

type MessageEx struct {
	Role      string `json:"role"`
	Content   string `json:"content"`
	ToolCalls []any  `json:"tool_calls"`
	Reasoning string `json:"reasoning"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
	CompletionTokens int `json:"completion_tokens"`
}
