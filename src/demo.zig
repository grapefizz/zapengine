const std = @import("std");
const engine = @import("engine.zig");

pub fn main() !void {
    const window = engine.initWindow(800, 600, "Zig Vulkan 2D Engine");
    if (window == null) {
        std.debug.print("Failed to create window\n", .{});
        return;
    }

    const sprite = engine.Sprite{
        .x = 100,
        .y = 100,
        .width = 64,
        .height = 64,
    };

    while (engine.pollInput(window.?)) {
        // Here you would update sprite position based on input
        engine.renderSprite(sprite);
        // In a real engine, you'd present the frame here
    }

    engine.deinitWindow(window.?);
}
