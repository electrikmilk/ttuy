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

![Ask](https://user-images.githubusercontent.com/4368524/192031382-ede561d5-bc1c-4939-aab2-87611db8341c.gif)

### Prompt

```go
if Prompt("Install Program? This will take up 586kb.") {
	fmt.Println("Installing...")
} else {
	fmt.Println("Installation Canceled.")
}
```

![Prompt](https://user-images.githubusercontent.com/4368524/192075069-9595cb54-5b9a-4095-a46d-dccc63924e6e.png)

### Yes/No Prompt

```go
YesNo("Are you sure?")
```

![YesNo](https://user-images.githubusercontent.com/4368524/192075153-653de8e4-ba7b-4294-8dc2-73bcf264c53a.gif)

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
var rows []row
for i := 0; i < 5; i++ {
	var rowRows = []string{"Cell 1", "Cell 2", "Cell 3", "Cell 4"}
	rows = append(rows, row{columns: rowRows})
}
Table(headers, rows)
```

![Screen Shot 2022-09-23 at 21 46 59](https://user-images.githubusercontent.com/4368524/192074972-ec501f77-a511-42ba-b2c1-cf0ac8a11661.png)

### Progress bar

```go
for i := 0; i <= 100; i++ {
    ttuy.Progress(i)
    time.Sleep(100 * time.Millisecond)
}
```

![Progress](https://user-images.githubusercontent.com/4368524/192031937-7c725cb2-b2eb-44ee-ba55-3f581ffe6d1c.gif)

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
