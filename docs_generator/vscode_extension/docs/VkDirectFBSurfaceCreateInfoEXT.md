# VkDirectFBSurfaceCreateInfoEXT(3) Manual Page

## Name

VkDirectFBSurfaceCreateInfoEXT - Structure specifying parameters of a newly created DirectFB surface object



## [](#_c_specification)C Specification

The `VkDirectFBSurfaceCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_directfb_surface
typedef struct VkDirectFBSurfaceCreateInfoEXT {
    VkStructureType                    sType;
    const void*                        pNext;
    VkDirectFBSurfaceCreateFlagsEXT    flags;
    IDirectFB*                         dfb;
    IDirectFBSurface*                  surface;
} VkDirectFBSurfaceCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `dfb` is a pointer to the `IDirectFB` main interface of DirectFB.
- `surface` is a pointer to a `IDirectFBSurface` surface interface.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDirectFBSurfaceCreateInfoEXT-dfb-04117)VUID-VkDirectFBSurfaceCreateInfoEXT-dfb-04117  
  `dfb` **must** point to a valid DirectFB `IDirectFB`
- [](#VUID-VkDirectFBSurfaceCreateInfoEXT-surface-04118)VUID-VkDirectFBSurfaceCreateInfoEXT-surface-04118  
  `surface` **must** point to a valid DirectFB `IDirectFBSurface`

Valid Usage (Implicit)

- [](#VUID-VkDirectFBSurfaceCreateInfoEXT-sType-sType)VUID-VkDirectFBSurfaceCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DIRECTFB_SURFACE_CREATE_INFO_EXT`
- [](#VUID-VkDirectFBSurfaceCreateInfoEXT-pNext-pNext)VUID-VkDirectFBSurfaceCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDirectFBSurfaceCreateInfoEXT-flags-zerobitmask)VUID-VkDirectFBSurfaceCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_EXT\_directfb\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_directfb_surface.html), [VkDirectFBSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectFBSurfaceCreateFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateDirectFBSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDirectFBSurfaceEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDirectFBSurfaceCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0