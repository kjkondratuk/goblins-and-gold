# Bubbletea UI Migration

This document describes the migration of the Goblins and Gold game from using a combination of go-prompt, promptui, and p-term to using charmbracelet/bubbletea and related components from charmbracelet/bubbles.

## Architecture Overview

The new UI architecture is based on the Bubble Tea framework, which follows the Model-View-Update (MVU) pattern. The application is structured as a single bubbletea application with multiple components displayed within a view.

### Components

1. **Main Model** (`ui/model.go`)
   - Manages the application state
   - Coordinates the different UI components
   - Handles layout and rendering

2. **Room View** (`ui/room_view.go`)
   - Displays the current room description
   - Shows the room name and description

3. **Command View** (`ui/command_view.go`)
   - Provides a command prompt for user input
   - Executes commands when the user presses Enter

4. **Message View** (`ui/message_view.go`)
   - Displays system messages and logs
   - Shows different message types with appropriate styling

5. **Logger** (`ui/logger.go`)
   - Redirects pterm logs to the message view
   - Handles different log levels

6. **Prompt Library** (`ui/prompt_lib.go`)
   - Implements the ux.PromptLib interface using bubbletea
   - Provides methods for text input and selection menus

### Layout

The UI is divided into three sections vertically:
1. Room description (top half)
2. Message box (quarter of the screen)
3. Command prompt (remaining space)

This layout provides a clear separation of concerns while maintaining a cohesive user experience.

## Changes Made

1. Added dependencies to `go.mod`:
   - github.com/charmbracelet/bubbletea
   - github.com/charmbracelet/bubbles
   - github.com/charmbracelet/lipgloss

2. Created a new UI package with components:
   - model.go - Main bubbletea model
   - room_view.go - Room description component
   - command_view.go - Command prompt component
   - message_view.go - Message box component
   - logger.go - Logger for redirecting pterm logs
   - prompt_lib.go - Implementation of PromptLib using bubbletea

3. Created a new game initialization function in `game/bubbletea.go`:
   - RunWithBubbleTea - Initializes and runs the game with the bubbletea UI

4. Updated `main.go` to use the new RunWithBubbleTea function

## Benefits

1. **Unified UI Framework**: The application now uses a single UI framework (bubbletea) instead of a combination of different libraries.

2. **Improved User Experience**: The UI is more cohesive and responsive, with a clear layout and consistent styling.

3. **Better Message Handling**: System messages are now displayed in a dedicated message box, making them easier to read and track.

4. **Extensibility**: The component-based architecture makes it easier to add new UI elements or modify existing ones.

5. **Maintainability**: The code is more organized and follows bubbletea best practices, making it easier to maintain and extend.

## Future Improvements

1. **Custom Styles**: Add more custom styling to improve the visual appeal of the UI.

2. **Keyboard Shortcuts**: Implement keyboard shortcuts for common actions.

3. **Help System**: Add a help system that displays available commands and their usage.

4. **Multiple Layouts**: Implement different layouts for different game states (e.g., combat, inventory management).

5. **Animations**: Add animations for transitions between different game states.