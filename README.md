# Shell Poker

Shell Poker is a cross-platform terminal Texas hold 'em poker game that doesn't involve money. Instead of currency, you play with chips, which reset after every game. It is completely open-source.

## How to play hold 'em poker?

You can learn to play hold' em [here](https://www.techopedia.com/gambling-guides/how-to-play-texas-holdem-poker).

## How do I install and run it?

1. If you haven't already, install The Go programming language here [here](https://go.dev/doc/install).
2. Open your shell as administrator and run ```go install github.com/Chopikashvili/shell-poker@latest```
3. Run ```[go/path]/bin/shell-poker``` in your preferred shell, where ```[go/path]``` is the GOPATH variable. View the GOPATH by running ```go env```. If you're using Git Bash, add ```winpty``` before the command or the game won't start. If you're using a Windows shell, replace forward slashes with back slashes and add ```.exe``` to the end of the path.

## How do I quit?

Press Ctrl+C to quit the game.

## What is the purpose of this?

It is my introductory project in Go!