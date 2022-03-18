package cli

import "fmt"

type Color string

const (
	ColorRed   Color = "\u001b[31m"
	ColorGreen Color = "\u001b[32m"
	ColorBlue  Color = "\u001b[36m"
	ColorReset Color = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(fmt.Sprintf("%s%s%s", string(color), message, string(ColorReset)))
}
