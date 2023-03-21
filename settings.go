package gogpt

import "io"

func (c *Context) SetHeader(header *Header) *Context {
	c.header = header
	return c
}

func (c *Context) SetBody(opt map[string]any) *Context {
	for k, v := range opt {
		c.body[k] = v
	}
	return c
}

func (c *Context) SetMaxTokens(max_tokens int) *Context {
	c.SetBody(map[string]any{"max_tokens": max_tokens})
	return c
}

func (c *Context) ToStreamContextWithCannel(ch chan string) *StreamContext {
	return &StreamContext{
		Context: &Context{
			body: map[string]any{
				"model":    c.body["model"],
				"messages": c.body["messages"],
				"stream":   true,
			},
			count:    c.count,
			maxTurns: c.maxTurns,
			output:   nil,
			stream:   true,
		},
		channel: ch,
	}
}

func (c *Context) ToStreamContextWithWriter() *StreamContext {
	return &StreamContext{
		Context: &Context{
			body: map[string]any{
				"model":    c.body["model"],
				"messages": c.body["messages"],
				"stream":   true,
			},
			count:    c.count,
			maxTurns: c.maxTurns,
			output:   c.output,
			stream:   true,
		},
		channel: nil,
	}
}

func (c *Context) SetMaxTurns(max_turns int) *Context {
	c.maxTurns = max_turns
	return c
}

func (c *Context) SetSystemPrompt(prompt string) *Context {
	c.appendMessage("system", prompt)
	return c
}

func (c *Context) SetModel(model string) *Context {
	c.body["model"] = model
	return c
}

func (c *Context) SetOutput(output io.Writer) *Context {
	c.output = output
	return c
}
