# VkBindIndexBufferIndirectCommandNV(3) Manual Page

## Name

VkBindIndexBufferIndirectCommandNV - Structure specifying input data for a single index buffer command token



## [](#_c_specification)C Specification

The `VkBindIndexBufferIndirectCommandNV` structure specifies the input data for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_NV` token.

```c++
// Provided by VK_NV_device_generated_commands
typedef struct VkBindIndexBufferIndirectCommandNV {
    VkDeviceAddress    bufferAddress;
    uint32_t           size;
    VkIndexType        indexType;
} VkBindIndexBufferIndirectCommandNV;
```

## [](#_members)Members

- `bufferAddress` specifies a physical address of the [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) used as index buffer.
- `size` is the byte size range which is available for this operation from the provided address.
- `indexType` is a [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value specifying how indices are treated. Instead of the Vulkan enum values, a custom `uint32_t` value **can** be mapped to [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) by specifying the `VkIndirectCommandsLayoutTokenNV`::`pIndexTypes` and `VkIndirectCommandsLayoutTokenNV`::`pIndexTypeValues` arrays.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBindIndexBufferIndirectCommandNV-None-02946)VUID-VkBindIndexBufferIndirectCommandNV-None-02946  
  The buffer’s usage flag from which the address was acquired **must** have the `VK_BUFFER_USAGE_INDEX_BUFFER_BIT` bit set
- [](#VUID-VkBindIndexBufferIndirectCommandNV-bufferAddress-02947)VUID-VkBindIndexBufferIndirectCommandNV-bufferAddress-02947  
  The `bufferAddress` **must** be aligned to the `indexType` used
- [](#VUID-VkBindIndexBufferIndirectCommandNV-None-02948)VUID-VkBindIndexBufferIndirectCommandNV-None-02948  
  Each element of the buffer from which the address was acquired and that is non-sparse **must** be bound completely and contiguously to a single `VkDeviceMemory` object

Valid Usage (Implicit)

- [](#VUID-VkBindIndexBufferIndirectCommandNV-indexType-parameter)VUID-VkBindIndexBufferIndirectCommandNV-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindIndexBufferIndirectCommandNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0