package replica

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

type Replica struct {
	asciiArt []string
}

func NewReplica(aa []string) *Replica {
	return &Replica{
		asciiArt: aa,
	}
}

func (i *Replica) Print() {
	for i, line := range i.asciiArt {
		randomize(line, i)
	}
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
