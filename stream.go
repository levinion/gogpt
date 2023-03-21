package gogpt

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strings"

	"net/http"

	"github.com/tidwall/gjson"
)

type StreamContext struct {
	*Context
	channel chan string
}

func NewStreamContext() *StreamContext {
	return &StreamContext{
		Context: &Context{
			body: map[string]any{
				"model":    "gpt-3.5-turbo",
				"messages": []map[string]string{},
				"stream":   true,
			},
			count:    0,
			maxTurns: 1,
			output:   nil,
			stream:   true,
		},
		channel: nil,
	}
}

func NewStreamContextWithWriter(writer io.Writer) *StreamContext {
	return &StreamContext{
		Context: &Context{
			body: map[string]any{
				"model":    "gpt-3.5-turbo",
				"messages": []map[string]string{},
				"stream":   true,
			},
			count:    0,
			maxTurns: 1,
			output:   writer,
			stream:   true,
		},
		channel: nil,
	}
}

func NewStreamContextWithChannel(ch chan string) *StreamContext {
	return &StreamContext{
		Context: &Context{
			body: map[string]any{
				"model":    "gpt-3.5-turbo",
				"messages": []map[string]string{},
				"stream":   true,
			},
			count:    0,
			maxTurns: 1,
			output:   nil,
			stream:   true,
		},
		channel: ch,
	}
}

func (c *StreamContext) post() {
	res := c.query()
	defer res.Body.Close()
	if c.output == nil && c.channel != nil {
		c.streamOutputWithChannel(res)
	} else if c.output != nil && c.channel == nil {
		c.streamOutputWithWriter(res)
	} else {
		log.Fatal("Please choose how to return!")
	}
}

func (c *StreamContext) Continue(prompt string) *StreamContext {
	c.checkAndReset()
	c.appendMessage("user", prompt)
	c.post()
	return c
}

func (c *StreamContext) streamOutputWithWriter(res *http.Response) {
	reader := bufio.NewReader(res.Body)
	buf := bytes.NewBufferString("")
	writer := io.MultiWriter(buf, c.output)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		output := strings.TrimPrefix(string(line), "data: ")
		content := gjson.Get(output, "choices.0.delta.content").String()
		writer.Write([]byte(content))
	}
	writer.Write([]byte{'\n'})
	c.appendMessage("assistant", buf.String())
	c.content = buf.String()
}

func (c *StreamContext) streamOutputWithChannel(res *http.Response) {
	reader := bufio.NewReader(res.Body)
	buf := bytes.NewBufferString("")
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		output := strings.TrimPrefix(string(line), "data: ")
		content := gjson.Get(output, "choices.0.delta.content").String()
		buf.WriteString(content)
		c.channel <- content
	}
	buf.WriteString("\n")
	close(c.channel)
	c.appendMessage("assistant", buf.String())
	c.content = buf.String()
}
