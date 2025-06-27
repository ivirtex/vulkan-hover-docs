# VkBufferCopy(3) Manual Page

## Name

VkBufferCopy - Structure specifying a buffer copy operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferCopy` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkBufferCopy {
    VkDeviceSize    srcOffset;
    VkDeviceSize    dstOffset;
    VkDeviceSize    size;
} VkBufferCopy;
```

## <a href="#_members" class="anchor"></a>Members

- `srcOffset` is the starting offset in bytes from the start of
  `srcBuffer`.

- `dstOffset` is the starting offset in bytes from the start of
  `dstBuffer`.

- `size` is the number of bytes to copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBufferCopy-size-01988"
  id="VUID-VkBufferCopy-size-01988"></a> VUID-VkBufferCopy-size-01988  
  The `size` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[vkCmdCopyBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferCopy"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
