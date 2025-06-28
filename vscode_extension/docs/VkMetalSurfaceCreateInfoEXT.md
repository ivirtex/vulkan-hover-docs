# VkMetalSurfaceCreateInfoEXT(3) Manual Page

## Name

VkMetalSurfaceCreateInfoEXT - Structure specifying parameters of a newly created Metal surface object



## [](#_c_specification)C Specification

The [VkMetalSurfaceCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMetalSurfaceCreateInfoEXT.html) structure is defined as:

```c++
// Provided by VK_EXT_metal_surface
typedef struct VkMetalSurfaceCreateInfoEXT {
    VkStructureType                 sType;
    const void*                     pNext;
    VkMetalSurfaceCreateFlagsEXT    flags;
    const CAMetalLayer*             pLayer;
} VkMetalSurfaceCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `pLayer` is a reference to a [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html) object representing a renderable surface.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkMetalSurfaceCreateInfoEXT-sType-sType)VUID-VkMetalSurfaceCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT`
- [](#VUID-VkMetalSurfaceCreateInfoEXT-pNext-pNext)VUID-VkMetalSurfaceCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMetalSurfaceCreateInfoEXT-flags-zerobitmask)VUID-VkMetalSurfaceCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_EXT\_metal\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_surface.html), [VkMetalSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMetalSurfaceCreateFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateMetalSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMetalSurfaceEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMetalSurfaceCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0