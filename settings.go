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

func (c *Context) SetStream() *Context {
	c.SetBody(map[string]any{"stream": true})
	return c
}

func (c *Context) SetMaxTurns(max_turns int) *Context {
	c.maxTurns = max_turns
	return c
}

func (c *Context) SetSystemPrompt(prompt string) *Context {
	c.appendMessage("system", prompt)
	return c
}

func (c *Context) SetOutput(output io.Writer) *Context {
	c.output = output
	return c
}

func (c *Context) DiscardOutput() *Context {
	c.allowOutput = false
	return c
}
