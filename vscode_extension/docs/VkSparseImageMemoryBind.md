# VkSparseImageMemoryBind(3) Manual Page

## Name

VkSparseImageMemoryBind - Structure specifying sparse image memory bind



## [](#_c_specification)C Specification

The `VkSparseImageMemoryBind` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkSparseImageMemoryBind {
    VkImageSubresource         subresource;
    VkOffset3D                 offset;
    VkExtent3D                 extent;
    VkDeviceMemory             memory;
    VkDeviceSize               memoryOffset;
    VkSparseMemoryBindFlags    flags;
} VkSparseImageMemoryBind;
```

## [](#_members)Members

- `subresource` is the image *aspect* and region of interest in the image.
- `offset` are the coordinates of the first texel within the image subresource to bind.
- `extent` is the size in texels of the region within the image subresource to bind. The extent **must** be a multiple of the sparse image block dimensions, except when binding sparse image blocks along the edge of an image subresource it **can** instead be such that any coordinate of `offset` + `extent` equals the corresponding dimensions of the image subresource.
- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object that the sparse image blocks of the image are bound to. If `memory` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the sparse image blocks are unbound.
- `memoryOffset` is an offset into [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object. If `memory` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), this value is ignored.
- `flags` are sparse memory binding flags.

## [](#_description)Description

Valid Usage

- [](#VUID-VkSparseImageMemoryBind-memory-01104)VUID-VkSparseImageMemoryBind-memory-01104  
  If the [`sparseResidencyAliased`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-sparseResidencyAliased) feature is not enabled, and if any other resources are bound to ranges of `memory`, the range of `memory` being bound **must** not overlap with those bound ranges
- [](#VUID-VkSparseImageMemoryBind-memory-01105)VUID-VkSparseImageMemoryBind-memory-01105  
  `memory` and `memoryOffset` **must** match the memory requirements of the calling command’s `image`, as described in section [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-association](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-association)
- [](#VUID-VkSparseImageMemoryBind-offset-01107)VUID-VkSparseImageMemoryBind-offset-01107  
  `offset.x` **must** be a multiple of the sparse image block width (`VkSparseImageFormatProperties`::`imageGranularity.width`) of the image
- [](#VUID-VkSparseImageMemoryBind-extent-09388)VUID-VkSparseImageMemoryBind-extent-09388  
  `extent.width` **must** be greater than `0`
- [](#VUID-VkSparseImageMemoryBind-extent-01108)VUID-VkSparseImageMemoryBind-extent-01108  
  `extent.width` **must** either be a multiple of the sparse image block width of the image, or else (`extent.width` + `offset.x`) **must** equal the width of the image subresource
- [](#VUID-VkSparseImageMemoryBind-offset-01109)VUID-VkSparseImageMemoryBind-offset-01109  
  `offset.y` **must** be a multiple of the sparse image block height (`VkSparseImageFormatProperties`::`imageGranularity.height`) of the image
- [](#VUID-VkSparseImageMemoryBind-extent-09389)VUID-VkSparseImageMemoryBind-extent-09389  
  `extent.height` **must** be greater than `0`
- [](#VUID-VkSparseImageMemoryBind-extent-01110)VUID-VkSparseImageMemoryBind-extent-01110  
  `extent.height` **must** either be a multiple of the sparse image block height of the image, or else (`extent.height` + `offset.y`) **must** equal the height of the image subresource
- [](#VUID-VkSparseImageMemoryBind-offset-01111)VUID-VkSparseImageMemoryBind-offset-01111  
  `offset.z` **must** be a multiple of the sparse image block depth (`VkSparseImageFormatProperties`::`imageGranularity.depth`) of the image
- [](#VUID-VkSparseImageMemoryBind-extent-09390)VUID-VkSparseImageMemoryBind-extent-09390  
  `extent.depth` **must** be greater than `0`
- [](#VUID-VkSparseImageMemoryBind-extent-01112)VUID-VkSparseImageMemoryBind-extent-01112  
  `extent.depth` **must** either be a multiple of the sparse image block depth of the image, or else (`extent.depth` + `offset.z`) **must** equal the depth of the image subresource
- [](#VUID-VkSparseImageMemoryBind-memory-02732)VUID-VkSparseImageMemoryBind-memory-02732  
  If `memory` was created with [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes` not equal to `0`, at least one handle type it contained **must** also have been set in [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes` when the image was created
- [](#VUID-VkSparseImageMemoryBind-memory-02733)VUID-VkSparseImageMemoryBind-memory-02733  
  If `memory` was created by a memory import operation, the external handle type of the imported memory **must** also have been set in [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes` when `image` was created

Valid Usage (Implicit)

- [](#VUID-VkSparseImageMemoryBind-subresource-parameter)VUID-VkSparseImageMemoryBind-subresource-parameter  
  `subresource` **must** be a valid [VkImageSubresource](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource.html) structure
- [](#VUID-VkSparseImageMemoryBind-memory-parameter)VUID-VkSparseImageMemoryBind-memory-parameter  
  If `memory` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-VkSparseImageMemoryBind-flags-parameter)VUID-VkSparseImageMemoryBind-flags-parameter  
  `flags` **must** be a valid combination of [VkSparseMemoryBindFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBindFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkImageSubresource](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource.html), [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html), [VkSparseImageMemoryBindInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryBindInfo.html), [VkSparseMemoryBindFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBindFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSparseImageMemoryBind).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0