# Game Configuration Guide

This document describes the configuration options for Goblins and Gold, including how to define items, monsters, players, and worlds.

## Configuration Files

Goblins and Gold uses YAML files for configuration. The main configuration files are:

- `data/worlds/test_world.yaml`: Defines the world structure, rooms, paths, and containers
- `data/monsters.yaml`: Defines monster types and their attributes
- `data/test_player.yaml`: Defines the player character's attributes
- `data/armor.yml`: Defines armor types and their properties
- `data/weapons.yml`: Defines weapon types (currently a placeholder)

## World Configuration

Worlds are defined in YAML files like `data/worlds/test_world.yaml`. A world consists of:

### Common Descriptors

Reusable text descriptions that can be referenced throughout the world definition:

```yaml
common_descriptors:
  stone_tunnel: &stone_tunnel This is an empty room with a compacted dirt floor and stone walls.
  simple_door: &simple_door A simple wooden door hangs in the doorway.
```

### Rooms

Rooms are the main locations in the world. Each room has:

- `name`: The room's name
- `description`: A description of the room
- `paths`: Connections to other rooms
- `containers`: Items that can be interacted with
- `mandatory_encounters`: Encounters that trigger when entering the room

Example:

```yaml
rooms:
  treasure-room-1:
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
    paths:
      - room: starting-room
        description: A fancy door with gold filigree.
```

### Containers

Containers can be placed in rooms and can hold items. Container properties include:

- `type`: The type of container (e.g., "Chest")
- `description`: A description of the container
- `interactions`: Available interactions (e.g., "Open", "Loot", "Unlock")
- `locked`: Optional lock properties if the container is locked
- `items`: Items contained within

Example of a locked container:

```yaml
- type: Chest
  interactions: [Open, Loot, Unlock]
  locked:
    dc: 10
    type: Dex
    crit_success:
      description: You deftly pick the lock and reveal the items hidden within.
    success:
      description: You are able to pick the lock, but attack it beyond repair.
    fail:
      description: You fail to pick the lock. The chest remains shut.
    crit_fail:
      damage:
        hp: 1
        type: Slashing
      description: You fail to pick the lock and hurt your finger in the process.
  description: A strongly built wooden chest with a strong iron lock.
  items:
    - type: Gold
      description: A pile of gold coins
      quantity: 150
      unit: Coins
```

### Starting Room

The `startingRoom` property defines which room the player starts in:

```yaml
startingRoom: starting-room
```

## Monster Configuration

Monsters are defined in `data/monsters.yaml`. Each monster has:

- `name`: The monster's name
- `description`: A description of the monster
- `ac`: Armor Class
- `hp`: Hit Points
- `stats`: Ability scores (str, dex, con, int, wis, cha)
- `inventory`: Items the monster carries
- `attacks`: Attacks the monster can perform

Example:

```yaml
goblin: &goblin
  name: Goblin
  ac: 15
  hp: 7
  description: A short humanoid, in rags, wielding a scimitar
  stats:
    str: 8
    dex: 14
    con: 10
    int: 10
    wis: 8
    cha: 8
  inventory:
    - type: Armor
      description: A rough-hewn suit of leather armor
      equip_info:
        slot: Torso
        ac_bonus: 2
      quantity: 1
      unit: suit
  attacks:
    scimitar:
      bonus: 4
      range: 5
      damage:
        - roll: 1d6
          bonus: 2
          damage_type: Slashing
```

## Player Configuration

Players are defined in `data/test_player.yaml`. A player has:

- `name`: The player's name
- `hp`: Hit Points
- `ac`: Armor Class
- `stats`: Ability scores (str, dex, con, int, wis, cha) and level
- `loadout`: Equipment the player has equipped
- `attacks`: Attacks the player can perform

Example:

```yaml
name: Arthur Sotheby
hp: 57
ac: 15
stats:
  lvl: 1
  str: 15
  dex: 12
  con: 16
  int: 8
  cha: 11
  wis: 10
loadout:
  torso:
    type: Armor
    equip_info:
      slot: Torso
      ac_bonus: 2
    description: Breastplate
    quantity: 1
    unit: piece
attacks:
  dagger:
    bonus: 3
    range: 5
    damage:
      - roll: 1d4
        bonus: 3
        damage_type: Piercing
```

## Armor Configuration

Armor types are defined in `data/armor.yml`. Each armor type has:

- `name`: The armor's name
- `attr`: The attribute that affects the armor's effectiveness (Dex or Str)
- `min`/`max`: Minimum or maximum attribute values that affect the armor
- `ac`: The base Armor Class provided by the armor
- `stealth`: Whether the armor imposes disadvantage on stealth checks
- `properties`: Categories of the armor (Light, Medium, Heavy)

Example:

```yaml
- name: "Chain Mail"
  attr: "Str"
  min: 13
  ac: 16
  stealth: false
  properties:
    - "Heavy"
```

## Weapon Configuration

Weapons can be defined in various files, including `data/weapons.yml` (currently a placeholder) and within monster definitions. A weapon typically has:

- `bonus`: Attack bonus
- `range`: Attack range
- `damage`: Damage properties
  - `roll`: Dice roll for damage (e.g., "1d6")
  - `bonus`: Damage bonus
  - `damage_type`: Type of damage (e.g., "Slashing", "Piercing")

Example:

```yaml
scimitar:
  bonus: 4
  range: 5
  damage:
    - roll: 1d6
      bonus: 2
      damage_type: Slashing
```