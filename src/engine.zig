const std = @import("std");
const c = @cImport({
    @cInclude("GLFW/glfw3.h");
});

pub const Sprite = struct {
    x: f32,
    y: f32,
    width: f32,
    height: f32,
    // In a real engine, you'd have a texture handle here
};

pub fn initWindow(width: i32, height: i32, title: [*c]const u8) ?*c.GLFWwindow {
    if (c.glfwInit() == 0) return null;
    c.glfwWindowHint(c.GLFW_CLIENT_API, c.GLFW_NO_API); // Vulkan only
    const window = c.glfwCreateWindow(width, height, title, null, null);
    if (window == null) {
        c.glfwTerminate();
        return null;
    }
    return window;
}

pub fn pollInput(window: *c.GLFWwindow) bool {
    c.glfwPollEvents();
    return c.glfwWindowShouldClose(window) == 0;
}

// Placeholder for Vulkan setup and sprite rendering
pub fn renderSprite(sprite: Sprite) void {
    // In a real engine, you'd record Vulkan commands here
    _ = sprite;
}

pub fn deinitWindow(window: *c.GLFWwindow) void {
    c.glfwDestroyWindow(window);
    c.glfwTerminate();
}
