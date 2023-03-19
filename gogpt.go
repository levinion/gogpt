package gogpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
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
	client := &http.Client{}
	jsonData, err := json.Marshal(c.body)
	checkErr(err)
	body := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", c.header.Url, body)
	checkErr(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.header.Auth)
	if c.body["stream"].(bool) {
		req.Header.Set("Accept", "text/event-stream")
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Connection", "keep-alive")
	}

	res, err := client.Do(req)
	checkErr(err)
	defer res.Body.Close()
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

func (c *Context) defaultOutput(res *http.Response) {
	content, err := io.ReadAll(res.Body)
	checkErr(err)
	c.content = gjson.Get(string(content), "choices.0.message.content").String()
	c.appendMessage("assistant", c.content)
	fmt.Fprintln(c.output, c.content)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(0)
	}
}
