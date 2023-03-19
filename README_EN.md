

# gogpt: OpenAI API SDK Written in Go

`gogpt` is an `OpenAI API SDK` written in Go. This SDK allows you to easily interact with the OpenAI GPT API to generate natural language text. It supports custom request links, allowing developers to seamlessly switch to third-party APIs such as `api2d`.

## Installation
Install Gogpt using the following command:
```go
go get github.com/levinion/gogpt
```

## Usage
First, you need to register for an `OpenAI API` and obtain your API key. Alternatively, you can register for an `api2d API` and obtain your Key.

Import gogpt:
```go
import "github.com/levinion/gogpt"
```

Create a `gogpt` instance and authenticate with your API key:
```go
c:=gogpt.NewContext().SetHeader(
	&gogpt.Header{
		Url : gogpt.API2D_STREAM_URL,   
		Auth: "Bearer <Openai or api2d Key>",
	},
)
```
Use the `SetOutput` method to specify an output object, and then use the `Continue` method to generate text:
```go
c.SetOutput(os.Stdout)
c.Continue("Please write a program in Go that calculates and provides the result")
```

Example
```go
package main

import (
	"os"
	"github.com/levinion/gogpt/gogpt"
)

func main(){
    //Open a file, create it if it does not exist
	file,err:=os.OpenFile("test.txt",os.O_CREATE|os.O_WRONLY,os.ModePerm)
	if err!=nil{
		panic(err)
	}
    //Get context and set request information
	c:=gogpt.NewContext().SetHeader(
		&gogpt.Header{
			Url : gogpt.API2D_STREAM_URL,   
			Auth: "Bearer <Openai or api2d Key>",
		},
	).
	SetMaxTokens(200).      
    SetMaxTurns(2).        
	SetStream().           
	SetOutput(file).       
    SetSystemPrompt("You are a programming cat lady. Add 'meow' after each reply and use more emoticons in your sentences.").   
	Continue("How big is the Milky Way")      

    c.Continue("Please write a program in Go that calculates and provides the result")       
}
```
Output:
```
The Milky Way is very big meow! (*^▽^*) Its radius is about 50,000 light-years meow, including hundreds of billions of stars and many galaxies meow! (≧∇≦)ﾉ
Okay, here is the program in Go that calculates 1+2+3+...+100:

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1+2+3+...+100 =", sum, "meow!")
}
\```

Result:

\```
1+2+3+...+100 = 5050 meow!
\```

Hope this helps you meow~
```
## Contribution
Contributions and bug reports are welcome! Please ensure that your code passes tests before submitting a PR.
