# VkBindIndexBufferIndirectCommandNV(3) Manual Page

## Name

VkBindIndexBufferIndirectCommandNV - Structure specifying input data for
a single index buffer command token



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindIndexBufferIndirectCommandNV` structure specifies the input
data for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_INDEX_BUFFER_NV` token.

``` c
// Provided by VK_NV_device_generated_commands
typedef struct VkBindIndexBufferIndirectCommandNV {
    VkDeviceAddress    bufferAddress;
    uint32_t           size;
    VkIndexType        indexType;
} VkBindIndexBufferIndirectCommandNV;
```

## <a href="#_members" class="anchor"></a>Members

- `bufferAddress` specifies a physical address of the
  [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) used as index buffer.

- `size` is the byte size range which is available for this operation
  from the provided address.

- `indexType` is a [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) value specifying how
  indices are treated. Instead of the Vulkan enum values, a custom
  `uint32_t` value **can** be mapped to [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html)
  by specifying the `VkIndirectCommandsLayoutTokenNV`::`pIndexTypes` and
  `VkIndirectCommandsLayoutTokenNV`::`pIndexTypeValues` arrays.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBindIndexBufferIndirectCommandNV-None-02946"
  id="VUID-VkBindIndexBufferIndirectCommandNV-None-02946"></a>
  VUID-VkBindIndexBufferIndirectCommandNV-None-02946  
  The buffer’s usage flag from which the address was acquired **must**
  have the `VK_BUFFER_USAGE_INDEX_BUFFER_BIT` bit set

- <a href="#VUID-VkBindIndexBufferIndirectCommandNV-bufferAddress-02947"
  id="VUID-VkBindIndexBufferIndirectCommandNV-bufferAddress-02947"></a>
  VUID-VkBindIndexBufferIndirectCommandNV-bufferAddress-02947  
  The `bufferAddress` **must** be aligned to the `indexType` used

- <a href="#VUID-VkBindIndexBufferIndirectCommandNV-None-02948"
  id="VUID-VkBindIndexBufferIndirectCommandNV-None-02948"></a>
  VUID-VkBindIndexBufferIndirectCommandNV-None-02948  
  Each element of the buffer from which the address was acquired and
  that is non-sparse **must** be bound completely and contiguously to a
  single `VkDeviceMemory` object

Valid Usage (Implicit)

- <a href="#VUID-VkBindIndexBufferIndirectCommandNV-indexType-parameter"
  id="VUID-VkBindIndexBufferIndirectCommandNV-indexType-parameter"></a>
  VUID-VkBindIndexBufferIndirectCommandNV-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindIndexBufferIndirectCommandNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
