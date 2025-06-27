# VkImageDrmFormatModifierListCreateInfoEXT(3) Manual Page

## Name

VkImageDrmFormatModifierListCreateInfoEXT - Specify that an image must
be created with a DRM format modifier from the provided list



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
includes a
[VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)
structure, then the image will be created with one of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-drm-format-modifier"
target="_blank" rel="noopener">Linux DRM format modifiers</a> listed in
the structure. The choice of modifier is implementation-dependent.

The
[VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)
structure is defined as:

``` c
// Provided by VK_EXT_image_drm_format_modifier
typedef struct VkImageDrmFormatModifierListCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           drmFormatModifierCount;
    const uint64_t*    pDrmFormatModifiers;
} VkImageDrmFormatModifierListCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `drmFormatModifierCount` is the length of the `pDrmFormatModifiers`
  array.

- `pDrmFormatModifiers` is a pointer to an array of *Linux DRM format
  modifiers*.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkImageDrmFormatModifierListCreateInfoEXT-pDrmFormatModifiers-02263"
  id="VUID-VkImageDrmFormatModifierListCreateInfoEXT-pDrmFormatModifiers-02263"></a>
  VUID-VkImageDrmFormatModifierListCreateInfoEXT-pDrmFormatModifiers-02263  
  Each *modifier* in `pDrmFormatModifiers` **must** be compatible with
  the parameters in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) and its
  `pNext` chain, as determined by querying
  [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)
  extended with
  [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)

Valid Usage (Implicit)

- <a href="#VUID-VkImageDrmFormatModifierListCreateInfoEXT-sType-sType"
  id="VUID-VkImageDrmFormatModifierListCreateInfoEXT-sType-sType"></a>
  VUID-VkImageDrmFormatModifierListCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_LIST_CREATE_INFO_EXT`

- <a
  href="#VUID-VkImageDrmFormatModifierListCreateInfoEXT-pDrmFormatModifiers-parameter"
  id="VUID-VkImageDrmFormatModifierListCreateInfoEXT-pDrmFormatModifiers-parameter"></a>
  VUID-VkImageDrmFormatModifierListCreateInfoEXT-pDrmFormatModifiers-parameter  
  `pDrmFormatModifiers` **must** be a valid pointer to an array of
  `drmFormatModifierCount` `uint64_t` values

- <a
  href="#VUID-VkImageDrmFormatModifierListCreateInfoEXT-drmFormatModifierCount-arraylength"
  id="VUID-VkImageDrmFormatModifierListCreateInfoEXT-drmFormatModifierCount-arraylength"></a>
  VUID-VkImageDrmFormatModifierListCreateInfoEXT-drmFormatModifierCount-arraylength  
  `drmFormatModifierCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_drm_format_modifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_drm_format_modifier.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageDrmFormatModifierListCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
