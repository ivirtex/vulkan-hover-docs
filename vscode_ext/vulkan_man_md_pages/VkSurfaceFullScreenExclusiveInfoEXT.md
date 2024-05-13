# VkSurfaceFullScreenExclusiveInfoEXT(3) Manual Page

## Name

VkSurfaceFullScreenExclusiveInfoEXT - Structure specifying the preferred
full-screen transition behavior



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) includes a
`VkSurfaceFullScreenExclusiveInfoEXT` structure, then that structure
specifies the application’s preferred full-screen transition behavior.

The `VkSurfaceFullScreenExclusiveInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_full_screen_exclusive
typedef struct VkSurfaceFullScreenExclusiveInfoEXT {
    VkStructureType             sType;
    void*                       pNext;
    VkFullScreenExclusiveEXT    fullScreenExclusive;
} VkSurfaceFullScreenExclusiveInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `fullScreenExclusive` is a
  [VkFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFullScreenExclusiveEXT.html) value
  specifying the preferred full-screen transition behavior.

## <a href="#_description" class="anchor"></a>Description

If this structure is not present, `fullScreenExclusive` is considered to
be `VK_FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT`.

Valid Usage (Implicit)

- <a href="#VUID-VkSurfaceFullScreenExclusiveInfoEXT-sType-sType"
  id="VUID-VkSurfaceFullScreenExclusiveInfoEXT-sType-sType"></a>
  VUID-VkSurfaceFullScreenExclusiveInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT`

- <a
  href="#VUID-VkSurfaceFullScreenExclusiveInfoEXT-fullScreenExclusive-parameter"
  id="VUID-VkSurfaceFullScreenExclusiveInfoEXT-fullScreenExclusive-parameter"></a>
  VUID-VkSurfaceFullScreenExclusiveInfoEXT-fullScreenExclusive-parameter  
  `fullScreenExclusive` **must** be a valid
  [VkFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFullScreenExclusiveEXT.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_full_screen_exclusive](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_full_screen_exclusive.html),
[VkFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFullScreenExclusiveEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfaceFullScreenExclusiveInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
