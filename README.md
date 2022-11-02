![ttuy](https://user-images.githubusercontent.com/4368524/192105220-b950d506-7cd0-43da-ae55-6dc13d249dc0.png)

# ttuy

<p>
    <a href="https://github.com/electrikmilk/ttuy/actions/workflows/go.yml"><img src="https://github.com/electrikmilk/ttuy/actions/workflows/go.yml/badge.svg?branch=main" alt="Build"></a>
    <a href="https://pkg.go.dev/github.com/electrikmilk/ttuy?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/electrikmilk/ttuy"><img src="https://goreportcard.com/badge/github.com/electrikmilk/ttuy"/></a>
</p>

**ttuy** (pronounced _tee-tee-YU-Y_ or like _TUI_) is an easy-to-use, performant, procedural TUI framework.

Procedural as in, no models, no complex abstractions, run a function, and it just works.

```console
go get github.com/electrikmilk/ttuy
```

```go
import "github.com/electrikmilk/ttuy"
```

## Widgets

### User Input

```go
var name string
ttuy.Ask("Enter your name", &name)
```

![Ask](https://user-images.githubusercontent.com/4368524/192031382-ede561d5-bc1c-4939-aab2-87611db8341c.gif)

### Prompt

```go
if ttuy.Prompt("Install Program? This will take up 586kb.") {
    fmt.Println("Installing...")
} else {
    fmt.Println("Installation Canceled.")
}
```

![Prompt](https://user-images.githubusercontent.com/4368524/192075153-653de8e4-ba7b-4294-8dc2-73bcf264c53a.gif)

### Yes/No Prompt

```go
if ttuy.YesNo("Are you sure?") {
    fmt.Println("Installing...")
} else {
    fmt.Println("Installation Canceled.")
}
```

![YesNo](https://user-images.githubusercontent.com/4368524/192075315-f9b87357-94da-4fb5-9184-e292ae941b9d.png)

### Styled Output

```go
ttuy.Style("Text", ttuy.Bold, ttuy.Green)
```

![Style](https://user-images.githubusercontent.com/4368524/192031417-28a22355-fc20-49eb-913a-dcb90155ff07.png)

### Menu

``` go
ttuy.Menu("Title", []ttuy.Option{
    {
      Label: "Option 1",
      Callback: func() {
        // ...
      },
    },
    {
      Label: "Option 2",
      Callback: func() {
        // ...
      },
      Disabled: true,
    },
    // ...
})
```

![Menu](https://user-images.githubusercontent.com/4368524/192031437-b2cf1abe-c7d0-477f-a703-942bc7be290e.gif)

### Table

```go
var headers = []string{"Header 1", "Header 2", "Header 3", "Header 4"}
var rows []ttuy.Row
for i := 0; i < 5; i++ {
    var rowRows = []string{"Cell 1", "Cell 2", "Cell 3", "Cell 4"}
    rows = append(rows, ttuy.Row{Columns: rowRows})
}
ttuy.Table(headers, rows)
```

![Table](https://user-images.githubusercontent.com/4368524/192122215-eaef5af4-e08f-49c3-a401-9c515655dc34.png)

### Grid

```go
var alphabet = strings.Split("abcdefghijklmnopqrstuvwxyz", "")
var rows []ttuy.Row
for i := 0; i < 10; i++ {
    var rowRows []string
    for c := 1; c < 5; c++ {
        var cellContent string
        for w := 0; w < 26; w++ {
            cellContent += alphabet[w]
        }
        rowRows = append(rowRows, cellContent)
    }
    rows = append(rows, ttuy.Row{columns: rowRows})
}
fmt.Println(ttuy.Grid(rows))
```

![Grid](https://user-images.githubusercontent.com/4368524/193421608-71051c8d-8f77-4cd8-b1c6-ff685d41b46b.png)

### Viewport

```go
ttuy.Viewport(fileContents)
```

![Viewport](https://user-images.githubusercontent.com/4368524/192436746-c354968e-0094-4ca9-8f3c-2e0dbb4fc11b.gif)

### Progress Bar

```go
for i := 0; i <= 100; i++ {
    ttuy.ProgressBar(i)
    time.Sleep(100 * time.Millisecond)
}
```

![ProgressBar](https://user-images.githubusercontent.com/4368524/192673338-dbe6de5e-7225-4456-9a86-27bdd2702a33.gif)

### Spinners

```go
go ttuy.Spinner("Indeterminate Progress", ttuy.Ticker | ttuy.DotDotDot | ttuy.Throbber | ttuy.Blinker)
for i := 0; i < 50; i++ {
    time.Sleep(30 * time.Millisecond)
}
ttuy.StopSpinner()
```

![Spinner](https://user-images.githubusercontent.com/4368524/192109402-bc6691f9-e988-44de-b249-2c2b0f9a7bd0.png)

### Typewriter

```go
ttuy.Typewriter("Typed out one character at a time")
```

![Typewriter](https://user-images.githubusercontent.com/4368524/192031967-2643fbef-7290-4c0c-8f8e-e538f806472b.gif)

Or time it specifically...

```go
ttuy.TypewriterTimed("Typed out at duration...", 1000)
```

![TypewriterTimed](https://user-images.githubusercontent.com/4368524/192031989-fbb7b350-ddf0-4c84-897e-630e91e292df.gif)

Print contents of file, some ASCII art for example...

```go
ttuy.File("logo.txt")
```

![File](https://user-images.githubusercontent.com/4368524/192108479-306d0f11-b3f9-406b-86e6-e7ce4fde66f2.png)
