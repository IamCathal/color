# color
Simple and effective styled text output for golang

## Usage

```go
package main

import (
    "fmt"
    "github.com/iamcathal/color"
)

func main() {
	fmt.Println(StyledText("Nothing like a good 12 pack of Galahads!", "bold", "red"))
	fmt.Println(StyledText("Nothing like a good 12 pack of Galahads!", "bold", "yellow", "reverse"))
	fmt.Println(StyledText("Nothing like a good 12 pack of Galahads!", "black", "bold", "lightmagenta"))
	fmt.Println(StyledText("Nothing like a good 12 pack of Galahads!", "magenta", "bold", "backgroundgreen"))
	fmt.Println(StyledText("Nothing like a good 12 pack of Galahads!", "bold", "blue"))
	fmt.Println(StyledText("Nothing like a good 12 pack of Galahads!", "white", "underlined"))
	fmt.Println(StyledText("Nothing like a good 12 pack of Galahads!", "hidden"))
}

```

<img alt="Preview of styled text" src="https://i.imgur.com/HYpTS1c.png"> 