//go:build windows

package colr

import (
	"os"
	"syscall"
)

func init() {
	handle := syscall.Handle(os.Stdout.Fd())
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")

	// if ansi doesnt work bc idk why
	if _, _, err := setConsoleModeProc.Call(uintptr(handle), 0x0001|0x0002|0x0004); err != nil && err.Error() != "The operation completed successfully." {
		reset = ""

		// Special
		bold = ""
		dim = ""
		italic = ""
		underline = ""
		blink = ""
		strikethrough = ""

		// Text colors
		black = ""
		red = ""
		green = ""
		yellow = ""
		blue = ""
		purple = ""
		cyan = ""
		gray = ""
		white = ""

		// back colors
		blackBack = ""
		redBack = ""
		greenBack = ""
		yellowBack = ""
		blueBack = ""
		purpleBack = ""
		cyanBack = ""
		grayBack = ""
		whiteBack = ""
	}
}
