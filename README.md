# Terminal Racer

A multiplayer terminal racing game where up to 5 players can join a lobby and race each other.

## Requirements

- Go 1.23+

## Run

```bash
go run .
```

## Controls

| Key | Action |
|-----|--------|
| Arrow Up | Accelerate |
| Arrow Down | Brake |
| Arrow Left | Steer left |
| Arrow Right | Steer right |

## Project Structure

```
TerminalRacer/
├── main.go         # Game loop and entry point
├── player.go       # Player model and movement
├── track.go        # Track rendering and boundaries
├── timer.go        # Game timer utility
└── assets/
    ├── assets.go   # Embedded asset loader
    └── ui/static/gameModels/
        ├── car.png
        └── grass.png
```
