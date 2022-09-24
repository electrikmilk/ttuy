![ttuy](https://user-images.githubusercontent.com/4368524/192105220-b950d506-7cd0-43da-ae55-6dc13d249dc0.png)

# ttuy

<p>
    <a href="https://pkg.go.dev/github.com/electrikmilk/ttuy?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
</p>

**ttuy** (pronounced _tee-tee-YU-Y_ or like _TUI_) is an easy-to-use procedural TUI Framework. Procedural as in, no models, no dealing with abstractions, just run a function, and it just works. It uses ANSI escape sequences for colors and manipulating the cursor.

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
ttuy.YesNo("Are you sure?")
```

![YesNo](https://user-images.githubusercontent.com/4368524/192075315-f9b87357-94da-4fb5-9184-e292ae941b9d.png)

### Styled Output

```go
ttuy.Style("Text", ttuy.Bold, ttuy.Green) // returns string
```

<img width="240" alt="Screen Shot 2022-09-23 at 14 06 15" src="https://user-images.githubusercontent.com/4368524/192031417-28a22355-fc20-49eb-913a-dcb90155ff07.png">

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
var rows []ttuy.row
for i := 0; i < 5; i++ {
	var rowRows = []string{"Cell 1", "Cell 2", "Cell 3", "Cell 4"}
	rows = append(rows, ttuy.row{columns: rowRows})
}
ttuy.Table(headers, rows)
```

![Screen Shot 2022-09-23 at 21 46 59](https://user-images.githubusercontent.com/4368524/192074972-ec501f77-a511-42ba-b2c1-cf0ac8a11661.png)

## Progress Bar

```go
for i := 0; i <= 100; i++ {
    ttuy.ProgressBar(i)
    time.Sleep(100 * time.Millisecond)
}
```

![ProgressBar](https://user-images.githubusercontent.com/4368524/192120185-cb9edca5-2fb3-4317-86c6-a8cc4e479770.gif)

### Spinners

```go
go Spinner("Indeterminate Progress", ttuy.Ticker | ttuy.DotDotDot | ttuy.Throbber | ttuy.Blinker)
for i := 0; i < 50; i++ {
    time.Sleep(30 * time.Millisecond)
}
StopSpinner()
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
