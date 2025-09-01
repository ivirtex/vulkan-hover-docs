# VkIndirectCommandsVertexBufferTokenEXT(3) Manual Page

## Name

VkIndirectCommandsVertexBufferTokenEXT - Structure specifying layout token info for a single index buffer command token



## [](#_c_specification)C Specification

The `VkIndirectCommandsVertexBufferTokenEXT` structure specifies the layout token info for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_EXT` token.

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkIndirectCommandsVertexBufferTokenEXT {
    uint32_t    vertexBindingUnit;
} VkIndirectCommandsVertexBufferTokenEXT;
```

## [](#_members)Members

- `vertexBindingUnit` is the vertex input binding number to be bound.

## [](#_description)Description

Valid Usage

- [](#VUID-VkIndirectCommandsVertexBufferTokenEXT-vertexBindingUnit-11134)VUID-VkIndirectCommandsVertexBufferTokenEXT-vertexBindingUnit-11134  
  `vertexBindingUnit` **must** be less than the total number of vertex input bindings in use by the current graphics state

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsTokenDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenDataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsVertexBufferTokenEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0