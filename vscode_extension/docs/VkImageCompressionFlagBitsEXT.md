# VkImageCompressionFlagBitsEXT(3) Manual Page

## Name

VkImageCompressionFlagBitsEXT - Bitmask specifying image compression controls



## [](#_c_specification)C Specification

Possible values of [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionControlEXT.html)::`flags`, specifying compression controls for an image, are:

```c++
// Provided by VK_EXT_image_compression_control
typedef enum VkImageCompressionFlagBitsEXT {
    VK_IMAGE_COMPRESSION_DEFAULT_EXT = 0,
    VK_IMAGE_COMPRESSION_FIXED_RATE_DEFAULT_EXT = 0x00000001,
    VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT = 0x00000002,
    VK_IMAGE_COMPRESSION_DISABLED_EXT = 0x00000004,
} VkImageCompressionFlagBitsEXT;
```

## [](#_description)Description

- `VK_IMAGE_COMPRESSION_DEFAULT_EXT` specifies that the default image compression setting is used. Implementations **must** not apply fixed-rate compression.
- `VK_IMAGE_COMPRESSION_FIXED_RATE_DEFAULT_EXT` specifies that the implementation **may** choose any supported fixed-rate compression setting in an implementation-defined manner based on the properties of the image.
- `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT` specifies that fixed-rate compression **may** be used and that the allowed compression rates are specified by [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionControlEXT.html)::`pFixedRateFlags`.
- `VK_IMAGE_COMPRESSION_DISABLED_EXT` specifies that all lossless and fixed-rate compression **should** be disabled.

If [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionControlEXT.html)::`flags` is `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT`, then the `i`th member of the `pFixedRateFlags` array specifies the allowed compression rates for the image’s `i`th plane.

Note

If `VK_IMAGE_COMPRESSION_DISABLED_EXT` is included in [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionControlEXT.html)::`flags`, both lossless and fixed-rate compression will be disabled. This is likely to have a negative impact on performance and is only intended to be used for debugging purposes.

## [](#_see_also)See Also

[VK\_EXT\_image\_compression\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_compression_control.html), [VkImageCompressionFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageCompressionFlagBitsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0