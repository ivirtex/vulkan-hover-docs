# VkBindVertexBufferIndirectCommandNV(3) Manual Page

## Name

VkBindVertexBufferIndirectCommandNV - Structure specifying input data for a single vertex buffer command token



## [](#_c_specification)C Specification

The `VkBindVertexBufferIndirectCommandNV` structure specifies the input data for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_VERTEX_BUFFER_NV` token.

```c++
// Provided by VK_NV_device_generated_commands
typedef struct VkBindVertexBufferIndirectCommandNV {
    VkDeviceAddress    bufferAddress;
    uint32_t           size;
    uint32_t           stride;
} VkBindVertexBufferIndirectCommandNV;
```

## [](#_members)Members

- `bufferAddress` specifies a physical address of the [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) used as vertex input binding.
- `size` is the byte size range which is available for this operation from the provided address.
- `stride` is the byte size stride for this vertex input binding as in `VkVertexInputBindingDescription`::`stride`. It is only used if `VkIndirectCommandsLayoutTokenNV`::`vertexDynamicStride` was set, otherwise the stride is inherited from the current bound graphics pipeline.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBindVertexBufferIndirectCommandNV-None-02949)VUID-VkBindVertexBufferIndirectCommandNV-None-02949  
  The buffer’s usage flag from which the address was acquired **must** have the `VK_BUFFER_USAGE_VERTEX_BUFFER_BIT` bit set

Valid Usage (Implicit)

- [](#VUID-VkBindVertexBufferIndirectCommandNV-bufferAddress-parameter)VUID-VkBindVertexBufferIndirectCommandNV-bufferAddress-parameter  
  `bufferAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindVertexBufferIndirectCommandNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0