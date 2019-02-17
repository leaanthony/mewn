# Mewn

A zero dependency asset embedder.

## About

Mewn is perhaps the easiest way to embed assets in a Go program. Here is an example:

```Go
package main

import (
	"fmt"

	"github.com/leaanthony/mewn"
)

func main() {
	myTest := mewn.String("./assets/hello.txt")
	fmt.Println(myTest)
}
```

If compiled with `go build`, this example will read `hello.txt` from disk.
If compiled with `mewn build`, it will embed the assets into the resultant binary.

## Mewn cli command

The `mewn` command does 3 things:

- If you run `mewn`, it will recursively look for mewn.\* calls in your .go files. It will then generate intermeriary go files with assets embedded. It does not compile them into a final binary.
- `mewn build` will do the above, but compile all the source, then delete the intermediary files. This makes things a bit cleaner.
- `mewn pack` will do the same as `wails build`, but will compile with the go flags `-ldflags "-w -s"` to compress the final binary even more.

For the `build` and `pack` subcommands, any other cli parameters will be passed on to `go build`.

## Caveats

This project was built for simple embedding of assets and as such, there are a number of things to consider when choosing whether or not to use it.

- Paths to assets need to be unique. If you try to access 2 files with the same relative path, it isn't going to work.
- It is _extremely_ unlikely that any new features will be added in the future. This is by choice, not necessity. I want this project to be extremely stable so if you choose to use it today, it should work exactly the same in 3 years time. If it doesn't currently do what you want, you are probably looking for a different project.

Bug reports are _very_ welcome! Almost as much as PRs to fix them!

## What does 'Mewn' mean?

Mewn (mare-oon as fast as you can say it, not meee-oon) is the [Welsh](https://en.wikipedia.org/wiki/Welsh_language) word for "in".

## Why go for a crazy Welsh name?

Well, it stands out as a project name (practically zero name clashes), Welsh is one of the oldest and coolest languages in Europe and I speak it. JRR Tolkien (heard of him?) was [obsessed with Welsh](http://www.bbc.co.uk/guides/z2hthyc). So much so, he based the Middle Earth language "Sindarin" on it and there's strong evidence he based LOTR on Welsh mythology. Yeah. And it's associated with Red Dragons. And it's on Duolingo 😉.

## Inspiration

Heavy inspiration was drawn from [packr](https://github.com/gobuffalo/packr) by the awesome Mark Bates. The scope of what I needed was far narrower than the packr project, thus Mewn was born. If Mewn doesn't fulfil your needs, it's likely that packr will.
