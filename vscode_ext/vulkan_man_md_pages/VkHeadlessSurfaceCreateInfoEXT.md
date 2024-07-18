# VkHeadlessSurfaceCreateInfoEXT(3) Manual Page

## Name

VkHeadlessSurfaceCreateInfoEXT - Structure specifying parameters of a
newly created headless surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkHeadlessSurfaceCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_headless_surface
typedef struct VkHeadlessSurfaceCreateInfoEXT {
    VkStructureType                    sType;
    const void*                        pNext;
    VkHeadlessSurfaceCreateFlagsEXT    flags;
} VkHeadlessSurfaceCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkHeadlessSurfaceCreateInfoEXT-sType-sType"
  id="VUID-VkHeadlessSurfaceCreateInfoEXT-sType-sType"></a>
  VUID-VkHeadlessSurfaceCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_HEADLESS_SURFACE_CREATE_INFO_EXT`

- <a href="#VUID-VkHeadlessSurfaceCreateInfoEXT-pNext-pNext"
  id="VUID-VkHeadlessSurfaceCreateInfoEXT-pNext-pNext"></a>
  VUID-VkHeadlessSurfaceCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkHeadlessSurfaceCreateInfoEXT-flags-zerobitmask"
  id="VUID-VkHeadlessSurfaceCreateInfoEXT-flags-zerobitmask"></a>
  VUID-VkHeadlessSurfaceCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_headless_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_headless_surface.html),
[VkHeadlessSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHeadlessSurfaceCreateFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateHeadlessSurfaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateHeadlessSurfaceEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkHeadlessSurfaceCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
