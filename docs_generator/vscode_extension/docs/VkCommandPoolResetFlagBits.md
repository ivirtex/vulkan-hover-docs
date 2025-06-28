# VkCommandPoolResetFlagBits(3) Manual Page

## Name

VkCommandPoolResetFlagBits - Bitmask controlling behavior of a command pool reset



## [](#_c_specification)C Specification

Bits which **can** be set in [vkResetCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetCommandPool.html)::`flags`, controlling the reset operation, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkCommandPoolResetFlagBits {
    VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT = 0x00000001,
} VkCommandPoolResetFlagBits;
```

## [](#_description)Description

- `VK_COMMAND_POOL_RESET_RELEASE_RESOURCES_BIT` specifies that resetting a command pool recycles all of the resources from the command pool back to the system.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandPoolResetFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolResetFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCommandPoolResetFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0