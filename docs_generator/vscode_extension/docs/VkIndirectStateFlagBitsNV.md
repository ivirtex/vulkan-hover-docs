# VkIndirectStateFlagBitsNV(3) Manual Page

## Name

VkIndirectStateFlagBitsNV - Bitmask specifying state that can be altered on the device



## [](#_c_specification)C Specification

A subset of the graphics pipeline state **can** be altered using indirect state flags:

```c++
// Provided by VK_NV_device_generated_commands
typedef enum VkIndirectStateFlagBitsNV {
    VK_INDIRECT_STATE_FLAG_FRONTFACE_BIT_NV = 0x00000001,
} VkIndirectStateFlagBitsNV;
```

## [](#_description)Description

- `VK_INDIRECT_STATE_FLAG_FRONTFACE_BIT_NV` allows to toggle the [VkFrontFace](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrontFace.html) rasterization state for subsequent drawing commands.

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkIndirectStateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectStateFlagsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectStateFlagBitsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0