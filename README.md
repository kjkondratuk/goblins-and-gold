# Goblins and Gold

A flexible, extensible text-based adventure game inspired by tabletop RPG systems.

![Version](https://img.shields.io/badge/version-0.1.0-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview

Goblins and Gold is a text-based adventure game that allows players to explore dungeons, interact with objects, and (eventually) battle monsters. The game is designed with extensibility in mind, making it easy to create and share custom worlds, monsters, and adventures.

### Key Features

- **Flexible World Design**: Create custom worlds using simple YAML files
- **Interactive Exploration**: Navigate through interconnected rooms and discover secrets
- **Container Interactions**: Open chests, loot items, and unlock locked containers
- **Skill Challenges**: Test your character's abilities with skill checks
- **Command-Line Interface**: Simple, intuitive command system with auto-completion

## Getting Started

### Installation

#### Option 1: Run from Source

```bash
# Clone the repository
git clone https://github.com/kjkondratuk/goblins-and-gold.git
cd goblins-and-gold

# Build and run
go build -o goblins-and-gold .
./goblins-and-gold  # On Linux/macOS
.\goblins-and-gold.exe  # On Windows
```

#### Option 2: Run with Make

```bash
# On Linux/macOS
make run

# On Windows (PowerShell)
go build -o bin\goblins-and-gold.exe .
.\bin\goblins-and-gold.exe
```

#### Option 3: Run Pre-built Binary

If you have a pre-built binary:

```bash
# On Linux/macOS
./goblins-and-gold

# On Windows
.\goblins-and-gold.exe
```

For detailed installation instructions, see the [Installation Guide](./docs/installation.md).

## Game Commands

- `look`: Examine your surroundings or a specific object
- `go`: Navigate to another room
- `interact`: Interact with objects in the environment
- `stats`: Display your character's statistics
- `quit`: Exit the game

## Creating Custom Adventures

Goblins and Gold is designed to be highly extensible. You can create custom adventures by defining:

- **Worlds**: Define rooms, paths, and containers
- **Monsters**: Create custom enemies with unique stats and attacks
- **Items**: Define weapons, armor, and other loot
- **Players**: Customize player characters with different stats and equipment

All game elements are defined in YAML files, making it easy to create and share custom content.

Example of a room definition:

```yaml
treasure-room:
  name: Treasure Room
  description: There are three chests located against the far wall.
  containers:
    - type: Chest
      description: A small wooden chest
      interactions: [Open, Loot]
      items:
        - type: Gold
          description: A few glittering coins
          quantity: 5
          unit: Coins
```

For detailed configuration options, see the [Game Configuration Guide](./docs/game_configuration.md).

## Game Mechanics

The game includes various mechanics such as:

- **Navigation**: Move between interconnected rooms
- **Interactions**: Open containers, loot items, unlock locks
- **Skill Checks**: Test character abilities against challenges
- **Combat** (planned): Turn-based combat with various attacks and defenses
- **Equipment** (planned): Equip weapons and armor for bonuses

For more details on game mechanics, see the [Game Mechanics Guide](./docs/game_mechanics.md).

## Documentation

- [Installation Guide](./docs/installation.md): Detailed instructions for installing and running the game
- [Game Configuration Guide](./docs/game_configuration.md): Documentation on configuring worlds, monsters, items, and players
- [Game Mechanics Guide](./docs/game_mechanics.md): Explanation of game mechanics and systems
- [Interactions](./interaction/interactions.md): Sequence diagrams showing how interactions work

## Roadmap

The following features are planned for future development:

1. Skill Checks (partially implemented)
2. Equipment (partially implemented)
3. Traps
4. Monsters (partially implemented)
5. Combat
6. NPC Interactions
7. Saving game state
8. A graphical tool to build worlds
9. Multiplayer

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License.
