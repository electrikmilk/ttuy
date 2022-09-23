# ttuy

Easy-to-use Procedural Terminal TUI

```console
go get github.com/electrikmilk/ttuy
```

```go
import "github.com/electrikmilk/ttuy"
```

## Components

### User Input

```go
var name string
ttuy.Ask("Enter your name", &name)
```

### Styled Output

```go
ttuy.Style("Text", ttuy.Bold, ttuy.Green) // returns string
```

### Menu

``` go
ttuy.Menu("Title", []ttuy.Option{
    {
      Label: "Option 1",
      Callback: func() {
        //
      },
    },
    {
      Label: "Option 2",
      Callback: func() {
        //
      },
    },
    // ...
  },
})
```

### Progress bar

```go
for i := 0; i <= 100; i++ {
    ttuy.Progress(i)
    time.Sleep(100 * time.Millisecond)
}
```

### Typewriter

```go
ttuy.Typewriter("Typed out one character at a time")

// Or time it specifically...
ttuy.TypewriterTimed("Typed out at duration...", 1000)
```

Print contents of file, some ASCII art for example...
```go
ttuy.File("logo.txt")
```

## Example

![Example](https://i.imgur.com/kLwzS6Q.gif)
