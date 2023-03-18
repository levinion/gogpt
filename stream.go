package gogpt

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
)

func (c *Context) streamOutput(res *req.Response) {
	reader := bufio.NewReader(res.Body)
	var buf bytes.Buffer
	writer := io.MultiWriter(&buf, c.output)
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
