const std = @import("std");
const engine = @import("engine.zig");

pub fn main() !void {
    const window = engine.initWindow(800, 600, "Zig Vulkan 2D Engine");
    if (window == null) {
        std.debug.print("Failed to create window\n", .{});
        return;
    }

    var sprite = engine.Sprite{
        .x = 100,
        .y = 100,
        .width = 64,
        .height = 64,
    };

    while (engine.pollInput(window.?)) {
        engine.moveSprite(&sprite, window.?, 2.0);
        engine.renderSprite(sprite);
    }

    engine.deinitWindow(window.?);
}
