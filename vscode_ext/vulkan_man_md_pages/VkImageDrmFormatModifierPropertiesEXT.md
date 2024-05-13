# VkImageDrmFormatModifierPropertiesEXT(3) Manual Page

## Name

VkImageDrmFormatModifierPropertiesEXT - Properties of an image's Linux
DRM format modifier



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierPropertiesEXT.html)
structure is defined as:

``` c
// Provided by VK_EXT_image_drm_format_modifier
typedef struct VkImageDrmFormatModifierPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           drmFormatModifier;
} VkImageDrmFormatModifierPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `drmFormatModifier` returns the image’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-drm-format-modifier"
  target="_blank" rel="noopener">Linux DRM format modifier</a>.

## <a href="#_description" class="anchor"></a>Description

If the `image` was created with
[VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html),
then the returned `drmFormatModifier` **must** belong to the list of
modifiers provided at time of image creation in
[VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)::`pDrmFormatModifiers`.
If the `image` was created with
[VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html),
then the returned `drmFormatModifier` **must** be the modifier provided
at time of image creation in
[VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)::`drmFormatModifier`.

Valid Usage (Implicit)

- <a href="#VUID-VkImageDrmFormatModifierPropertiesEXT-sType-sType"
  id="VUID-VkImageDrmFormatModifierPropertiesEXT-sType-sType"></a>
  VUID-VkImageDrmFormatModifierPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_PROPERTIES_EXT`

- <a href="#VUID-VkImageDrmFormatModifierPropertiesEXT-pNext-pNext"
  id="VUID-VkImageDrmFormatModifierPropertiesEXT-pNext-pNext"></a>
  VUID-VkImageDrmFormatModifierPropertiesEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_drm_format_modifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_drm_format_modifier.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageDrmFormatModifierPropertiesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageDrmFormatModifierPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
