# VkScreenBufferPropertiesQNX(3) Manual Page

## Name

VkScreenBufferPropertiesQNX - Properties of External Memory QNX Screen Buffers



## [](#_c_specification)C Specification

The `VkScreenBufferPropertiesQNX` structure returned is defined as:

```c++
// Provided by VK_QNX_external_memory_screen_buffer
typedef struct VkScreenBufferPropertiesQNX {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       allocationSize;
    uint32_t           memoryTypeBits;
} VkScreenBufferPropertiesQNX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `allocationSize` is the size of the external memory.
- `memoryTypeBits` is a bitmask containing one bit set for every memory type which the specified Screen buffer **can** be imported as.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkScreenBufferPropertiesQNX-sType-sType)VUID-VkScreenBufferPropertiesQNX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SCREEN_BUFFER_PROPERTIES_QNX`
- [](#VUID-VkScreenBufferPropertiesQNX-pNext-pNext)VUID-VkScreenBufferPropertiesQNX-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferFormatPropertiesQNX.html)
- [](#VUID-VkScreenBufferPropertiesQNX-sType-unique)VUID-VkScreenBufferPropertiesQNX-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_QNX\_external\_memory\_screen\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QNX_external_memory_screen_buffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetScreenBufferPropertiesQNX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkScreenBufferPropertiesQNX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0