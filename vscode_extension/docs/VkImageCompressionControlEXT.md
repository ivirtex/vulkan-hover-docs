# VkImageCompressionControlEXT(3) Manual Page

## Name

VkImageCompressionControlEXT - Specify image compression properties



## [](#_c_specification)C Specification

If the `pNext` list of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) includes a `VkImageCompressionControlEXT` structure, then that structure describes compression controls for this image.

The `VkImageCompressionControlEXT` structure is defined as:

```c++
// Provided by VK_EXT_image_compression_control
typedef struct VkImageCompressionControlEXT {
    VkStructureType                         sType;
    const void*                             pNext;
    VkImageCompressionFlagsEXT              flags;
    uint32_t                                compressionControlPlaneCount;
    VkImageCompressionFixedRateFlagsEXT*    pFixedRateFlags;
} VkImageCompressionControlEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkImageCompressionFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFlagBitsEXT.html) describing compression controls for the image.
- `compressionControlPlaneCount` is the number of entries in the `pFixedRateFlags` array.
- `pFixedRateFlags` is `NULL` or a pointer to an array of [VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFixedRateFlagsEXT.html) bitfields describing allowed fixed-rate compression rates of each image plane. It is ignored if `flags` does not include `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT`.

## [](#_description)Description

If enabled, fixed-rate compression is done in an implementation-defined manner and **may** be applied at block granularity. In that case, a write to an individual texel **may** modify the value of other texels in the same block.

Valid Usage

- [](#VUID-VkImageCompressionControlEXT-flags-06747)VUID-VkImageCompressionControlEXT-flags-06747  
  `flags` **must** be one of `VK_IMAGE_COMPRESSION_DEFAULT_EXT`, `VK_IMAGE_COMPRESSION_FIXED_RATE_DEFAULT_EXT`, `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT`, or `VK_IMAGE_COMPRESSION_DISABLED_EXT`
- [](#VUID-VkImageCompressionControlEXT-flags-06748)VUID-VkImageCompressionControlEXT-flags-06748  
  If `flags` includes `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT`, `pFixedRateFlags` **must** not be `NULL`

Valid Usage (Implicit)

- [](#VUID-VkImageCompressionControlEXT-sType-sType)VUID-VkImageCompressionControlEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_COMPRESSION_CONTROL_EXT`

Note

Some combinations of compression properties may not be supported. For example, some implementations may not support different fixed-rate compression rates per plane of a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar) and will not be able to enable fixed-rate compression for any plane if the requested rates differ.

## [](#_see_also)See Also

[VK\_EXT\_image\_compression\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_compression_control.html), [VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFixedRateFlagsEXT.html), [VkImageCompressionFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageCompressionControlEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0