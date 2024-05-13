# VkImageAspectFlagBits(3) Manual Page

## Name

VkImageAspectFlagBits - Bitmask specifying which aspects of an image are
included in a view



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in an aspect mask to specify aspects of an
image for purposes such as identifying a subresource, are:

``` c
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

## <a href="#_description" class="anchor"></a>Description

- `VK_IMAGE_ASPECT_NONE` specifies no image aspect, or the image aspect
  is not applicable.

- `VK_IMAGE_ASPECT_COLOR_BIT` specifies the color aspect.

- `VK_IMAGE_ASPECT_DEPTH_BIT` specifies the depth aspect.

- `VK_IMAGE_ASPECT_STENCIL_BIT` specifies the stencil aspect.

- `VK_IMAGE_ASPECT_METADATA_BIT` specifies the metadata aspect used for
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#sparsememory"
  target="_blank" rel="noopener">sparse resource</a> operations.

- `VK_IMAGE_ASPECT_PLANE_0_BIT` specifies plane 0 of a *multi-planar*
  image format.

- `VK_IMAGE_ASPECT_PLANE_1_BIT` specifies plane 1 of a *multi-planar*
  image format.

- `VK_IMAGE_ASPECT_PLANE_2_BIT` specifies plane 2 of a *multi-planar*
  image format.

- `VK_IMAGE_ASPECT_MEMORY_PLANE_0_BIT_EXT` specifies *memory plane* 0.

- `VK_IMAGE_ASPECT_MEMORY_PLANE_1_BIT_EXT` specifies *memory plane* 1.

- `VK_IMAGE_ASPECT_MEMORY_PLANE_2_BIT_EXT` specifies *memory plane* 2.

- `VK_IMAGE_ASPECT_MEMORY_PLANE_3_BIT_EXT` specifies *memory plane* 3.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html),
[VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageMemoryRequirements.html),
[VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalTextureInfoEXT.html),
[VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlags.html),
[VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html),
[VkImportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalTextureInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageAspectFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
