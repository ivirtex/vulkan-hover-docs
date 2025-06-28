# VkImageDrmFormatModifierExplicitCreateInfoEXT(3) Manual Page

## Name

VkImageDrmFormatModifierExplicitCreateInfoEXT - Specify that an image be created with the provided DRM format modifier and explicit memory layout



## [](#_c_specification)C Specification

If the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) includes a [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html) structure, then the image will be created with the [Linux DRM format modifier](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-drm-format-modifier) and memory layout defined by the structure.

The [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html) structure is defined as:

```c++
// Provided by VK_EXT_image_drm_format_modifier
typedef struct VkImageDrmFormatModifierExplicitCreateInfoEXT {
    VkStructureType               sType;
    const void*                   pNext;
    uint64_t                      drmFormatModifier;
    uint32_t                      drmFormatModifierPlaneCount;
    const VkSubresourceLayout*    pPlaneLayouts;
} VkImageDrmFormatModifierExplicitCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `drmFormatModifier` is the *Linux DRM format modifier* with which the image will be created.
- `drmFormatModifierPlaneCount` is the number of *memory planes* in the image (as reported by [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierPropertiesEXT.html)) as well as the length of the `pPlaneLayouts` array.
- `pPlaneLayouts` is a pointer to an array of [VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout.html) structures describing the image’s *memory planes*.

## [](#_description)Description

The `i`th member of `pPlaneLayouts` describes the layout of the image’s `i`th *memory plane* (that is, `VK_IMAGE_ASPECT_MEMORY_PLANE_i_BIT_EXT`). In each element of `pPlaneLayouts`, the implementation **must** ignore `size`. The implementation calculates the size of each plane, which the application **can** query with [vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout.html).

When creating an image with [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html), it is the application’s responsibility to satisfy all valid usage requirements. However, the implementation **must** validate that the provided `pPlaneLayouts`, when combined with the provided `drmFormatModifier` and other creation parameters in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) and its `pNext` chain, produce a valid image. (This validation is necessarily implementation-dependent and outside the scope of Vulkan, and therefore not described by valid usage requirements). If this validation fails, then [vkCreateImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImage.html) returns `VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT`.

Valid Usage

- [](#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifier-02264)VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifier-02264  
  `drmFormatModifier` **must** be compatible with the parameters in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) and its `pNext` chain, as determined by querying [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html) extended with [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)
- [](#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifierPlaneCount-02265)VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifierPlaneCount-02265  
  `drmFormatModifierPlaneCount` **must** be equal to the [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierPlaneCount` associated with [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`format` and `drmFormatModifier`, as found by querying [VkDrmFormatModifierPropertiesListEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierPropertiesListEXT.html)
- [](#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-size-02267)VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-size-02267  
  For each element of `pPlaneLayouts`, `size` **must** be 0
- [](#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-arrayPitch-02268)VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-arrayPitch-02268  
  For each element of `pPlaneLayouts`, `arrayPitch` **must** be 0 if [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`arrayLayers` is 1
- [](#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-depthPitch-02269)VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-depthPitch-02269  
  For each element of `pPlaneLayouts`, `depthPitch` **must** be 0 if [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`extent.depth` is 1

Valid Usage (Implicit)

- [](#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-sType-sType)VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_EXPLICIT_CREATE_INFO_EXT`
- [](#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-pPlaneLayouts-parameter)VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-pPlaneLayouts-parameter  
  `pPlaneLayouts` **must** be a valid pointer to an array of `drmFormatModifierPlaneCount` [VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout.html) structures
- [](#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifierPlaneCount-arraylength)VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifierPlaneCount-arraylength  
  `drmFormatModifierPlaneCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_image\_drm\_format\_modifier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_drm_format_modifier.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageDrmFormatModifierExplicitCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0