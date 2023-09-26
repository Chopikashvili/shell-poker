# Shell Poker

Shell Poker is a cross-platform terminal Texas hold 'em poker game that doesn't involve money. Instead of currency, you play with chips, which reset after every game. It is completely open-source.

## How to play hold 'em poker?

You can learn to play hold' em [here](https://www.techopedia.com/gambling-guides/how-to-play-texas-holdem-poker).

## How do I install and run it?

1. If you haven't already, install The Go programming language here [here](https://go.dev/doc/install).
2. Open your shell as administrator and run ```go install github.com/Chopikashvili/shell-poker@latest```
3. Unless you tinkered with the GOPATH, you need to run  ```C:/"Program Files"/Go/bin/shell-poker```. If you changed the GOPATH value, run ```[new/path]/bin/shell-poker``` in your preferred shell, where ```[new/path]``` is the GOPATH. If you're using Git Bash, add ```winpty``` before the command or the game won't start. If you're using a Windows shell, replace forward slashes with back slashes and add ```.exe``` to the end of the path.

## How do I quit?

Press Ctrl+C to quit the game.

## What is the purpose of this?

It is my introductory project in Go!