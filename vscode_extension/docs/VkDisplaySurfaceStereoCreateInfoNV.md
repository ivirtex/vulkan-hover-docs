# VkDisplaySurfaceStereoCreateInfoNV(3) Manual Page

## Name

VkDisplaySurfaceStereoCreateInfoNV - Structure specifying stereo parameters of a newly created display plane surface object



## [](#_c_specification)C Specification

The `VkDisplaySurfaceStereoCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_display_stereo
typedef struct VkDisplaySurfaceStereoCreateInfoNV {
    VkStructureType                 sType;
    const void*                     pNext;
    VkDisplaySurfaceStereoTypeNV    stereoType;
} VkDisplaySurfaceStereoCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`stereoType` is a [VkDisplaySurfaceStereoTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceStereoTypeNV.html) value specifying the type of 3D stereo presentation the display will be configured for.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDisplaySurfaceStereoCreateInfoNV-sType-sType)VUID-VkDisplaySurfaceStereoCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_SURFACE_STEREO_CREATE_INFO_NV`
- [](#VUID-VkDisplaySurfaceStereoCreateInfoNV-stereoType-parameter)VUID-VkDisplaySurfaceStereoCreateInfoNV-stereoType-parameter  
  `stereoType` **must** be a valid [VkDisplaySurfaceStereoTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceStereoTypeNV.html) value

## [](#_see_also)See Also

[VK\_NV\_display\_stereo](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_display_stereo.html), [VkDisplaySurfaceStereoTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceStereoTypeNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplaySurfaceStereoCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0