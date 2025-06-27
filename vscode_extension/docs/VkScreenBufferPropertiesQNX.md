# VkScreenBufferPropertiesQNX(3) Manual Page

## Name

VkScreenBufferPropertiesQNX - Properties of External Memory QNX Screen
Buffers



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkScreenBufferPropertiesQNX` structure returned is defined as:

``` c
// Provided by VK_QNX_external_memory_screen_buffer
typedef struct VkScreenBufferPropertiesQNX {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       allocationSize;
    uint32_t           memoryTypeBits;
} VkScreenBufferPropertiesQNX;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `allocationSize` is the size of the external memory.

- `memoryTypeBits` is a bitmask containing one bit set for every memory
  type which the specified Screen buffer **can** be imported as.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkScreenBufferPropertiesQNX-sType-sType"
  id="VUID-VkScreenBufferPropertiesQNX-sType-sType"></a>
  VUID-VkScreenBufferPropertiesQNX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SCREEN_BUFFER_PROPERTIES_QNX`

- <a href="#VUID-VkScreenBufferPropertiesQNX-pNext-pNext"
  id="VUID-VkScreenBufferPropertiesQNX-pNext-pNext"></a>
  VUID-VkScreenBufferPropertiesQNX-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferFormatPropertiesQNX.html)

- <a href="#VUID-VkScreenBufferPropertiesQNX-sType-unique"
  id="VUID-VkScreenBufferPropertiesQNX-sType-unique"></a>
  VUID-VkScreenBufferPropertiesQNX-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QNX_external_memory_screen_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QNX_external_memory_screen_buffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetScreenBufferPropertiesQNX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkScreenBufferPropertiesQNX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
