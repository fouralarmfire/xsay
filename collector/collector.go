package collector

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

type TextCollector struct {
	lines []string
}

func NewTextCollector() *TextCollector {
	return &TextCollector{}
}

func (c *TextCollector) ReceivedInput() (bool, []string) {
	if (c.receivedStdin() || c.receivedArgs()) && c.checkLines() {
		return true, c.lines
	}
	return false, []string{}
}

func (c *TextCollector) receivedStdin() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			c.appendText(scanner.Text())
		}
		return true
	}

	return false
}

func (c *TextCollector) receivedArgs() bool {
	args := os.Args[1:]
	text := strings.Join(args, " ")
	if len(args) > 0 && len(text) < int(c.screenWidth())-13 {
		return c.appendText(text)
	}
	return false
}

func (c *TextCollector) checkLines() bool {
	if len(c.lines) == 0 {
		return false
	}
	return true
}

func (c *TextCollector) appendText(text string) bool {
	re := regexp.MustCompile(`[\S]`)
	if re.MatchString(text) {
		c.lines = append(c.lines, text)
		return true
	}
	return false
}

func (c *TextCollector) screenWidth() int {
	w, _, err := terminal.GetSize(0)
	if err != nil {
		fmt.Errorf("failed to get terminal size: %s", err)
		return 0
	}
	return w
}
