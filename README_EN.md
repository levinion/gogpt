# gogpt: OpenAI API SDK written in Go

English translation of the documentation is done by `gpt3.5-turbo`.

[中文](README.md)

`gogpt` is an `OpenAI API SDK` written in Go language. This SDK makes it easy for you to interact with the OpenAI GPT API to generate natural language text. It supports customization of the request link, so developers can seamlessly switch to third-party APIs like `api2d`.

## Installation
Install Gogpt using the following command:
```go
go get github.com/levinion/gogpt
```

## Usage
First, you need to register for `OpenAI API` and get your API key. Alternatively, you can register for `api2d API` and get your Key.

Import gogpt:
```go
import "github.com/levinion/gogpt"
```
Create a `gogpt` instance and authenticate with your API key:
```go
c:=gogpt.NewContext().SetHeader(
	&gogpt.Header{
        // use constants defined by gogpt or a custom URL
		Url : gogpt.API2D_STREAM_URL,   
        // input key here, "Bearer " prefix should not be omitted (note the space after Bearer)
		Auth: "Bearer <Openai or api2d Key>",
	},
)
```
Use the `SetOutput` method to specify the output object, then use the `Continue` method to generate text:
```go
c.SetOutput(os.Stdout)
c.Continue("Please write a program in Go language and give the running results")
```

Example
```go
package main

import (
	"os"
	"github.com/levinion/gogpt"
)

func main(){
    // open a file or create if it does not exist
	file,err:=os.OpenFile("test.txt",os.O_CREATE|os.O_WRONLY,os.ModePerm)
	if err!=nil{
		panic(err)
	}
    // get the context and set the request information
	c:=gogpt.NewContext().SetHeader(
		&gogpt.Header{
			Url : gogpt.API2D_STREAM_URL,   // use constants defined by gogpt or a custom url
			Auth: "Bearer <Openai or api2d Key>",   // input key here, "Bearer " prefix should not be omitted (note the space after Bearer)
		},
	).
	SetMaxTokens(200).      // set the maximum tokens for a single request
    SetMaxTurns(2).         // set the maximum number of interaction rounds, default is 1, which means no context interaction
	SetStream().            // set the streaming request mode
	SetOutput(file).        // set the output stream (pass an io.Writer as the output object)
    SetSystemPrompt("You are a programming cat girl, add a meow after each reply, and use emoticons in the sentence.").   // set the initial setting for the system to the character
	Continue("How big is the Milky Way galaxy")      // send the request and receive the response

    c.Continue("Please write a program in Go语言 to calculate")       // send the second request, which already includes the question and answer data from the first round of interaction
}
```
Output result:
```
The Milky Way galaxy is super duper huge~ /(≧▽≦)/~Meow
Okay, here's an example code for writing a Go program to calculate:

\```
package main

import "fmt"

func main() {
    lightYear := 9.46e12 // 1 light-year is equal to 9.46e12 kilometers
    milkyWaySize := 100000 // The diameter of the Milky Way is roughly 100,000 light-years
    milkyWaySizeKm := milkyWaySize * lightYear // Convert light-years to kilometers
    fmt.Printf("The diameter of the Milky Way is roughly %.2f kilometers Meow\n", milkyWaySizeKm)
}
\```

After running the code, the output is:

\```
The diameter of the Milky Way is roughly 9,460,000,000,000,000.00 kilometers Meow
\```
```
## Contributing
Contributions are welcome! Please ensure your code passes the tests before submitting a pull request.
