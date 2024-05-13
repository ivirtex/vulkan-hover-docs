# VkBindImagePlaneMemoryInfo(3) Manual Page

## Name

VkBindImagePlaneMemoryInfo - Structure specifying how to bind an image
plane to memory



## <a href="#_c_specification" class="anchor"></a>C Specification

In order to bind *planes* of a *disjoint image*, add a
`VkBindImagePlaneMemoryInfo` structure to the `pNext` chain of
[VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html).

The `VkBindImagePlaneMemoryInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkBindImagePlaneMemoryInfo {
    VkStructureType          sType;
    const void*              pNext;
    VkImageAspectFlagBits    planeAspect;
} VkBindImagePlaneMemoryInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkBindImagePlaneMemoryInfo VkBindImagePlaneMemoryInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `planeAspect` is a `VkImageAspectFlagBits` value specifying the aspect
  of the disjoint image plane to bind.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBindImagePlaneMemoryInfo-planeAspect-02283"
  id="VUID-VkBindImagePlaneMemoryInfo-planeAspect-02283"></a>
  VUID-VkBindImagePlaneMemoryInfo-planeAspect-02283  
  If the image’s `tiling` is `VK_IMAGE_TILING_LINEAR` or
  `VK_IMAGE_TILING_OPTIMAL`, then `planeAspect` **must** be a single
  valid <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-planes-image-aspect"
  target="_blank" rel="noopener">multi-planar aspect mask</a> bit

- <a href="#VUID-VkBindImagePlaneMemoryInfo-planeAspect-02284"
  id="VUID-VkBindImagePlaneMemoryInfo-planeAspect-02284"></a>
  VUID-VkBindImagePlaneMemoryInfo-planeAspect-02284  
  If the image’s `tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`,
  then `planeAspect` **must** be a single valid *memory plane* for the
  image (that is, `aspectMask` **must** specify a plane index that is
  less than the
  [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierPlaneCount`
  associated with the image’s `format` and
  [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierPropertiesEXT.html)::`drmFormatModifier`)

Valid Usage (Implicit)

- <a href="#VUID-VkBindImagePlaneMemoryInfo-sType-sType"
  id="VUID-VkBindImagePlaneMemoryInfo-sType-sType"></a>
  VUID-VkBindImagePlaneMemoryInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO`

- <a href="#VUID-VkBindImagePlaneMemoryInfo-planeAspect-parameter"
  id="VUID-VkBindImagePlaneMemoryInfo-planeAspect-parameter"></a>
  VUID-VkBindImagePlaneMemoryInfo-planeAspect-parameter  
  `planeAspect` **must** be a valid
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindImagePlaneMemoryInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
