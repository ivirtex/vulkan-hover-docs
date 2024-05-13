# VkSparseImageMemoryBind(3) Manual Page

## Name

VkSparseImageMemoryBind - Structure specifying sparse image memory bind



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSparseImageMemoryBind` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `subresource` is the image *aspect* and region of interest in the
  image.

- `offset` are the coordinates of the first texel within the image
  subresource to bind.

- `extent` is the size in texels of the region within the image
  subresource to bind. The extent **must** be a multiple of the sparse
  image block dimensions, except when binding sparse image blocks along
  the edge of an image subresource it **can** instead be such that any
  coordinate of `offset` + `extent` equals the corresponding dimensions
  of the image subresource.

- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object that the
  sparse image blocks of the image are bound to. If `memory` is
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the sparse image blocks are
  unbound.

- `memoryOffset` is an offset into [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  object. If `memory` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), this
  value is ignored.

- `flags` are sparse memory binding flags.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSparseImageMemoryBind-memory-01104"
  id="VUID-VkSparseImageMemoryBind-memory-01104"></a>
  VUID-VkSparseImageMemoryBind-memory-01104  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseResidencyAliased"
  target="_blank" rel="noopener"><code>sparseResidencyAliased</code></a>
  feature is not enabled, and if any other resources are bound to ranges
  of `memory`, the range of `memory` being bound **must** not overlap
  with those bound ranges

- <a href="#VUID-VkSparseImageMemoryBind-memory-01105"
  id="VUID-VkSparseImageMemoryBind-memory-01105"></a>
  VUID-VkSparseImageMemoryBind-memory-01105  
  `memory` and `memoryOffset` **must** match the memory requirements of
  the calling command’s `image`, as described in section <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-association"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-association</a>

- <a href="#VUID-VkSparseImageMemoryBind-offset-01107"
  id="VUID-VkSparseImageMemoryBind-offset-01107"></a>
  VUID-VkSparseImageMemoryBind-offset-01107  
  `offset.x` **must** be a multiple of the sparse image block width
  (`VkSparseImageFormatProperties`::`imageGranularity.width`) of the
  image

- <a href="#VUID-VkSparseImageMemoryBind-extent-09388"
  id="VUID-VkSparseImageMemoryBind-extent-09388"></a>
  VUID-VkSparseImageMemoryBind-extent-09388  
  `extent.width` **must** be greater than `0`

- <a href="#VUID-VkSparseImageMemoryBind-extent-01108"
  id="VUID-VkSparseImageMemoryBind-extent-01108"></a>
  VUID-VkSparseImageMemoryBind-extent-01108  
  `extent.width` **must** either be a multiple of the sparse image block
  width of the image, or else (`extent.width` + `offset.x`) **must**
  equal the width of the image subresource

- <a href="#VUID-VkSparseImageMemoryBind-offset-01109"
  id="VUID-VkSparseImageMemoryBind-offset-01109"></a>
  VUID-VkSparseImageMemoryBind-offset-01109  
  `offset.y` **must** be a multiple of the sparse image block height
  (`VkSparseImageFormatProperties`::`imageGranularity.height`) of the
  image

- <a href="#VUID-VkSparseImageMemoryBind-extent-09389"
  id="VUID-VkSparseImageMemoryBind-extent-09389"></a>
  VUID-VkSparseImageMemoryBind-extent-09389  
  `extent.height` **must** be greater than `0`

- <a href="#VUID-VkSparseImageMemoryBind-extent-01110"
  id="VUID-VkSparseImageMemoryBind-extent-01110"></a>
  VUID-VkSparseImageMemoryBind-extent-01110  
  `extent.height` **must** either be a multiple of the sparse image
  block height of the image, or else (`extent.height` + `offset.y`)
  **must** equal the height of the image subresource

- <a href="#VUID-VkSparseImageMemoryBind-offset-01111"
  id="VUID-VkSparseImageMemoryBind-offset-01111"></a>
  VUID-VkSparseImageMemoryBind-offset-01111  
  `offset.z` **must** be a multiple of the sparse image block depth
  (`VkSparseImageFormatProperties`::`imageGranularity.depth`) of the
  image

- <a href="#VUID-VkSparseImageMemoryBind-extent-09390"
  id="VUID-VkSparseImageMemoryBind-extent-09390"></a>
  VUID-VkSparseImageMemoryBind-extent-09390  
  `extent.depth` **must** be greater than `0`

- <a href="#VUID-VkSparseImageMemoryBind-extent-01112"
  id="VUID-VkSparseImageMemoryBind-extent-01112"></a>
  VUID-VkSparseImageMemoryBind-extent-01112  
  `extent.depth` **must** either be a multiple of the sparse image block
  depth of the image, or else (`extent.depth` + `offset.z`) **must**
  equal the depth of the image subresource

- <a href="#VUID-VkSparseImageMemoryBind-memory-02732"
  id="VUID-VkSparseImageMemoryBind-memory-02732"></a>
  VUID-VkSparseImageMemoryBind-memory-02732  
  If `memory` was created with
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes`
  not equal to `0`, at least one handle type it contained **must** also
  have been set in
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes`
  when the image was created

- <a href="#VUID-VkSparseImageMemoryBind-memory-02733"
  id="VUID-VkSparseImageMemoryBind-memory-02733"></a>
  VUID-VkSparseImageMemoryBind-memory-02733  
  If `memory` was created by a memory import operation, the external
  handle type of the imported memory **must** also have been set in
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes`
  when `image` was created

Valid Usage (Implicit)

- <a href="#VUID-VkSparseImageMemoryBind-subresource-parameter"
  id="VUID-VkSparseImageMemoryBind-subresource-parameter"></a>
  VUID-VkSparseImageMemoryBind-subresource-parameter  
  `subresource` **must** be a valid
  [VkImageSubresource](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource.html) structure

- <a href="#VUID-VkSparseImageMemoryBind-memory-parameter"
  id="VUID-VkSparseImageMemoryBind-memory-parameter"></a>
  VUID-VkSparseImageMemoryBind-memory-parameter  
  If `memory` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `memory`
  **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) handle

- <a href="#VUID-VkSparseImageMemoryBind-flags-parameter"
  id="VUID-VkSparseImageMemoryBind-flags-parameter"></a>
  VUID-VkSparseImageMemoryBind-flags-parameter  
  `flags` **must** be a valid combination of
  [VkSparseMemoryBindFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBindFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html),
[VkImageSubresource](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource.html),
[VkOffset3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset3D.html),
[VkSparseImageMemoryBindInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryBindInfo.html),
[VkSparseMemoryBindFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBindFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSparseImageMemoryBind"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
