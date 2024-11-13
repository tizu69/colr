package colr

import "fmt"

var (
	reset = "\033[0m"

	// Special
	bold          = "\033[1m"
	dim           = "\033[2m"
	italic        = "\033[3m"
	underline     = "\033[4m"
	blink         = "\033[5m"
	strikethrough = "\033[9m"

	// Text colors
	black  = "\033[30m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"

	// back colors
	blackBack  = "\033[40m"
	redBack    = "\033[41m"
	greenBack  = "\033[42m"
	yellowBack = "\033[43m"
	blueBack   = "\033[44m"
	purpleBack = "\033[45m"
	cyanBack   = "\033[46m"
	grayBack   = "\033[47m"
	whiteBack  = "\033[107m"
)

func use(color string, s any) string {
	switch s := s.(type) {
	case string:
		return color + s + reset
	default:
		return color + fmt.Sprint(s) + reset
	}
}

func Bold(s any) string          { return use(bold, s) }
func Dim(s any) string           { return use(dim, s) }
func Italic(s any) string        { return use(italic, s) }
func Underline(s any) string     { return use(underline, s) }
func Blink(s any) string         { return use(blink, s) }
func Strikethrough(s any) string { return use(strikethrough, s) }

func Black(s any) string  { return use(black, s) }
func Red(s any) string    { return use(red, s) }
func Green(s any) string  { return use(green, s) }
func Yellow(s any) string { return use(yellow, s) }
func Blue(s any) string   { return use(blue, s) }
func Purple(s any) string { return use(purple, s) }
func Cyan(s any) string   { return use(cyan, s) }
func Gray(s any) string   { return use(gray, s) }
func White(s any) string  { return use(white, s) }

func OnBlack(s any) string  { return use(blackBack, s) }
func OnRed(s any) string    { return use(redBack, s) }
func OnGreen(s any) string  { return use(greenBack, s) }
func OnYellow(s any) string { return use(yellowBack, s) }
func OnBlue(s any) string   { return use(blueBack, s) }
func OnPurple(s any) string { return use(purpleBack, s) }
func OnCyan(s any) string   { return use(cyanBack, s) }
func OnGray(s any) string   { return use(grayBack, s) }
func OnWhite(s any) string  { return use(whiteBack, s) }

func BlackOnBlack(s any) string   { return use(black+blackBack, s) }
func BlackOnRed(s any) string     { return use(black+redBack, s) }
func BlackOnGreen(s any) string   { return use(black+greenBack, s) }
func BlackOnYellow(s any) string  { return use(black+yellowBack, s) }
func BlackOnBlue(s any) string    { return use(black+blueBack, s) }
func BlackOnPurple(s any) string  { return use(black+purpleBack, s) }
func BlackOnCyan(s any) string    { return use(black+cyanBack, s) }
func BlackOnGray(s any) string    { return use(black+grayBack, s) }
func BlackOnWhite(s any) string   { return use(black+whiteBack, s) }
func RedOnBlack(s any) string     { return use(red+blackBack, s) }
func RedOnRed(s any) string       { return use(red+redBack, s) }
func RedOnGreen(s any) string     { return use(red+greenBack, s) }
func RedOnYellow(s any) string    { return use(red+yellowBack, s) }
func RedOnBlue(s any) string      { return use(red+blueBack, s) }
func RedOnPurple(s any) string    { return use(red+purpleBack, s) }
func RedOnCyan(s any) string      { return use(red+cyanBack, s) }
func RedOnGray(s any) string      { return use(red+grayBack, s) }
func RedOnWhite(s any) string     { return use(red+whiteBack, s) }
func GreenOnBlack(s any) string   { return use(green+blackBack, s) }
func GreenOnRed(s any) string     { return use(green+redBack, s) }
func GreenOnGreen(s any) string   { return use(green+greenBack, s) }
func GreenOnYellow(s any) string  { return use(green+yellowBack, s) }
func GreenOnBlue(s any) string    { return use(green+blueBack, s) }
func GreenOnPurple(s any) string  { return use(green+purpleBack, s) }
func GreenOnCyan(s any) string    { return use(green+cyanBack, s) }
func GreenOnGray(s any) string    { return use(green+grayBack, s) }
func GreenOnWhite(s any) string   { return use(green+whiteBack, s) }
func YellowOnBlack(s any) string  { return use(yellow+blackBack, s) }
func YellowOnRed(s any) string    { return use(yellow+redBack, s) }
func YellowOnGreen(s any) string  { return use(yellow+greenBack, s) }
func YellowOnYellow(s any) string { return use(yellow+yellowBack, s) }
func YellowOnBlue(s any) string   { return use(yellow+blueBack, s) }
func YellowOnPurple(s any) string { return use(yellow+purpleBack, s) }
func YellowOnCyan(s any) string   { return use(yellow+cyanBack, s) }
func YellowOnGray(s any) string   { return use(yellow+grayBack, s) }
func YellowOnWhite(s any) string  { return use(yellow+whiteBack, s) }
func BlueOnBlack(s any) string    { return use(blue+blackBack, s) }
func BlueOnRed(s any) string      { return use(blue+redBack, s) }
func BlueOnGreen(s any) string    { return use(blue+greenBack, s) }
func BlueOnYellow(s any) string   { return use(blue+yellowBack, s) }
func BlueOnBlue(s any) string     { return use(blue+blueBack, s) }
func BlueOnPurple(s any) string   { return use(blue+purpleBack, s) }
func BlueOnCyan(s any) string     { return use(blue+cyanBack, s) }
func BlueOnGray(s any) string     { return use(blue+grayBack, s) }
func BlueOnWhite(s any) string    { return use(blue+whiteBack, s) }
func PurpleOnBlack(s any) string  { return use(purple+blackBack, s) }
func PurpleOnRed(s any) string    { return use(purple+redBack, s) }
func PurpleOnGreen(s any) string  { return use(purple+greenBack, s) }
func PurpleOnYellow(s any) string { return use(purple+yellowBack, s) }
func PurpleOnBlue(s any) string   { return use(purple+blueBack, s) }
func PurpleOnPurple(s any) string { return use(purple+purpleBack, s) }
func PurpleOnCyan(s any) string   { return use(purple+cyanBack, s) }
func PurpleOnGray(s any) string   { return use(purple+grayBack, s) }
func PurpleOnWhite(s any) string  { return use(purple+whiteBack, s) }
func CyanOnBlack(s any) string    { return use(cyan+blackBack, s) }
func CyanOnRed(s any) string      { return use(cyan+redBack, s) }
func CyanOnGreen(s any) string    { return use(cyan+greenBack, s) }
func CyanOnYellow(s any) string   { return use(cyan+yellowBack, s) }
func CyanOnBlue(s any) string     { return use(cyan+blueBack, s) }
func CyanOnPurple(s any) string   { return use(cyan+purpleBack, s) }
func CyanOnCyan(s any) string     { return use(cyan+cyanBack, s) }
func CyanOnGray(s any) string     { return use(cyan+grayBack, s) }
func CyanOnWhite(s any) string    { return use(cyan+whiteBack, s) }
func GrayOnBlack(s any) string    { return use(gray+blackBack, s) }
func GrayOnRed(s any) string      { return use(gray+redBack, s) }
func GrayOnGreen(s any) string    { return use(gray+greenBack, s) }
func GrayOnYellow(s any) string   { return use(gray+yellowBack, s) }
func GrayOnBlue(s any) string     { return use(gray+blueBack, s) }
func GrayOnPurple(s any) string   { return use(gray+purpleBack, s) }
func GrayOnCyan(s any) string     { return use(gray+cyanBack, s) }
func GrayOnGray(s any) string     { return use(gray+grayBack, s) }
func GrayOnWhite(s any) string    { return use(gray+whiteBack, s) }
func WhiteOnBlack(s any) string   { return use(white+blackBack, s) }
func WhiteOnRed(s any) string     { return use(white+redBack, s) }
func WhiteOnGreen(s any) string   { return use(white+greenBack, s) }
func WhiteOnYellow(s any) string  { return use(white+yellowBack, s) }
func WhiteOnBlue(s any) string    { return use(white+blueBack, s) }
func WhiteOnPurple(s any) string  { return use(white+purpleBack, s) }
func WhiteOnCyan(s any) string    { return use(white+cyanBack, s) }
func WhiteOnGray(s any) string    { return use(white+grayBack, s) }
func WhiteOnWhite(s any) string   { return use(white+whiteBack, s) }
