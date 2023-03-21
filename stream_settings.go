package gogpt

import "io"

func (c *StreamContext) SetHeader(header *Header) *StreamContext {
	c.header = header
	return c
}

func (c *StreamContext) SetBody(opt map[string]any) *StreamContext {
	for k, v := range opt {
		c.body[k] = v
	}
	return c
}

func (c *StreamContext) SetMaxTokens(max_tokens int) *StreamContext {
	c.SetBody(map[string]any{"max_tokens": max_tokens})
	return c
}

func (c *StreamContext) SetMaxTurns(max_turns int) *StreamContext {
	c.maxTurns = max_turns
	return c
}

func (c *StreamContext) SetSystemPrompt(prompt string) *StreamContext {
	c.appendMessage("system", prompt)
	return c
}

func (c *StreamContext) SetModel(model string) *StreamContext {
	c.body["model"] = model
	return c
}

func (c *StreamContext) SetOutput(output io.Writer) *StreamContext {
	c.output = output
	c.channel = nil
	return c
}

func (c *StreamContext) SetChannel(ch chan string) *StreamContext {
	c.output = nil
	c.channel = ch
	return c
}
