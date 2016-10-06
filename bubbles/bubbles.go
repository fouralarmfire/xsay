package bubbles

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Bubble struct {
	lines          []string
	defaultMessage string
	maxLen         int
}

func NewBubble(dm string) *Bubble {
	return &Bubble{
		defaultMessage: dm,
	}
}

func (b *Bubble) CustomMessage(text []string) {
	b.lines = text
	b.getMaxLineLength()
	if len(b.lines) == 1 {
		b.oneLineBubble()
	} else {
		b.multiLineBubble()
	}
}

func (b *Bubble) DefaultMessage() {
	b.lines = append(b.lines, b.defaultMessage)
	b.getMaxLineLength()
	b.oneLineBubble()
}

func (b *Bubble) multiLineBubble() {
	var del0, del1 string
	fmt.Printf(color.MagentaString("   %s\n", b.topLine()))
	for i, line := range b.lines {
		if i == 0 {
			del0 = b.delimeters("first", 0)
			del1 = b.delimeters("first", 1)
		} else if i == len(b.lines)-1 {
			del0 = b.delimeters("last", 0)
			del1 = b.delimeters("last", 1)
		} else {
			del0 = b.delimeters("middle", 0)
			del1 = b.delimeters("middle", 1)
		}
		fmt.Printf(color.CyanString(" %s  %s %s\n", del0, b.pad(strings.TrimSpace(line)), del1))
	}
	fmt.Printf(color.MagentaString("   %s\n", b.bottomLine()))
}

func (b *Bubble) oneLineBubble() {
	fmt.Printf(color.MagentaString("         %s\n", b.topLine()))
	fmt.Printf(color.CyanString("       %s  %s  %s\n", b.delimeters("only", 0), b.lines[0], b.delimeters("only", 1)))
	fmt.Printf(color.MagentaString("         %s\n", b.bottomLine()))
}

func (b *Bubble) delimeters(loc string, index int) string {
	var dels = map[string][]string{
		"first":  []string{"/", "\\"},
		"middle": []string{"|", "|"},
		"last":   []string{"\\", "/"},
		"only":   []string{"<", ">"},
	}
	return dels[loc][index]
}

func (b *Bubble) getMaxLineLength() {
	var length int
	for _, line := range b.lines {
		if len(line) > length {
			length = len(line)
		}
	}
	b.maxLen = length
}

func (b *Bubble) pad(text string) string {
	var padding = strings.Repeat(" ", b.maxLen-len(text))
	return text + padding
}

func (b *Bubble) topLine() string {
	return strings.Repeat("_", b.maxLen+2)
}

func (b *Bubble) bottomLine() string {
	return strings.Repeat("-", b.maxLen+2)
}
