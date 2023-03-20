package gogpt

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/tidwall/gjson"
	"net/http"
)

func (c *Context) streamOutput(res *http.Response) {
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
		// fmt.Println(i,content)
		writer.Write([]byte(content))
	}
	writer.Write([]byte{'\n'})
	c.appendMessage("assistant", buf.String())
	c.content = buf.String()
	c.done=true
}
