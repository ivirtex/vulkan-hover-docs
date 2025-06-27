# VkMetalSurfaceCreateInfoEXT(3) Manual Page

## Name

VkMetalSurfaceCreateInfoEXT - Structure specifying parameters of a newly
created Metal surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkMetalSurfaceCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMetalSurfaceCreateInfoEXT.html)
structure is defined as:

``` c
// Provided by VK_EXT_metal_surface
typedef struct VkMetalSurfaceCreateInfoEXT {
    VkStructureType                 sType;
    const void*                     pNext;
    VkMetalSurfaceCreateFlagsEXT    flags;
    const CAMetalLayer*             pLayer;
} VkMetalSurfaceCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `pLayer` is a reference to a [CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html) object
  representing a renderable surface.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkMetalSurfaceCreateInfoEXT-sType-sType"
  id="VUID-VkMetalSurfaceCreateInfoEXT-sType-sType"></a>
  VUID-VkMetalSurfaceCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_METAL_SURFACE_CREATE_INFO_EXT`

- <a href="#VUID-VkMetalSurfaceCreateInfoEXT-pNext-pNext"
  id="VUID-VkMetalSurfaceCreateInfoEXT-pNext-pNext"></a>
  VUID-VkMetalSurfaceCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMetalSurfaceCreateInfoEXT-flags-zerobitmask"
  id="VUID-VkMetalSurfaceCreateInfoEXT-flags-zerobitmask"></a>
  VUID-VkMetalSurfaceCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_surface.html),
[VkMetalSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMetalSurfaceCreateFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateMetalSurfaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateMetalSurfaceEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMetalSurfaceCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
