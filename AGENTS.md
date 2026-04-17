# Flappy Bird in Go with Ebitengine

## Objective

    Create a simple Flappy Bird clone in Go using Ebitengine.

## Project Rules

- Keep the implementation simple and educational.
- Do not add dependencies beyond the Go standard library and Ebitengine, unless
  clearly necessary.
- Always explain before making major changes.
- Implement in small, testable steps.
- Prioritize readable code over unnecessary abstractions.
- Separate game logic from rendering when it makes sense.
- Always ensure the project runs with go run ..

## Code Style Guidelines

- Use Go standard formatting (`gofmt`)
- Constants in UPPER_CASE with descriptive names
- Struct fields: PascalCase for public, camelCase for private
- Package-level variables and functions: PascalCase for public
- Use meaningful variables names (e.g., `screenWidth`, `playerHeight`)
- Group related constants together
- Embeded structs when appropriate (e.g. `Payer` embeds `Vector2D`)

## Error handling

- Use `log.Fatal(err)` for critical errors that should terminate
- Check errors immediately after operations that can fail
- Handle file loading errors gracefully

## Main Commands

- `go run main.go` - Run the game
- `go build -o flappy main.go` - Build executable
- `go test ./..` - Run all tests
- `go test --run TestName` - Run single test
- `go mod tidy` - Clean up dependencies

## Incremental Goals

1. Open a window
2. Draw the bird
3. Implement gravity
4. Implement jump
5. Create pipes
6. Detect collisions
7. Add scoring
8. Add restart
