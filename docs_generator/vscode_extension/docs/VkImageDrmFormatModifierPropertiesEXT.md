# VkImageDrmFormatModifierPropertiesEXT(3) Manual Page

## Name

VkImageDrmFormatModifierPropertiesEXT - Properties of an image's Linux DRM format modifier



## [](#_c_specification)C Specification

The [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierPropertiesEXT.html) structure is defined as:

```c++
// Provided by VK_EXT_image_drm_format_modifier
typedef struct VkImageDrmFormatModifierPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           drmFormatModifier;
} VkImageDrmFormatModifierPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `drmFormatModifier` returns the image’s [Linux DRM format modifier](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-drm-format-modifier).

## [](#_description)Description

If the `image` was created with [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html), then the returned `drmFormatModifier` **must** belong to the list of modifiers provided at time of image creation in [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)::`pDrmFormatModifiers`. If the `image` was created with [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html), then the returned `drmFormatModifier` **must** be the modifier provided at time of image creation in [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)::`drmFormatModifier`.

Valid Usage (Implicit)

- [](#VUID-VkImageDrmFormatModifierPropertiesEXT-sType-sType)VUID-VkImageDrmFormatModifierPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_PROPERTIES_EXT`
- [](#VUID-VkImageDrmFormatModifierPropertiesEXT-pNext-pNext)VUID-VkImageDrmFormatModifierPropertiesEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_EXT\_image\_drm\_format\_modifier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_drm_format_modifier.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageDrmFormatModifierPropertiesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageDrmFormatModifierPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0