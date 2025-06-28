# VkImagePlaneMemoryRequirementsInfo(3) Manual Page

## Name

VkImagePlaneMemoryRequirementsInfo - Structure specifying image plane for memory requirements



## [](#_c_specification)C Specification

To determine the memory requirements for a plane of a disjoint image, add a `VkImagePlaneMemoryRequirementsInfo` structure to the `pNext` chain of the `VkImageMemoryRequirementsInfo2` structure.

The `VkImagePlaneMemoryRequirementsInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkImagePlaneMemoryRequirementsInfo {
    VkStructureType          sType;
    const void*              pNext;
    VkImageAspectFlagBits    planeAspect;
} VkImagePlaneMemoryRequirementsInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkImagePlaneMemoryRequirementsInfo VkImagePlaneMemoryRequirementsInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `planeAspect` is a [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) value specifying the aspect corresponding to the image plane to query.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-02281)VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-02281  
  If the image’s `tiling` is `VK_IMAGE_TILING_LINEAR` or `VK_IMAGE_TILING_OPTIMAL`, then `planeAspect` **must** be a single valid [multi-planar aspect mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar-image-aspect) bit
- [](#VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-02282)VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-02282  
  If the image’s `tiling` is `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`, then `planeAspect` **must** be a single valid *memory plane* for the image (that is, `aspectMask` **must** specify a plane index that is less than the [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierPlaneCount` associated with the image’s `format` and [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierPropertiesEXT.html)::`drmFormatModifier`)

Valid Usage (Implicit)

- [](#VUID-VkImagePlaneMemoryRequirementsInfo-sType-sType)VUID-VkImagePlaneMemoryRequirementsInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO`
- [](#VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-parameter)VUID-VkImagePlaneMemoryRequirementsInfo-planeAspect-parameter  
  `planeAspect` **must** be a valid [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImagePlaneMemoryRequirementsInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0