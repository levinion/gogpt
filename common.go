package gogpt

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Header struct {
	Url  string
	Auth string
}

func (c *Context) query() *http.Response {
	c.count++
	client := &http.Client{}
	jsonData, err := json.Marshal(c.body)
	checkErr(err)
	body := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", c.header.Url, body)
	checkErr(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.header.Auth)
	if c.stream {
		req.Header.Set("Accept", "text/event-stream")
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Connection", "keep-alive")
	}
	res, err := client.Do(req)
	checkErr(err)
	return res
}

func (c *Context) checkAndReset() {
	if c.count >= c.maxTurns {
		c.Clear()
	}
}

func (c *Context) Clear(){
	c.body["messages"] = []map[string]string{}
	c.count = 0
}
