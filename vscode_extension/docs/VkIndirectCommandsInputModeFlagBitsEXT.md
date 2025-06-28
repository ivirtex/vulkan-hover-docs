# VkIndirectCommandsInputModeFlagBitsEXT(3) Manual Page

## Name

VkIndirectCommandsInputModeFlagBitsEXT - Bitmask specifying allowed usage of an indirect commands layout



## [](#_c_specification)C Specification

Bits which are set in [VkIndirectCommandsIndexBufferTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsIndexBufferTokenEXT.html)::`mode`, specifying how an index buffer is used, are:

```c++
// Provided by VK_EXT_device_generated_commands
typedef enum VkIndirectCommandsInputModeFlagBitsEXT {
    VK_INDIRECT_COMMANDS_INPUT_MODE_VULKAN_INDEX_BUFFER_EXT = 0x00000001,
    VK_INDIRECT_COMMANDS_INPUT_MODE_DXGI_INDEX_BUFFER_EXT = 0x00000002,
} VkIndirectCommandsInputModeFlagBitsEXT;
```

## [](#_description)Description

- `VK_INDIRECT_COMMANDS_INPUT_MODE_VULKAN_INDEX_BUFFER_EXT` specifies that the indirect buffer contains [VkBindIndexBufferIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindIndexBufferIndirectCommandEXT.html).
- `VK_INDIRECT_COMMANDS_INPUT_MODE_DXGI_INDEX_BUFFER_EXT` specifies that the indirect buffer contains `D3D12_INDEX_BUFFER_VIEW`.

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsIndexBufferTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsIndexBufferTokenEXT.html), [VkIndirectCommandsInputModeFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsInputModeFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsInputModeFlagBitsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0