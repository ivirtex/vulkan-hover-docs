# VkImageDrmFormatModifierListCreateInfoEXT(3) Manual Page

## Name

VkImageDrmFormatModifierListCreateInfoEXT - Specify that an image must be created with a DRM format modifier from the provided list



## [](#_c_specification)C Specification

If the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) includes a [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html) structure, then the image will be created with one of the [Linux DRM format modifiers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-drm-format-modifier) listed in the structure. The choice of modifier is implementation-dependent.

The [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html) structure is defined as:

```c++
// Provided by VK_EXT_image_drm_format_modifier
typedef struct VkImageDrmFormatModifierListCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           drmFormatModifierCount;
    const uint64_t*    pDrmFormatModifiers;
} VkImageDrmFormatModifierListCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `drmFormatModifierCount` is the length of the `pDrmFormatModifiers` array.
- `pDrmFormatModifiers` is a pointer to an array of *Linux DRM format modifiers*.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImageDrmFormatModifierListCreateInfoEXT-pDrmFormatModifiers-02263)VUID-VkImageDrmFormatModifierListCreateInfoEXT-pDrmFormatModifiers-02263  
  Each *modifier* in `pDrmFormatModifiers` **must** be compatible with the parameters in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) and its `pNext` chain, as determined by querying [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html) extended with [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)

Valid Usage (Implicit)

- [](#VUID-VkImageDrmFormatModifierListCreateInfoEXT-sType-sType)VUID-VkImageDrmFormatModifierListCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_LIST_CREATE_INFO_EXT`
- [](#VUID-VkImageDrmFormatModifierListCreateInfoEXT-pDrmFormatModifiers-parameter)VUID-VkImageDrmFormatModifierListCreateInfoEXT-pDrmFormatModifiers-parameter  
  `pDrmFormatModifiers` **must** be a valid pointer to an array of `drmFormatModifierCount` `uint64_t` values
- [](#VUID-VkImageDrmFormatModifierListCreateInfoEXT-drmFormatModifierCount-arraylength)VUID-VkImageDrmFormatModifierListCreateInfoEXT-drmFormatModifierCount-arraylength  
  `drmFormatModifierCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_image\_drm\_format\_modifier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_image_drm_format_modifier.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageDrmFormatModifierListCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0