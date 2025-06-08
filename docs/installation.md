# Installation and Execution Guide

This document provides instructions for installing and running Goblins and Gold, both from source code and as a standalone binary.

## Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher (for building from source)
- Git (for cloning the repository)

## Building from Source

### Clone the Repository

```bash
git clone https://github.com/kjkondratuk/goblins-and-gold.git
cd goblins-and-gold
```

### Build and Run

The project includes a Makefile that simplifies the build and run process:

```bash
# On Linux/macOS
make run

# On Windows (PowerShell)
go build -o bin\goblins-and-gold.exe .
.\bin\goblins-and-gold.exe
```

### Manual Build

If you prefer to build manually without using the Makefile:

```bash
# Build the executable
go build -o goblins-and-gold .

# Run the executable
./goblins-and-gold  # On Linux/macOS
.\goblins-and-gold.exe  # On Windows
```

## Running as a Standalone Binary

If you have downloaded a pre-built binary:

1. Extract the archive if necessary
2. Open a terminal/command prompt
3. Navigate to the directory containing the binary
4. Run the executable:

```bash
# On Linux/macOS
./goblins-and-gold

# On Windows
.\goblins-and-gold.exe
```

## Debug Mode

To run the game in debug mode, which enables additional debug commands:

```bash
# On Linux/macOS
DEBUG=true ./goblins-and-gold

# On Windows (PowerShell)
$env:DEBUG="true"; .\goblins-and-gold.exe
```

## Game Data

The game loads data from YAML files in the `data` directory:
- World data: `data/worlds/test_world.yaml`
- Monster data: `data/monsters.yaml`
- Player data: `data/test_player.yaml`
- Armor data: `data/armor.yml`
- Weapons data: `data/weapons.yml`

To use custom data files, you would need to modify the paths in the source code (specifically in `game/bubbletea.go`) and rebuild the application.