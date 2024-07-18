# VkDrmFormatModifierPropertiesList2EXT(3) Manual Page

## Name

VkDrmFormatModifierPropertiesList2EXT - Structure specifying the list of
DRM format modifiers supported for a format



## <a href="#_c_specification" class="anchor"></a>C Specification

The list of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-drm-format-modifier"
target="_blank" rel="noopener">Linux DRM format modifiers</a> compatible
with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) **can** be obtained by adding a
[VkDrmFormatModifierPropertiesList2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesList2EXT.html)
structure to the `pNext` chain of
[VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html).

The
[VkDrmFormatModifierPropertiesList2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesList2EXT.html)
structure is defined as:

``` c
// Provided by VK_EXT_image_drm_format_modifier with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
typedef struct VkDrmFormatModifierPropertiesList2EXT {
    VkStructureType                       sType;
    void*                                 pNext;
    uint32_t                              drmFormatModifierCount;
    VkDrmFormatModifierProperties2EXT*    pDrmFormatModifierProperties;
} VkDrmFormatModifierPropertiesList2EXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `drmFormatModifierCount` is an inout parameter related to the number
  of modifiers compatible with the `format`, as described below.

- `pDrmFormatModifierProperties` is either `NULL` or a pointer to an
  array of
  [VkDrmFormatModifierProperties2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierProperties2EXT.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

If `pDrmFormatModifierProperties` is `NULL`, the number of modifiers
compatible with the queried `format` is returned in
`drmFormatModifierCount`. Otherwise, the application **must** set
`drmFormatModifierCount` to the length of the array
`pDrmFormatModifierProperties`; the function will write at most
`drmFormatModifierCount` elements to the array, and will return in
`drmFormatModifierCount` the number of elements written.

Among the elements in array `pDrmFormatModifierProperties`, each
returned `drmFormatModifier` **must** be unique.

Among the elements in array `pDrmFormatModifierProperties`, the bits
reported in `drmFormatModifierTilingFeatures` **must** include the bits
reported in the corresponding element of
`VkDrmFormatModifierPropertiesListEXT`::`pDrmFormatModifierProperties`.

Valid Usage (Implicit)

- <a href="#VUID-VkDrmFormatModifierPropertiesList2EXT-sType-sType"
  id="VUID-VkDrmFormatModifierPropertiesList2EXT-sType-sType"></a>
  VUID-VkDrmFormatModifierPropertiesList2EXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_LIST_2_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_drm_format_modifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_drm_format_modifier.html),
[VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkDrmFormatModifierProperties2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierProperties2EXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDrmFormatModifierPropertiesList2EXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
