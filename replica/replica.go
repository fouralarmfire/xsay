package replica

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

type Replica struct {
	asciiArt []string
}

func NewReplica(path string) *Replica {
	return &Replica{
		asciiArt: readAscii(path),
	}
}

func (i *Replica) Print() {
	for i, line := range i.asciiArt {
		randomize(line, i)
	}
}

func readAscii(path string) []string {
	lines := []string{}
	file, err := os.Open(path)
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

func randomize(str string, ind int) {
	color.Set(randomColour(ind))
	defer color.Unset()
	fmt.Println(str)
}

func randomColour(ind int) color.Attribute {
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
