# VkBindIndexBufferIndirectCommandEXT(3) Manual Page

## Name

VkBindIndexBufferIndirectCommandEXT - Structure specifying input data for a single index buffer command token



## [](#_c_specification)C Specification

The `VkBindIndexBufferIndirectCommandEXT` structure specifies the input data for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_EXT` token.

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkBindIndexBufferIndirectCommandEXT {
    VkDeviceAddress    bufferAddress;
    uint32_t           size;
    VkIndexType        indexType;
} VkBindIndexBufferIndirectCommandEXT;
```

## [](#_members)Members

- `bufferAddress` specifies a physical address of the [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) used as index buffer.
- `size` is the byte size range which is available for this operation from the provided address.
- `indexType` is a [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value specifying how indices are treated.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBindIndexBufferIndirectCommandEXT-None-11117)VUID-VkBindIndexBufferIndirectCommandEXT-None-11117  
  The buffer’s usage flags from which the address was acquired **must** have the `VK_BUFFER_USAGE_INDEX_BUFFER_BIT` bit set
- [](#VUID-VkBindIndexBufferIndirectCommandEXT-bufferAddress-11118)VUID-VkBindIndexBufferIndirectCommandEXT-bufferAddress-11118  
  The `bufferAddress` **must** be aligned to the [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) of the `indexType` used

Valid Usage (Implicit)

- [](#VUID-VkBindIndexBufferIndirectCommandEXT-bufferAddress-parameter)VUID-VkBindIndexBufferIndirectCommandEXT-bufferAddress-parameter  
  `bufferAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkBindIndexBufferIndirectCommandEXT-indexType-parameter)VUID-VkBindIndexBufferIndirectCommandEXT-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindIndexBufferIndirectCommandEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0