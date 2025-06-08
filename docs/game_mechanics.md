# Game Mechanics

This document describes the core mechanics of Goblins and Gold, including interactions, combat, and navigation.

## Command System

Goblins and Gold uses a command-based interface. The main commands include:

- `go`: Navigate to another room
- `look`: Examine your surroundings or a specific object
- `interact`: Interact with objects in the environment
- `stats`: Display your character's statistics
- `quit`: Exit the game

In debug mode (enabled with the `DEBUG=true` environment variable), additional debug commands are available.

## Navigation

The world in Goblins and Gold is structured as a graph of interconnected rooms. Each room can have multiple paths leading to other rooms. To navigate between rooms:

1. Use the `look` command to see available paths
2. Use the `go` command followed by the direction or room name to move

Example:
```
> look
You are in the Starting Room. There are paths leading to:
- Dead End Room (through an empty archway)
- Treasure Room (through a fancy door with gold filigree)
- Dirt Hall (through a simple wooden door)

> go treasure room
You move through the fancy door with gold filigree and enter the Treasure Room.
```

## Interactions

### Container Interactions

Containers (like chests) can be interacted with using the `interact` command. The available interactions depend on the container and may include:

- **Open**: View the contents of a container
- **Loot**: Take items from a container
- **Unlock**: Attempt to unlock a locked container

#### Opening Containers

```
> look
You see a small wooden chest against the wall.

> interact open chest
You open the chest and see:
- A few glittering coins (5 Coins)
```

#### Looting Containers

```
> interact loot chest
You take the few glittering coins (5 Coins).
```

#### Unlocking Containers

Locked containers require an unlock attempt, which involves a skill check:

```
> interact unlock chest
The chest is locked. You attempt to pick the lock...
[Dexterity check: 12 vs DC 10]
Success! You are able to pick the lock, but attack it beyond repair.
```

The outcome of unlock attempts depends on the roll result:
- **Critical Success**: Unlock with additional benefits
- **Success**: Unlock successfully
- **Failure**: Fail to unlock
- **Critical Failure**: Fail with negative consequences (e.g., taking damage)

## Skill Checks

Skill checks are used for various interactions, like unlocking containers. A skill check involves:

1. Rolling a die (typically a d20)
2. Adding the relevant ability modifier
3. Comparing the result to a Difficulty Class (DC)

The outcome depends on how the roll compares to the DC:
- Roll ≥ DC + 5: Critical Success
- Roll ≥ DC: Success
- Roll < DC: Failure
- Roll ≤ DC - 5: Critical Failure

## Combat (Planned Feature)

Combat is a planned feature that will involve:

1. Initiative determination
2. Turn-based actions
3. Attack rolls and damage calculation
4. Special abilities and effects

Monsters are defined with various attacks and statistics that will be used in combat encounters.

## Equipment (Planned Feature)

Equipment management is a planned feature that will allow players to:

1. Equip weapons and armor
2. Manage inventory
3. Benefit from equipment bonuses

## Sequence Diagrams

For detailed interaction flows, see the [Interactions](../interaction/interactions.md) document, which includes sequence diagrams for:

- Opening containers
- Looting containers
- Unlocking containers

## Future Mechanics

The following mechanics are planned for future development:

1. Skill Checks (partially implemented)
2. Equipment (partially implemented)
3. Traps
4. Monsters (partially implemented)
5. Combat
6. NPC Interactions
7. Saving game state