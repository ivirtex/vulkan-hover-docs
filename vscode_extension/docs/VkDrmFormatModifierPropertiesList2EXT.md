# VkDrmFormatModifierPropertiesList2EXT(3) Manual Page

## Name

VkDrmFormatModifierPropertiesList2EXT - Structure specifying the list of DRM format modifiers supported for a format



## [](#_c_specification)C Specification

The list of [Linux DRM format modifiers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-drm-format-modifier) compatible with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) **can** be obtained by adding a [VkDrmFormatModifierPropertiesList2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierPropertiesList2EXT.html) structure to the `pNext` chain of [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties2.html).

The [VkDrmFormatModifierPropertiesList2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierPropertiesList2EXT.html) structure is defined as:

```c++
// Provided by VK_EXT_image_drm_format_modifier with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
typedef struct VkDrmFormatModifierPropertiesList2EXT {
    VkStructureType                       sType;
    void*                                 pNext;
    uint32_t                              drmFormatModifierCount;
    VkDrmFormatModifierProperties2EXT*    pDrmFormatModifierProperties;
} VkDrmFormatModifierPropertiesList2EXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `drmFormatModifierCount` is an inout parameter related to the number of modifiers compatible with the `format`, as described below.
- `pDrmFormatModifierProperties` is either `NULL` or a pointer to an array of [VkDrmFormatModifierProperties2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierProperties2EXT.html) structures.

## [](#_description)Description

If `pDrmFormatModifierProperties` is `NULL`, the number of modifiers compatible with the queried `format` is returned in `drmFormatModifierCount`. Otherwise, the application **must** set `drmFormatModifierCount` to the length of the array `pDrmFormatModifierProperties`; the function will write at most `drmFormatModifierCount` elements to the array, and will return in `drmFormatModifierCount` the number of elements written.

Among the elements in array `pDrmFormatModifierProperties`, each returned `drmFormatModifier` **must** be unique.

Among the elements in array `pDrmFormatModifierProperties`, the bits reported in `drmFormatModifierTilingFeatures` **must** include the bits reported in the corresponding element of `VkDrmFormatModifierPropertiesListEXT`::`pDrmFormatModifierProperties`.

Valid Usage (Implicit)

- [](#VUID-VkDrmFormatModifierPropertiesList2EXT-sType-sType)VUID-VkDrmFormatModifierPropertiesList2EXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_LIST_2_EXT`

## [](#_see_also)See Also

[VK\_EXT\_image\_drm\_format\_modifier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_drm_format_modifier.html), [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkDrmFormatModifierProperties2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDrmFormatModifierProperties2EXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDrmFormatModifierPropertiesList2EXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0