# VkDirectFBSurfaceCreateInfoEXT(3) Manual Page

## Name

VkDirectFBSurfaceCreateInfoEXT - Structure specifying parameters of a
newly created DirectFB surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDirectFBSurfaceCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_directfb_surface
typedef struct VkDirectFBSurfaceCreateInfoEXT {
    VkStructureType                    sType;
    const void*                        pNext;
    VkDirectFBSurfaceCreateFlagsEXT    flags;
    IDirectFB*                         dfb;
    IDirectFBSurface*                  surface;
} VkDirectFBSurfaceCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `dfb` is a pointer to the `IDirectFB` main interface of DirectFB.

- `surface` is a pointer to a `IDirectFBSurface` surface interface.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDirectFBSurfaceCreateInfoEXT-dfb-04117"
  id="VUID-VkDirectFBSurfaceCreateInfoEXT-dfb-04117"></a>
  VUID-VkDirectFBSurfaceCreateInfoEXT-dfb-04117  
  `dfb` **must** point to a valid DirectFB `IDirectFB`

- <a href="#VUID-VkDirectFBSurfaceCreateInfoEXT-surface-04118"
  id="VUID-VkDirectFBSurfaceCreateInfoEXT-surface-04118"></a>
  VUID-VkDirectFBSurfaceCreateInfoEXT-surface-04118  
  `surface` **must** point to a valid DirectFB `IDirectFBSurface`

Valid Usage (Implicit)

- <a href="#VUID-VkDirectFBSurfaceCreateInfoEXT-sType-sType"
  id="VUID-VkDirectFBSurfaceCreateInfoEXT-sType-sType"></a>
  VUID-VkDirectFBSurfaceCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DIRECTFB_SURFACE_CREATE_INFO_EXT`

- <a href="#VUID-VkDirectFBSurfaceCreateInfoEXT-pNext-pNext"
  id="VUID-VkDirectFBSurfaceCreateInfoEXT-pNext-pNext"></a>
  VUID-VkDirectFBSurfaceCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDirectFBSurfaceCreateInfoEXT-flags-zerobitmask"
  id="VUID-VkDirectFBSurfaceCreateInfoEXT-flags-zerobitmask"></a>
  VUID-VkDirectFBSurfaceCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_directfb_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_directfb_surface.html),
[VkDirectFBSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectFBSurfaceCreateFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateDirectFBSurfaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDirectFBSurfaceEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDirectFBSurfaceCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
