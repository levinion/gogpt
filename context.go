package gogpt

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

type Context struct {
	header   *Header
	body     map[string]any
	content  string
	count    int
	maxTurns int
	output   io.Writer
	stream   bool
}

func NewContext() *Context {
	return &Context{
		body: map[string]any{
			"model":    "gpt-3.5-turbo",
			"messages": []map[string]string{},
			"stream":   false,
		},
		count:    0,
		maxTurns: 1,
		output:   nil,
		stream:   false,
	}
}

func NewContextWithWriter(writer io.Writer) *Context {
	return &Context{
		body: map[string]any{
			"model":    "gpt-3.5-turbo",
			"messages": []map[string]string{},
			"stream":   false,
		},
		count:    0,
		maxTurns: 1,
		output:   writer,
		stream:   false,
	}
}

func (c *Context) post() {
	res := c.query()
	defer res.Body.Close()
	c.defaultOutput(res)
}

func (c *Context) Continue(prompt string) *Context {
	c.checkAndReset()
	c.appendMessage("user", prompt)
	c.post()
	return c
}

func (c *Context) GetContent() string {
	if len(c.content) == 0 {
		fmt.Println("请求失败")
		os.Exit(0)
	}
	return c.content
}

func (c *Context) appendMessage(role string, content string) {
	c.body["messages"] =
		append(c.body["messages"].([]map[string]string), map[string]string{
			"role":    role,
			"content": content,
		})
}

func (c *Context) defaultOutput(res *http.Response) {
	content, err := io.ReadAll(res.Body)
	checkErr(err)
	c.content = gjson.Get(string(content), "choices.0.message.content").String()
	c.appendMessage("assistant", c.content)
	if c.output != nil {
		fmt.Fprintln(c.output, c.content)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(0)
	}
}
