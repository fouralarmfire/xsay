# xsay

Customisable & colourful version of `cowsay`.

Repeats args and stdin and prints a defult message if it doesn't feel like doing
either of those things.

```
go get github.com/fouralarmfire/xsay # will error but it's NBD
```

Put something like this in your own package:

main.go
```
package main

import (
  "fmt"

  "github.com/fouralarmfire/xsay"
)

func main() {
  x := xsay.New("woohoo.txt", "WOOHOO!")

  fmt.Print("\n")
  x.Say()
  fmt.Print("\n")
}
```

woohoo.txt
```
                \
                  ¶¶¶¶¶¶¶¶
             ¶¶¶¶¶        ¶¶¶¶¶
           ¶¶¶                 ¶¶¶
         ¶¶¶                     ¶¶¶
        ¶¶                         ¶¶
       ¶       ¶¶¶     ¶¶¶          ¶¶
      ¶          ¶¶      ¶¶          ¶¶
     ¶¶           ¶¶      ¶¶         ¶¶
     ¶             ¶¶      ¶¶   ¶¶¶   ¶¶
    ¶¶      ¶¶     ¶¶      ¶¶     ¶¶  ¶¶
    ¶¶    ¶¶¶      ¶¶      ¶¶      ¶¶ ¶¶
    ¶¶   ¶¶¶¶¶                  ¶¶ ¶¶ ¶¶
     ¶   ¶¶  ¶¶                 ¶¶    ¶¶
     ¶¶       ¶¶              ¶¶¶    ¶¶
      ¶¶       ¶¶            ¶¶¶     ¶¶
       ¶¶        ¶¶¶¶     ¶¶¶¶      ¶¶
        ¶¶          ¶¶¶¶¶¶¶        ¶¶
          ¶¶                     ¶¶
            ¶¶¶               ¶¶¶
               ¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶
```

Build and run for non-stop fun.
