
# Roma X Game Server API

## Overview
This repository contains the server implementation for the Roma X slot game, providing a RESTful API to handle game mechanics including spins, payouts, and special features. The implementation follows a clean architecture pattern using the Fiber web framework.

## Game Understanding

### Core Game Mechanics
Roma X is a 5-reel, 3-row slot game with the following key features:

1. **Basic Structure**
- 5×3 grid layout
- 15 fixed paylines
- Left-to-right winning combinations
- Win calculation: (Bet × Pay) ÷ 15

2. **Symbols & Payouts**
```
High-Value Symbols:
- Helmet: 3x(18), 4x(180), 5x(1200)
- Eye: 3x(18), 4x(120), 5x(600)
- Lion: 3x(12), 4x(60), 5x(240)

Medium-Value Symbols:
- Sword: 3x(12), 4x(36), 5x(144)
- Shield: 3x(6), 4x(18), 5x(90)

Low-Value Symbols:
- Cup: 3x(6), 4x(12), 5x(36)
- Grapes: 3x(6), 4x(12), 5x(36)

Special Symbols:
- Wild: Substitutes all symbols except Bonus
- Bonus: Triggers Bonus Game
```

3. **Special Features**
- Free Games:
  * 4 matches = 3 free games
  * 5 matches = 5 free games
  * 6 matches = 10 free games
  * 7 matches = 20 free games
- Bonus Game:
  * Triggered by 3 Bonus symbols
  * Combat mechanics with options:
    - Dual swords (20x bet)
    - Single sword (10x bet)
    - Lion (3x bet)

### Game Constraints
- Minimum bet: 1.2
- Maximum bet: 1,200
- Maximum payout: 12,000,000
- Maximum win odds: 10,000x

## Technical Implementation

### Project Structure
```
roma-x-api/
├── engine/
│   ├── game/
│   │   ├── constants.go
│   │   ├── models.go
│   │   ├── paylines.go
│   │   └── symbols.go
│   ├── handlers/
│   │   └── spin.go
│   └── service/
│       └── game_service.go
├── main.go
├── go.mod
└── README.md
```

### API Endpoint

#### POST /spin
Handles game spin requests and generates outcomes.

Request:
```json
{
    "bet": 10.0,
    "player_id": "player123"
}
```

Response:
```json
{
    "grid": {
        "symbols": [["Symbol", "Symbol", "Symbol"], ...]
    },
    "winning_lines": [
        {
            "payline_number": 1,
            "symbols": "Lion",
            "count": 3,
            "payout": 80.0
        },...
    ],
    "total_win": 80.0,
    "free_games": 0,
    "is_bonus_game": false
}
```

### Key Features
1. **Random Outcome Generation**
   - 30% base win probability
   - 3% bonus game trigger chance
   - 5% free games trigger chance

2. **Payline Handling**
   - 15 unique payline patterns
   - All patterns start from leftmost reel
   - Wild symbol substitutions
   - Multiple winning lines possible

3. **Payout Calculation**
   - Symbol-specific payout tables
   - Bet multiplication
   - Maximum payout limits
   - Bonus game multipliers

## Setup and Running

1. Clone the repository
```bash
git clone https://github.com/Techbite-sudo/Roma-x-api.git
cd roma-x-api
```

2. Install dependencies
```bash
go mod tidy
```

3. Run the server
```bash
go run main.go
```

The server will start on port 3000.

## Testing
```bash
curl -X POST http://localhost:3000/spin \
     -H "Content-Type: application/json" \
     -d '{"bet": 10.0, "player_id": "player123"}'
```

## Dependencies
- Fiber v2 - Web framework
- Go 1.21 or higher

## Error Handling
- Invalid bet amounts
- Request validation
- Maximum payout limits
- System malfunctions

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Author
[Boniface Mwema]

## Acknowledgments
- Game design inspiration from Roma X slot game
- Fiber framework community
```