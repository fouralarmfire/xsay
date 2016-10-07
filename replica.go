package xsay

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

type replica struct {
	asciiPath string
}

func newReplica(path string) *replica {
	return &replica{
		asciiPath: path,
	}
}

func (r *replica) readAscii() []string {
	lines := []string{}
	file, err := os.Open(r.asciiPath)
	if err != nil {
		return append(lines, "invalid file :(")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func (r *replica) printAscii() {
	for i, line := range r.readAscii() {
		r.randomize(line, i)
	}
}

func (r *replica) randomize(str string, ind int) {
	color.Set(r.randomColour(ind))
	defer color.Unset()
	fmt.Println(str)
}

func (r *replica) randomColour(ind int) color.Attribute {
	var colours = []color.Attribute{
		color.FgRed,
		color.FgGreen,
		color.FgYellow,
		color.FgBlue,
		color.FgMagenta,
		color.FgCyan,
	}

	rand.Seed(int64(ind-(-1)) * time.Now().Unix())
	return colours[rand.Int()%len(colours)]
}
