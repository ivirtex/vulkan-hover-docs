# VkCommandBufferLevel(3) Manual Page

## Name

VkCommandBufferLevel - Enumerant specifying a command buffer level



## [](#_c_specification)C Specification

Possible values of [VkCommandBufferAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferAllocateInfo.html)::`level`, specifying the command buffer level, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkCommandBufferLevel {
    VK_COMMAND_BUFFER_LEVEL_PRIMARY = 0,
    VK_COMMAND_BUFFER_LEVEL_SECONDARY = 1,
} VkCommandBufferLevel;
```

## [](#_description)Description

- `VK_COMMAND_BUFFER_LEVEL_PRIMARY` specifies a primary command buffer.
- `VK_COMMAND_BUFFER_LEVEL_SECONDARY` specifies a secondary command buffer.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBufferAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferAllocateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCommandBufferLevel)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0