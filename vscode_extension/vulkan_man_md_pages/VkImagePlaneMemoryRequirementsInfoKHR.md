# VkImagePlaneMemoryRequirementsInfo(3) Manual Page

## Name

VkImagePlaneMemoryRequirementsInfo - Structure specifying image plane
for memory requirements



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the memory requirements for a plane of a disjoint image,
add a `VkImagePlaneMemoryRequirementsInfo` structure to the `pNext`
chain of the `VkImageMemoryRequirementsInfo2` structure.

The `VkImagePlaneMemoryRequirementsInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkImagePlaneMemoryRequirementsInfo {
    VkStructureType          sType;
    const void*              pNext;
    VkImageAspectFlagBits    planeAspect;
} VkImagePlaneMemoryRequirementsInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkImagePlaneMemoryRequirementsInfo VkImagePlaneMemoryRequirementsInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `planeAspect` is a [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html)
  value specifying the aspect corresponding to the image plane to query.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-02281"
  id="VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-02281"></a>
  VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-02281  
  If the image’s `tiling` is `VK_IMAGE_TILING_LINEAR` or
  `VK_IMAGE_TILING_OPTIMAL`, then `planeAspect` **must** be a single
  valid <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-planes-image-aspect"
  target="_blank" rel="noopener">multi-planar aspect mask</a> bit

- <a href="#VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-02282"
  id="VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-02282"></a>
  VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-02282  
  If the image’s `tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`,
  then `planeAspect` **must** be a single valid *memory plane* for the
  image (that is, `aspectMask` **must** specify a plane index that is
  less than the
  [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierPlaneCount`
  associated with the image’s `format` and
  [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierPropertiesEXT.html)::`drmFormatModifier`)

Valid Usage (Implicit)

- <a href="#VUID-VkImagePlaneMemoryRequirementsInfo-sType-sType"
  id="VUID-VkImagePlaneMemoryRequirementsInfo-sType-sType"></a>
  VUID-VkImagePlaneMemoryRequirementsInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO`

- <a href="#VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-parameter"
  id="VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-parameter"></a>
  VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-parameter  
  `planeAspect` **must** be a valid
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImagePlaneMemoryRequirementsInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
