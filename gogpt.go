package gogpt

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"io"
	"os"
)

type Context struct {
	header      *Header
	body        map[string]any
	content     string
	count       int
	maxTurns    int
	allowOutput bool
	output      io.Writer
}

type Header struct {
	Url  string
	Auth string
}

func NewContext() *Context {
	return &Context{
		body: map[string]any{
			"model":    "gpt-3.5-turbo",
			"messages": []map[string]string{},
			"stream":   false,
		},
		count:       0,
		maxTurns:    1,
		output:      os.Stdout,
		allowOutput: true,
	}
}

func (c *Context) post() {
	c.count++
	res, err := req.C().R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", c.header.Auth).
		SetBody(c.body).
		Post(c.header.Url)
	checkErr(err)
	if !c.allowOutput {
		c.discardOutput()
	}
	if !c.body["stream"].(bool) {
		// fmt.Println("defaultOutput: ")
		c.defaultOutput(res)
	} else {
		// fmt.Println("streamOutput: ")
		c.streamOutput(res)
	}
}

func (c *Context) Continue(prompt string) *Context {
	if c.count >= c.maxTurns {
		c.body["message"] = []map[string]string{}
		c.count = 0
	}
	c.appendMessage("user", prompt)
	c.post()
	return c
}

func (c *Context) GetContent() string {
	if len(c.content) == 0 {
		c.post()
	}
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

func (c *Context) discardOutput() {
	c.output = io.Discard
}

func (c *Context) defaultOutput(res *req.Response) {
	c.content = gjson.Get(res.String(), "choices.0.message.content").String()
	c.appendMessage("assistant", c.content)
	fmt.Fprintln(c.output, c.content)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(0)
	}
}
