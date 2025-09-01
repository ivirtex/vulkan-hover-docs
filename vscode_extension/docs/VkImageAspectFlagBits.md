# VkImageAspectFlagBits(3) Manual Page

## Name

VkImageAspectFlagBits - Bitmask specifying which aspects of an image are included in a view



## [](#_c_specification)C Specification

Bits which **can** be set in an aspect mask to specify aspects of an image for purposes such as identifying a subresource, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkImageAspectFlagBits {
    VK_IMAGE_ASPECT_COLOR_BIT = 0x00000001,
    VK_IMAGE_ASPECT_DEPTH_BIT = 0x00000002,
    VK_IMAGE_ASPECT_STENCIL_BIT = 0x00000004,
    VK_IMAGE_ASPECT_METADATA_BIT = 0x00000008,
  // Provided by VK_VERSION_1_1
    VK_IMAGE_ASPECT_PLANE_0_BIT = 0x00000010,
  // Provided by VK_VERSION_1_1
    VK_IMAGE_ASPECT_PLANE_1_BIT = 0x00000020,
  // Provided by VK_VERSION_1_1
    VK_IMAGE_ASPECT_PLANE_2_BIT = 0x00000040,
  // Provided by VK_VERSION_1_3
    VK_IMAGE_ASPECT_NONE = 0,
  // Provided by VK_EXT_image_drm_format_modifier
    VK_IMAGE_ASPECT_MEMORY_PLANE_0_BIT_EXT = 0x00000080,
  // Provided by VK_EXT_image_drm_format_modifier
    VK_IMAGE_ASPECT_MEMORY_PLANE_1_BIT_EXT = 0x00000100,
  // Provided by VK_EXT_image_drm_format_modifier
    VK_IMAGE_ASPECT_MEMORY_PLANE_2_BIT_EXT = 0x00000200,
  // Provided by VK_EXT_image_drm_format_modifier
    VK_IMAGE_ASPECT_MEMORY_PLANE_3_BIT_EXT = 0x00000400,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_IMAGE_ASPECT_PLANE_0_BIT_KHR = VK_IMAGE_ASPECT_PLANE_0_BIT,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_IMAGE_ASPECT_PLANE_1_BIT_KHR = VK_IMAGE_ASPECT_PLANE_1_BIT,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_IMAGE_ASPECT_PLANE_2_BIT_KHR = VK_IMAGE_ASPECT_PLANE_2_BIT,
  // Provided by VK_KHR_maintenance4
    VK_IMAGE_ASPECT_NONE_KHR = VK_IMAGE_ASPECT_NONE,
} VkImageAspectFlagBits;
```

## [](#_description)Description

- `VK_IMAGE_ASPECT_NONE` specifies no image aspect, or the image aspect is not applicable.
- `VK_IMAGE_ASPECT_COLOR_BIT` specifies the color aspect.
- `VK_IMAGE_ASPECT_DEPTH_BIT` specifies the depth aspect.
- `VK_IMAGE_ASPECT_STENCIL_BIT` specifies the stencil aspect.
- `VK_IMAGE_ASPECT_METADATA_BIT` specifies the metadata aspect used for [sparse resource](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#sparsememory) operations.
- `VK_IMAGE_ASPECT_PLANE_0_BIT` specifies plane 0 of a *multi-planar* image format.
- `VK_IMAGE_ASPECT_PLANE_1_BIT` specifies plane 1 of a *multi-planar* image format.
- `VK_IMAGE_ASPECT_PLANE_2_BIT` specifies plane 2 of a *multi-planar* image format.
- `VK_IMAGE_ASPECT_MEMORY_PLANE_0_BIT_EXT` specifies *memory plane* 0.
- `VK_IMAGE_ASPECT_MEMORY_PLANE_1_BIT_EXT` specifies *memory plane* 1.
- `VK_IMAGE_ASPECT_MEMORY_PLANE_2_BIT_EXT` specifies *memory plane* 2.
- `VK_IMAGE_ASPECT_MEMORY_PLANE_3_BIT_EXT` specifies *memory plane* 3.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImagePlaneMemoryInfo.html), [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageMemoryRequirements.html), [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html), [VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlags.html), [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePlaneMemoryRequirementsInfo.html), [VkImportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMetalTextureInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageAspectFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0