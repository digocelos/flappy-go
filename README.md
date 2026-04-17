# Flappy Go

Minimal Flappy Bird clone written in Go with [Ebitengine](https://ebitengine.org/). It keeps the core loop in a single package, handles gravity, jumping, pipes, scoring, and restart logic without extra dependencies.

## Running

1. `go run main.go` — launches the window and lets you play immediately.
2. `go build -o flappy main.go` — creates an executable named `flappy` that you can run directly.

The game tracks score internally, displays it via `ebitenutil.DebugPrintAt`, and you can restart after a crash by pressing `Space` again.
