# VkCopyBufferInfo2(3) Manual Page

## Name

VkCopyBufferInfo2 - Structure specifying parameters of a buffer copy
command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyBufferInfo2` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_copy_commands2
typedef VkCopyBufferInfo2 VkCopyBufferInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcBuffer` is the source buffer.

- `dstBuffer` is the destination buffer.

- `regionCount` is the number of regions to copy.

- `pRegions` is a pointer to an array of
  [VkBufferCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCopy2.html) structures specifying the regions
  to copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCopyBufferInfo2-srcOffset-00113"
  id="VUID-VkCopyBufferInfo2-srcOffset-00113"></a>
  VUID-VkCopyBufferInfo2-srcOffset-00113  
  The `srcOffset` member of each element of `pRegions` **must** be less
  than the size of `srcBuffer`

- <a href="#VUID-VkCopyBufferInfo2-dstOffset-00114"
  id="VUID-VkCopyBufferInfo2-dstOffset-00114"></a>
  VUID-VkCopyBufferInfo2-dstOffset-00114  
  The `dstOffset` member of each element of `pRegions` **must** be less
  than the size of `dstBuffer`

- <a href="#VUID-VkCopyBufferInfo2-size-00115"
  id="VUID-VkCopyBufferInfo2-size-00115"></a>
  VUID-VkCopyBufferInfo2-size-00115  
  The `size` member of each element of `pRegions` **must** be less than
  or equal to the size of `srcBuffer` minus `srcOffset`

- <a href="#VUID-VkCopyBufferInfo2-size-00116"
  id="VUID-VkCopyBufferInfo2-size-00116"></a>
  VUID-VkCopyBufferInfo2-size-00116  
  The `size` member of each element of `pRegions` **must** be less than
  or equal to the size of `dstBuffer` minus `dstOffset`

- <a href="#VUID-VkCopyBufferInfo2-pRegions-00117"
  id="VUID-VkCopyBufferInfo2-pRegions-00117"></a>
  VUID-VkCopyBufferInfo2-pRegions-00117  
  The union of the source regions, and the union of the destination
  regions, specified by the elements of `pRegions`, **must** not overlap
  in memory

- <a href="#VUID-VkCopyBufferInfo2-srcBuffer-00118"
  id="VUID-VkCopyBufferInfo2-srcBuffer-00118"></a>
  VUID-VkCopyBufferInfo2-srcBuffer-00118  
  `srcBuffer` **must** have been created with
  `VK_BUFFER_USAGE_TRANSFER_SRC_BIT` usage flag

- <a href="#VUID-VkCopyBufferInfo2-srcBuffer-00119"
  id="VUID-VkCopyBufferInfo2-srcBuffer-00119"></a>
  VUID-VkCopyBufferInfo2-srcBuffer-00119  
  If `srcBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkCopyBufferInfo2-dstBuffer-00120"
  id="VUID-VkCopyBufferInfo2-dstBuffer-00120"></a>
  VUID-VkCopyBufferInfo2-dstBuffer-00120  
  `dstBuffer` **must** have been created with
  `VK_BUFFER_USAGE_TRANSFER_DST_BIT` usage flag

- <a href="#VUID-VkCopyBufferInfo2-dstBuffer-00121"
  id="VUID-VkCopyBufferInfo2-dstBuffer-00121"></a>
  VUID-VkCopyBufferInfo2-dstBuffer-00121  
  If `dstBuffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

Valid Usage (Implicit)

- <a href="#VUID-VkCopyBufferInfo2-sType-sType"
  id="VUID-VkCopyBufferInfo2-sType-sType"></a>
  VUID-VkCopyBufferInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_BUFFER_INFO_2`

- <a href="#VUID-VkCopyBufferInfo2-pNext-pNext"
  id="VUID-VkCopyBufferInfo2-pNext-pNext"></a>
  VUID-VkCopyBufferInfo2-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyBufferInfo2-srcBuffer-parameter"
  id="VUID-VkCopyBufferInfo2-srcBuffer-parameter"></a>
  VUID-VkCopyBufferInfo2-srcBuffer-parameter  
  `srcBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkCopyBufferInfo2-dstBuffer-parameter"
  id="VUID-VkCopyBufferInfo2-dstBuffer-parameter"></a>
  VUID-VkCopyBufferInfo2-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkCopyBufferInfo2-pRegions-parameter"
  id="VUID-VkCopyBufferInfo2-pRegions-parameter"></a>
  VUID-VkCopyBufferInfo2-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkBufferCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCopy2.html) structures

- <a href="#VUID-VkCopyBufferInfo2-regionCount-arraylength"
  id="VUID-VkCopyBufferInfo2-regionCount-arraylength"></a>
  VUID-VkCopyBufferInfo2-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-VkCopyBufferInfo2-commonparent"
  id="VUID-VkCopyBufferInfo2-commonparent"></a>
  VUID-VkCopyBufferInfo2-commonparent  
  Both of `dstBuffer`, and `srcBuffer` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkBufferCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCopy2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdCopyBuffer2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBuffer2.html),
[vkCmdCopyBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBuffer2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyBufferInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
