# ttuy
Basic Terminal TUI

```console
go get github.com/electrikmilk/ttuy
```

```go
import "github.com/electrikmilk/ttuy"
```

## Components

Progress bar
```go
for i := 0; i <= 100; i++ {
    ttuy.Progress(i)
    time.Sleep(100 * time.Millisecond)
}
```

Typewriter
```go
ttuy.Typewriter("Typed out one character at a time")
```

Print contents of file, some ASCII art for example...
```go
ttuy.File("logo.txt")
```

Menu
``` go
ttuy.Menu("Title", []ttuy.Option{
    ttuy.Option{
      Label: "Option 1",
      Callback: func() {
        //
      },
    },
    ttuy.Option{
      Label: "Option 2",
      Callback: func() {
        //
      },
    },
    ...
  },
})
```

Input prompt
```go
var name string
ttuy.Ask("Enter your name", &name)
```

## Example

![Example](https://i.imgur.com/kLwzS6Q.gif)
