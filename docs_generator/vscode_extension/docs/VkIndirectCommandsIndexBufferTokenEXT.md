# VkIndirectCommandsIndexBufferTokenEXT(3) Manual Page

## Name

VkIndirectCommandsIndexBufferTokenEXT - Structure specifying layout token info for a single index buffer command token



## [](#_c_specification)C Specification

The `VkIndirectCommandsIndexBufferTokenEXT` structure specifies the layout token info for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_EXT` token.

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkIndirectCommandsIndexBufferTokenEXT {
    VkIndirectCommandsInputModeFlagBitsEXT    mode;
} VkIndirectCommandsIndexBufferTokenEXT;
```

## [](#_members)Members

- `mode` specifies the mode to use with this token.

## [](#_description)Description

This allows for easy layering of Vulkan atop other APIs. When `VK_INDIRECT_COMMANDS_INPUT_MODE_DXGI_INDEX_BUFFER_EXT` is specified, the indirect buffer can contain a `D3D12_INDEX_BUFFER_VIEW` instead of [VkBindIndexBufferIndirectCommandEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindIndexBufferIndirectCommandEXT.html) as D3D’s DXGI format value is mapped to the [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html). It works as both structs are otherwise binary compatible.

Valid Usage

- [](#VUID-VkIndirectCommandsIndexBufferTokenEXT-mode-11135)VUID-VkIndirectCommandsIndexBufferTokenEXT-mode-11135  
  `mode` **must** be non-zero
- [](#VUID-VkIndirectCommandsIndexBufferTokenEXT-mode-11136)VUID-VkIndirectCommandsIndexBufferTokenEXT-mode-11136  
  `mode` **must** be one of the bits supported in [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-supportedIndirectCommandsInputModes)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`supportedIndirectCommandsInputModes`

Valid Usage (Implicit)

- [](#VUID-VkIndirectCommandsIndexBufferTokenEXT-mode-parameter)VUID-VkIndirectCommandsIndexBufferTokenEXT-mode-parameter  
  `mode` **must** be a valid [VkIndirectCommandsInputModeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsInputModeFlagBitsEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsInputModeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsInputModeFlagBitsEXT.html), [VkIndirectCommandsTokenDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenDataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsIndexBufferTokenEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0