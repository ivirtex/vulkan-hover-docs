# VkSparseImageFormatFlagBits(3) Manual Page

## Name

VkSparseImageFormatFlagBits - Bitmask specifying additional information
about a sparse image resource



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties.html)::`flags`,
specifying additional information about the sparse resource, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkSparseImageFormatFlagBits {
    VK_SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT = 0x00000001,
    VK_SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT = 0x00000002,
    VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT = 0x00000004,
} VkSparseImageFormatFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SPARSE_IMAGE_FORMAT_SINGLE_MIPTAIL_BIT` specifies that the image
  uses a single mip tail region for all array layers.

- `VK_SPARSE_IMAGE_FORMAT_ALIGNED_MIP_SIZE_BIT` specifies that the first
  mip level whose dimensions are not integer multiples of the
  corresponding dimensions of the sparse image block begins the mip tail
  region.

- `VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT` specifies that the
  image uses non-standard sparse image block dimensions, and the
  `imageGranularity` values do not match the standard sparse image block
  dimensions for the given format.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkSparseImageFormatFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSparseImageFormatFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
