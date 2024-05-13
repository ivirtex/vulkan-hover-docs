# VkBufferCopy2(3) Manual Page

## Name

VkBufferCopy2 - Structure specifying a buffer copy operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferCopy2` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkBufferCopy2 {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceSize       srcOffset;
    VkDeviceSize       dstOffset;
    VkDeviceSize       size;
} VkBufferCopy2;
```

or the equivalent

``` c
// Provided by VK_KHR_copy_commands2
typedef VkBufferCopy2 VkBufferCopy2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcOffset` is the starting offset in bytes from the start of
  `srcBuffer`.

- `dstOffset` is the starting offset in bytes from the start of
  `dstBuffer`.

- `size` is the number of bytes to copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBufferCopy2-size-01988"
  id="VUID-VkBufferCopy2-size-01988"></a>
  VUID-VkBufferCopy2-size-01988  
  The `size` **must** be greater than `0`

Valid Usage (Implicit)

- <a href="#VUID-VkBufferCopy2-sType-sType"
  id="VUID-VkBufferCopy2-sType-sType"></a>
  VUID-VkBufferCopy2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_COPY_2`

- <a href="#VUID-VkBufferCopy2-pNext-pNext"
  id="VUID-VkBufferCopy2-pNext-pNext"></a>
  VUID-VkBufferCopy2-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkCopyBufferInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyBufferInfo2.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferCopy2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
