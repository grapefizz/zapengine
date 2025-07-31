# ZapEngine

ZapEngine is a simple 2D game engine written in Zig, using Vulkan for rendering and GLFW for windowing and input.

## Features
- Basic window creation with GLFW
- Vulkan (MoltenVK) rendering backend (scaffolded, not yet drawing)
- Sprite struct and input polling (expandable)

## Requirements
- Zig 0.14.1 or newer
- Vulkan SDK (LunarG) for macOS
- MoltenVK (see below)
- GLFW (install via Homebrew: `brew install glfw`)

## Setup
1. **Install the Vulkan SDK** from [LunarG](https://vulkan.lunarg.com/sdk/home).
2. **Install MoltenVK**:
   - Download from [MoltenVK Releases](https://github.com/KhronosGroup/MoltenVK/releases)
   - Copy `libMoltenVK.dylib` to `/usr/local/lib/`
3. **Install GLFW**:
   - `brew install glfw`
4. **Clone this repository** and build:
   ```sh
   zig build run-demo
   ```

## Project Structure
- `src/engine.zig` — Core engine code (window, input, sprite, rendering stub)
- `src/demo.zig` — Example usage and main loop
- `build.zig` — Build script for Zig

## Notes
- The current version only opens a window. Vulkan rendering is not yet implemented.
- To add rendering, you will need to implement Vulkan instance, device, swapchain, and drawing code in `engine.zig`.

## License
MIT
