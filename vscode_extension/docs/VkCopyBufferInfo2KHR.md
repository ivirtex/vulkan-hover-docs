# VkCopyBufferInfo2(3) Manual Page

## Name

VkCopyBufferInfo2 - Structure specifying parameters of a buffer copy command



## [](#_c_specification)C Specification

The `VkCopyBufferInfo2` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkCopyBufferInfo2 {
    VkStructureType         sType;
    const void*             pNext;
    VkBuffer                srcBuffer;
    VkBuffer                dstBuffer;
    uint32_t                regionCount;
    const VkBufferCopy2*    pRegions;
} VkCopyBufferInfo2;
```

or the equivalent

```c++
// Provided by VK_KHR_copy_commands2
typedef VkCopyBufferInfo2 VkCopyBufferInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcBuffer` is the source buffer.
- `dstBuffer` is the destination buffer.
- `regionCount` is the number of regions to copy.
- `pRegions` is a pointer to an array of [VkBufferCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCopy2.html) structures specifying the regions to copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyBufferInfo2-srcOffset-00113)VUID-VkCopyBufferInfo2-srcOffset-00113  
  The `srcOffset` member of each element of `pRegions` **must** be less than the size of `srcBuffer`
- [](#VUID-VkCopyBufferInfo2-dstOffset-00114)VUID-VkCopyBufferInfo2-dstOffset-00114  
  The `dstOffset` member of each element of `pRegions` **must** be less than the size of `dstBuffer`
- [](#VUID-VkCopyBufferInfo2-size-00115)VUID-VkCopyBufferInfo2-size-00115  
  The `size` member of each element of `pRegions` **must** be less than or equal to the size of `srcBuffer` minus `srcOffset`
- [](#VUID-VkCopyBufferInfo2-size-00116)VUID-VkCopyBufferInfo2-size-00116  
  The `size` member of each element of `pRegions` **must** be less than or equal to the size of `dstBuffer` minus `dstOffset`
- [](#VUID-VkCopyBufferInfo2-pRegions-00117)VUID-VkCopyBufferInfo2-pRegions-00117  
  The union of the source regions, and the union of the destination regions, specified by the elements of `pRegions`, **must** not overlap in memory
- [](#VUID-VkCopyBufferInfo2-srcBuffer-00118)VUID-VkCopyBufferInfo2-srcBuffer-00118  
  `srcBuffer` **must** have been created with `VK_BUFFER_USAGE_TRANSFER_SRC_BIT` usage flag
- [](#VUID-VkCopyBufferInfo2-srcBuffer-00119)VUID-VkCopyBufferInfo2-srcBuffer-00119  
  If `srcBuffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkCopyBufferInfo2-dstBuffer-00120)VUID-VkCopyBufferInfo2-dstBuffer-00120  
  `dstBuffer` **must** have been created with `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag
- [](#VUID-VkCopyBufferInfo2-dstBuffer-00121)VUID-VkCopyBufferInfo2-dstBuffer-00121  
  If `dstBuffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object

Valid Usage (Implicit)

- [](#VUID-VkCopyBufferInfo2-sType-sType)VUID-VkCopyBufferInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_BUFFER_INFO_2`
- [](#VUID-VkCopyBufferInfo2-pNext-pNext)VUID-VkCopyBufferInfo2-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyBufferInfo2-srcBuffer-parameter)VUID-VkCopyBufferInfo2-srcBuffer-parameter  
  `srcBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkCopyBufferInfo2-dstBuffer-parameter)VUID-VkCopyBufferInfo2-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkCopyBufferInfo2-pRegions-parameter)VUID-VkCopyBufferInfo2-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount` valid [VkBufferCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCopy2.html) structures
- [](#VUID-VkCopyBufferInfo2-regionCount-arraylength)VUID-VkCopyBufferInfo2-regionCount-arraylength  
  `regionCount` **must** be greater than `0`
- [](#VUID-VkCopyBufferInfo2-commonparent)VUID-VkCopyBufferInfo2-commonparent  
  Both of `dstBuffer`, and `srcBuffer` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkBufferCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCopy2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBuffer2.html), [vkCmdCopyBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBuffer2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyBufferInfo2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0