# xsay

Customisable & colourful version of `cowsay`.

Repeats args and stdin and prints a defult message if it doesn't feel like doing
either of those things.

```
go get github.com/fouralarmfire/xsay # will error but it's NBD
```

Put something like this in your own package:

```
package main

import (
  "fmt"

  xsaypkg "github.com/fouralarmfire/xsay/mainframe"
)

func main() {
  me := xsaypkg.NewMainframe(myAscii(), defaultMessage())

  fmt.Print("\n")
  me.Say()
  fmt.Print("\n")
}

func myAscii() []string {
  return []string{
    "            \\",
    "              ¶¶¶¶¶¶¶¶",
    "         ¶¶¶¶¶        ¶¶¶¶¶",
    "       ¶¶¶                 ¶¶¶",
    "     ¶¶¶                     ¶¶¶",
    "    ¶¶                         ¶¶",
    "   ¶       ¶¶¶     ¶¶¶          ¶¶",
    "  ¶          ¶¶      ¶¶          ¶¶",
    " ¶¶           ¶¶      ¶¶         ¶¶",
    " ¶             ¶¶      ¶¶   ¶¶¶   ¶¶",
    "¶¶      ¶¶     ¶¶      ¶¶     ¶¶  ¶¶",
    "¶¶    ¶¶¶      ¶¶      ¶¶      ¶¶ ¶¶",
    "¶¶   ¶¶¶¶¶                  ¶¶ ¶¶ ¶¶",
    " ¶   ¶¶  ¶¶                 ¶¶    ¶¶",
    " ¶¶       ¶¶              ¶¶¶    ¶¶",
    "  ¶¶       ¶¶            ¶¶¶     ¶¶",
    "   ¶¶        ¶¶¶¶     ¶¶¶¶      ¶¶",
    "    ¶¶          ¶¶¶¶¶¶¶        ¶¶",
    "      ¶¶                     ¶¶",
    "        ¶¶¶               ¶¶¶",
    "           ¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶",
  }
}

func defaultMessage() string {
  return "WOOHOO!"
}
```

Build and run for non-stop fun.
