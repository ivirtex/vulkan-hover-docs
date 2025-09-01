# VkHeadlessSurfaceCreateInfoEXT(3) Manual Page

## Name

VkHeadlessSurfaceCreateInfoEXT - Structure specifying parameters of a newly created headless surface object



## [](#_c_specification)C Specification

The `VkHeadlessSurfaceCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_headless_surface
typedef struct VkHeadlessSurfaceCreateInfoEXT {
    VkStructureType                    sType;
    const void*                        pNext;
    VkHeadlessSurfaceCreateFlagsEXT    flags;
} VkHeadlessSurfaceCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkHeadlessSurfaceCreateInfoEXT-sType-sType)VUID-VkHeadlessSurfaceCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_HEADLESS_SURFACE_CREATE_INFO_EXT`
- [](#VUID-VkHeadlessSurfaceCreateInfoEXT-pNext-pNext)VUID-VkHeadlessSurfaceCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkHeadlessSurfaceCreateInfoEXT-flags-zerobitmask)VUID-VkHeadlessSurfaceCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_EXT\_headless\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_headless_surface.html), [VkHeadlessSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHeadlessSurfaceCreateFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateHeadlessSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateHeadlessSurfaceEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkHeadlessSurfaceCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0