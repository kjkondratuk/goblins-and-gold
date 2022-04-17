# goblins-and-gold

## What is this?
A text-based adventure game loosely based on popular tabletop RPG systems.
This project is very much in its infancy and has its fair share of spaghetti
code already as this idea is still taking shape.

## Objectives
1. Worlds/scenarios should be easy to create
2. Worlds/scenarios should be easy to share
3. The world should generally feel open and explorable to the player
4. The command-line user interface should be simple and should not involve typing
long commands in order to do simple tasks
5. Longer commands should occur interactively to give the user as much contextual
information about their interaction as possible
6. Leverage command-line completion and command aliases to simplify user input

## What's Working?

In its current state, this engine supports the following features:

1. Reads world and player data from YAMl files
2. A player can get a rough sense of their surroundings via room descriptions
3. A player can traverse the world graph between rooms utilizing paths
4. The world can be populated with containers that players can open
5. Openable containers can contain arbitrary collections of loot items

## What's Next?

The following are features that are on the roadmap in order of percieved priority: 

1. Skill Checks
2. Equipment
3. Traps
4. Monsters
5. Combat
6. NPC Interactions
7. Saving game state

## What's After that?

The following are some additional goals, but I'm not too sure when I'll get around to
actually implementing these:

1. A graphical tool to build worlds
2. Multiplayer

## How does it work?

* [Interactions](./interaction/interactions.md)

## How do I make maps?

Check out the [test map](./data/worlds/test_world.md) ([Source](./data/worlds/test_world.yaml))