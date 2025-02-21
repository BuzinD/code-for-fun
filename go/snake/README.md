# Snake Game in Go

## Description
This is a classic Snake game implemented in Go using the `tcell` library for terminal graphics. The game features a growing snake, an apple to eat, and an increasing speed mechanic as the score rises.

## Features
- Classic Snake gameplay.
- Speed increases every 2 points scored.
- Simple keyboard controls.
- Terminal-based graphics using `tcell`.
- Game over detection.
- Using goroutines and channel for handling user and game actions

## Installation
### Prerequisites
- Go (1.18+ recommended)

### Clone the code


### Install dependencies
```sh
go mod tidy
```

### Run the game
```sh
go run main.go
```

## Controls
- **Arrow Keys**: Move the snake (`Up`, `Down`, `Left`, `Right`)
- **P**: Pause the game
- **ESC**: Exit the game

## Game Rules
- Eat the `*` (apple) to grow and increase your score.
- The snake moves automatically.
- Every 2 apples, the speed increases.
- Hitting the wall or the snake itself ends the game.

## Code Overview
- `main.go` - Contains the game logic.
- `snake` - Struct representing the snake.
- `gameState` - Holds game status, score, and ticker for movement.
- `tcell` - Used for rendering the game in the terminal.

## Known Issues & TODOs
- Improve game-over screen.
- Add high-score tracking.
- Implement difficulty levels.

## License
This project is licensed under the MIT License.

