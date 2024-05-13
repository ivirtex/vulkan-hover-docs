# VkBindImageMemoryDeviceGroupInfo(3) Manual Page

## Name

VkBindImageMemoryDeviceGroupInfo - Structure specifying device within a
group to bind to



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindImageMemoryDeviceGroupInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkBindImageMemoryDeviceGroupInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           deviceIndexCount;
    const uint32_t*    pDeviceIndices;
    uint32_t           splitInstanceBindRegionCount;
    const VkRect2D*    pSplitInstanceBindRegions;
} VkBindImageMemoryDeviceGroupInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_bind_memory2 with VK_KHR_device_group
typedef VkBindImageMemoryDeviceGroupInfo VkBindImageMemoryDeviceGroupInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `deviceIndexCount` is the number of elements in `pDeviceIndices`.

- `pDeviceIndices` is a pointer to an array of device indices.

- `splitInstanceBindRegionCount` is the number of elements in
  `pSplitInstanceBindRegions`.

- `pSplitInstanceBindRegions` is a pointer to an array of
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures describing which regions of the
  image are attached to each instance of memory.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of
[VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html) includes a
`VkBindImageMemoryDeviceGroupInfo` structure, then that structure
determines how memory is bound to images across multiple devices in a
device group.

If `deviceIndexCount` is greater than zero, then on device index i
`image` is attached to the instance of the memory on the physical device
with device index `pDeviceIndices`\[i\].

Let N be the number of physical devices in the logical device. If
`splitInstanceBindRegionCount` is greater than zero, then
`pSplitInstanceBindRegions` is a pointer to an array of N<sup>2</sup>
rectangles, where the image region specified by the rectangle at element
i\*N+j in resource instance i is bound to the memory instance j. The
blocks of the memory that are bound to each sparse image block region
use an offset in memory, relative to `memoryOffset`, computed as if the
whole image was being bound to a contiguous range of memory. In other
words, horizontally adjacent image blocks use consecutive blocks of
memory, vertically adjacent image blocks are separated by the number of
bytes per block multiplied by the width in blocks of `image`, and the
block at (0,0) corresponds to memory starting at `memoryOffset`.

If `splitInstanceBindRegionCount` and `deviceIndexCount` are zero and
the memory comes from a memory heap with the
`VK_MEMORY_HEAP_MULTI_INSTANCE_BIT` bit set, then it is as if
`pDeviceIndices` contains consecutive indices from zero to the number of
physical devices in the logical device, minus one. In other words, by
default each physical device attaches to its own instance of the memory.

If `splitInstanceBindRegionCount` and `deviceIndexCount` are zero and
the memory comes from a memory heap without the
`VK_MEMORY_HEAP_MULTI_INSTANCE_BIT` bit set, then it is as if
`pDeviceIndices` contains an array of zeros. In other words, by default
each physical device attaches to instance zero.

Valid Usage

- <a href="#VUID-VkBindImageMemoryDeviceGroupInfo-deviceIndexCount-01633"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-deviceIndexCount-01633"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-deviceIndexCount-01633  
  At least one of `deviceIndexCount` and `splitInstanceBindRegionCount`
  **must** be zero

- <a href="#VUID-VkBindImageMemoryDeviceGroupInfo-deviceIndexCount-01634"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-deviceIndexCount-01634"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-deviceIndexCount-01634  
  `deviceIndexCount` **must** either be zero or equal to the number of
  physical devices in the logical device

- <a href="#VUID-VkBindImageMemoryDeviceGroupInfo-pDeviceIndices-01635"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-pDeviceIndices-01635"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-pDeviceIndices-01635  
  All elements of `pDeviceIndices` **must** be valid device indices

- <a
  href="#VUID-VkBindImageMemoryDeviceGroupInfo-splitInstanceBindRegionCount-01636"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-splitInstanceBindRegionCount-01636"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-splitInstanceBindRegionCount-01636  
  `splitInstanceBindRegionCount` **must** either be zero or equal to the
  number of physical devices in the logical device squared

- <a
  href="#VUID-VkBindImageMemoryDeviceGroupInfo-pSplitInstanceBindRegions-01637"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-pSplitInstanceBindRegions-01637"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-pSplitInstanceBindRegions-01637  
  Elements of `pSplitInstanceBindRegions` that correspond to the same
  instance of an image **must** not overlap

- <a href="#VUID-VkBindImageMemoryDeviceGroupInfo-offset-01638"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-offset-01638"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-offset-01638  
  The `offset.x` member of any element of `pSplitInstanceBindRegions`
  **must** be a multiple of the sparse image block width
  (`VkSparseImageFormatProperties`::`imageGranularity.width`) of all
  non-metadata aspects of the image

- <a href="#VUID-VkBindImageMemoryDeviceGroupInfo-offset-01639"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-offset-01639"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-offset-01639  
  The `offset.y` member of any element of `pSplitInstanceBindRegions`
  **must** be a multiple of the sparse image block height
  (`VkSparseImageFormatProperties`::`imageGranularity.height`) of all
  non-metadata aspects of the image

- <a href="#VUID-VkBindImageMemoryDeviceGroupInfo-extent-01640"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-extent-01640"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-extent-01640  
  The `extent.width` member of any element of
  `pSplitInstanceBindRegions` **must** either be a multiple of the
  sparse image block width of all non-metadata aspects of the image, or
  else `extent.width` + `offset.x` **must** equal the width of the image
  subresource

- <a href="#VUID-VkBindImageMemoryDeviceGroupInfo-extent-01641"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-extent-01641"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-extent-01641  
  The `extent.height` member of any element of
  `pSplitInstanceBindRegions` **must** either be a multiple of the
  sparse image block height of all non-metadata aspects of the image, or
  else `extent.height` + `offset.y` **must** equal the height of the
  image subresource

Valid Usage (Implicit)

- <a href="#VUID-VkBindImageMemoryDeviceGroupInfo-sType-sType"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-sType-sType"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO`

- <a
  href="#VUID-VkBindImageMemoryDeviceGroupInfo-pDeviceIndices-parameter"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-pDeviceIndices-parameter"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-pDeviceIndices-parameter  
  If `deviceIndexCount` is not `0`, `pDeviceIndices` **must** be a valid
  pointer to an array of `deviceIndexCount` `uint32_t` values

- <a
  href="#VUID-VkBindImageMemoryDeviceGroupInfo-pSplitInstanceBindRegions-parameter"
  id="VUID-VkBindImageMemoryDeviceGroupInfo-pSplitInstanceBindRegions-parameter"></a>
  VUID-VkBindImageMemoryDeviceGroupInfo-pSplitInstanceBindRegions-parameter  
  If `splitInstanceBindRegionCount` is not `0`,
  `pSplitInstanceBindRegions` **must** be a valid pointer to an array of
  `splitInstanceBindRegionCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindImageMemoryDeviceGroupInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
