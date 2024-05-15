# Vulkan Hover Docs

This extension provides hover documentation for Vulkan API based on the official Khronos Group man pages for Vulkan.

## Features

![preview](https://github.com/ivirtex/vulkan-hover-docs/blob/master/vscode_ext/images/screenshot.png?raw=true)

## Known Issues

- Doesn't work with Vulkan C++ bindings due to naming differences, which are not easily detectable and can resolve to multiple Vulkan functions e.g.
`commandBuffer.buildAccelerationStructuresKHR` could be matched to `vkCmdBuildAccelerationStructuresKHR` or `vkBuildAccelerationStructuresKHR`.
