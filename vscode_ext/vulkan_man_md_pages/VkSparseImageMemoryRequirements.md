# VkSparseImageMemoryRequirements(3) Manual Page

## Name

VkSparseImageMemoryRequirements - Structure specifying sparse image
memory requirements



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSparseImageMemoryRequirements` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkSparseImageMemoryRequirements {
    VkSparseImageFormatProperties    formatProperties;
    uint32_t                         imageMipTailFirstLod;
    VkDeviceSize                     imageMipTailSize;
    VkDeviceSize                     imageMipTailOffset;
    VkDeviceSize                     imageMipTailStride;
} VkSparseImageMemoryRequirements;
```

## <a href="#_members" class="anchor"></a>Members

- `formatProperties` is a
  [VkSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties.html)
  structure specifying properties of the image format.

- `imageMipTailFirstLod` is the first mip level at which image
  subresources are included in the mip tail region.

- `imageMipTailSize` is the memory size (in bytes) of the mip tail
  region. If `formatProperties.flags` contains
  `VK_SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT`, this is the size of the
  whole mip tail, otherwise this is the size of the mip tail of a single
  array layer. This value is guaranteed to be a multiple of the sparse
  block size in bytes.

- `imageMipTailOffset` is the opaque memory offset used with
  [VkSparseImageOpaqueMemoryBindInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageOpaqueMemoryBindInfo.html)
  to bind the mip tail region(s).

- `imageMipTailStride` is the offset stride between each array-layer’s
  mip tail, if `formatProperties.flags` does not contain
  `VK_SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT` (otherwise the value is
  undefined).

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties.html),
[VkSparseImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageMemoryRequirements2.html),
[vkGetImageSparseMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSparseMemoryRequirements.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSparseImageMemoryRequirements"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
