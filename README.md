# colr

An extremely lightweight cross-platform package to colorize text in terminals.

This is not meant for maximal compatibility, nor is it meant to handle a plethora of scenarios.
It will simply wrap a message with the necessary characters, if the OS handles it.

[Based on go-color](https://github.com/TwiN/go-color/blob/master/README.md)

## Usage

```console
go get g.tizu.dev/colr
```

### Using Functions

```go
package main

import "g.tizu.dev/colr"

func main() {
	// Special
	println(colr.Bold("hello!"))
	println(colr.Dim("hello!"))
	println(colr.Italic("hello!"))
	println(colr.Underline("hello!"))
	println(colr.Blink("hello!"))
	println(colr.Strikethrough("hello!"))

	// Text colors
	println(colr.Black("hello!"))
	println(colr.Red("hello!"))
	println(colr.Green("hello!"))
	println(colr.Yellow("hello!"))
	println(colr.Blue("hello!"))
	println(colr.Purple("hello!"))
	println(colr.Cyan("hello!"))
	println(colr.Gray("hello!"))
	println(colr.White("hello!"))

	// back colors
	println(colr.OnBlack("hello!"))
	println(colr.OnRed("hello!"))
	println(colr.OnGreen("hello!"))
	println(colr.OnYellow("hello!"))
	println(colr.OnBlue("hello!"))
	println(colr.OnPurple("hello!"))
	println(colr.OnCyan("hello!"))
	println(colr.OnGray("hello!"))
	println(colr.OnWhite("hello!"))

	// .. or combined
	println(colr.BlackOnBlack("hello!"))
	println(colr.BlackOnRed("hello!"))
	println(colr.BlackOnGreen("hello!"))
	println(colr.BlackOnYellow("hello!"))
	println(colr.BlackOnBlue("hello!"))
	println(colr.BlackOnPurple("hello!"))
	println(colr.BlackOnCyan("hello!"))
	println(colr.BlackOnGray("hello!"))
	println(colr.BlackOnWhite("hello!"))
	println(colr.RedOnBlack("hello!"))
	println(colr.RedOnRed("hello!"))
	println(colr.RedOnGreen("hello!"))
	println(colr.RedOnYellow("hello!"))
	println(colr.RedOnBlue("hello!"))
	println(colr.RedOnPurple("hello!"))
	println(colr.RedOnCyan("hello!"))
	println(colr.RedOnGray("hello!"))
	println(colr.RedOnWhite("hello!"))
	println(colr.GreenOnBlack("hello!"))
	println(colr.GreenOnRed("hello!"))
	println(colr.GreenOnGreen("hello!"))
	println(colr.GreenOnYellow("hello!"))
	println(colr.GreenOnBlue("hello!"))
	println(colr.GreenOnPurple("hello!"))
	println(colr.GreenOnCyan("hello!"))
	println(colr.GreenOnGray("hello!"))
	println(colr.GreenOnWhite("hello!"))
	println(colr.YellowOnBlack("hello!"))
	println(colr.YellowOnRed("hello!"))
	println(colr.YellowOnGreen("hello!"))
	println(colr.YellowOnYellow("hello!"))
	println(colr.YellowOnBlue("hello!"))
	println(colr.YellowOnPurple("hello!"))
	println(colr.YellowOnCyan("hello!"))
	println(colr.YellowOnGray("hello!"))
	println(colr.YellowOnWhite("hello!"))
	println(colr.BlueOnBlack("hello!"))
	println(colr.BlueOnRed("hello!"))
	println(colr.BlueOnGreen("hello!"))
	println(colr.BlueOnYellow("hello!"))
	println(colr.BlueOnBlue("hello!"))
	println(colr.BlueOnPurple("hello!"))
	println(colr.BlueOnCyan("hello!"))
	println(colr.BlueOnGray("hello!"))
	println(colr.BlueOnWhite("hello!"))
	println(colr.PurpleOnBlack("hello!"))
	println(colr.PurpleOnRed("hello!"))
	println(colr.PurpleOnGreen("hello!"))
	println(colr.PurpleOnYellow("hello!"))
	println(colr.PurpleOnBlue("hello!"))
	println(colr.PurpleOnPurple("hello!"))
	println(colr.PurpleOnCyan("hello!"))
	println(colr.PurpleOnGray("hello!"))
	println(colr.PurpleOnWhite("hello!"))
	println(colr.CyanOnBlack("hello!"))
	println(colr.CyanOnRed("hello!"))
	println(colr.CyanOnGreen("hello!"))
	println(colr.CyanOnYellow("hello!"))
	println(colr.CyanOnBlue("hello!"))
	println(colr.CyanOnPurple("hello!"))
	println(colr.CyanOnCyan("hello!"))
	println(colr.CyanOnGray("hello!"))
	println(colr.CyanOnWhite("hello!"))
	println(colr.GrayOnBlack("hello!"))
	println(colr.GrayOnRed("hello!"))
	println(colr.GrayOnGreen("hello!"))
	println(colr.GrayOnYellow("hello!"))
	println(colr.GrayOnBlue("hello!"))
	println(colr.GrayOnPurple("hello!"))
	println(colr.GrayOnCyan("hello!"))
	println(colr.GrayOnGray("hello!"))
	println(colr.GrayOnWhite("hello!"))
	println(colr.WhiteOnBlack("hello!"))
	println(colr.WhiteOnRed("hello!"))
	println(colr.WhiteOnGreen("hello!"))
	println(colr.WhiteOnYellow("hello!"))
	println(colr.WhiteOnBlue("hello!"))
	println(colr.WhiteOnPurple("hello!"))
	println(colr.WhiteOnCyan("hello!"))
	println(colr.WhiteOnGray("hello!"))
	println(colr.WhiteOnWhite("hello!"))
}
```

### Examples

```go
println("My name is " + colr.Green("John Doe") + " and I have " + colr.Red(32) + " apples.")
```
