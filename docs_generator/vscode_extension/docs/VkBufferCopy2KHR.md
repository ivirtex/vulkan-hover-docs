# VkBufferCopy2(3) Manual Page

## Name

VkBufferCopy2 - Structure specifying a buffer copy operation



## [](#_c_specification)C Specification

The `VkBufferCopy2` structure is defined as:

```c++
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

```c++
// Provided by VK_KHR_copy_commands2
typedef VkBufferCopy2 VkBufferCopy2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcOffset` is the starting offset in bytes from the start of `srcBuffer`.
- `dstOffset` is the starting offset in bytes from the start of `dstBuffer`.
- `size` is the number of bytes to copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBufferCopy2-size-01988)VUID-VkBufferCopy2-size-01988  
  The `size` **must** be greater than `0`

Valid Usage (Implicit)

- [](#VUID-VkBufferCopy2-sType-sType)VUID-VkBufferCopy2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_COPY_2`
- [](#VUID-VkBufferCopy2-pNext-pNext)VUID-VkBufferCopy2-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCopyBufferInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferInfo2.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferCopy2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0