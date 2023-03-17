package gogpt

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"os"
)

type Context struct {
	header  *Header
	body    *map[string]any
	content string
	count   int
}

type Header struct {
	Url  string
	Auth string
}

func NewContext() *Context {
	return &Context{
		body: &map[string]any{
			"model":    "gpt-3.5-turbo",
			"messages": []map[string]string{},
		},
		count: 0,
	}
}

func (c *Context) SetHeader(header *Header) *Context {
	c.header = header
	return c
}

func (c *Context) SetBody(opt map[string]any) *Context {
	for k, v := range opt {
		(*c.body)[k] = v
	}
	return c
}

func (c *Context) Post() {
	c.count++
	res, err := req.C().R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", c.header.Auth).
		SetBody(c.body).
		Post(c.header.Url)
	checkErr(err)
	c.content = getContent(res.String())
	c.appendMessage("assistant", c.content)
}

func (c *Context) Continue(prompt string) {
	c.appendMessage("user", prompt)
	c.Post()
}

func (c *Context) GetContent() string {
	if len(c.content) == 0 {
		c.Post()
	}
	if len(c.content) == 0 {
		fmt.Println("请求失败")
		os.Exit(0)
	}
	return c.content
}

func (c *Context) SetSystemPrompt(prompt string) *Context {
	c.appendMessage("system", prompt)
	return c
}

func (c *Context) appendMessage(role string, content string) {
	(*c.body)["messages"] =
		append((*c.body)["messages"].([]map[string]string), map[string]string{
			"role":    role,
			"content": content,
		})
}

func getContent(s string) string {
	return gjson.Get(s, "choices.0.message.content").String()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
