# VkImageCompressionPropertiesEXT(3) Manual Page

## Name

VkImageCompressionPropertiesEXT - Compression properties of an image



## [](#_c_specification)C Specification

To query the compression properties of an image, add a [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionPropertiesEXT.html) structure to the `pNext` chain of the [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html) structure in a call to [vkGetImageSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2.html) or [vkGetImageSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2.html).

To determine the compression rates that are supported for a given image format, add a [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionPropertiesEXT.html) structure to the `pNext` chain of the [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) structure in a call to [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html).

Note

Since fixed-rate compression is disabled by default, the [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionPropertiesEXT.html) structure passed to [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html) will not indicate any fixed-rate compression support unless a [VkImageCompressionControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionControlEXT.html) structure is also included in the `pNext` chain of the [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html) structure passed to the same command.

The `VkImageCompressionPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_image_compression_control
typedef struct VkImageCompressionPropertiesEXT {
    VkStructureType                        sType;
    void*                                  pNext;
    VkImageCompressionFlagsEXT             imageCompressionFlags;
    VkImageCompressionFixedRateFlagsEXT    imageCompressionFixedRateFlags;
} VkImageCompressionPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `imageCompressionFlags` returns a value describing the compression controls that apply to the image. The value will be either `VK_IMAGE_COMPRESSION_DEFAULT_EXT` to indicate no fixed-rate compression, `VK_IMAGE_COMPRESSION_FIXED_RATE_EXPLICIT_EXT` to indicate fixed-rate compression, or `VK_IMAGE_COMPRESSION_DISABLED_EXT` to indicate no compression.
- `imageCompressionFixedRateFlags` returns a [VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFixedRateFlagsEXT.html) value describing the compression rates that apply to the specified aspect of the image.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkImageCompressionPropertiesEXT-sType-sType)VUID-VkImageCompressionPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_COMPRESSION_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_image\_compression\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_compression_control.html), [VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFixedRateFlagsEXT.html), [VkImageCompressionFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCompressionFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageCompressionPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0