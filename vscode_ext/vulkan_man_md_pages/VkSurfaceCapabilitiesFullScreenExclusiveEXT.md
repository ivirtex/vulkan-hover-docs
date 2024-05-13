# VkSurfaceCapabilitiesFullScreenExclusiveEXT(3) Manual Page

## Name

VkSurfaceCapabilitiesFullScreenExclusiveEXT - Structure describing full
screen exclusive capabilities of a surface



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSurfaceCapabilitiesFullScreenExclusiveEXT` structure is defined
as:

``` c
// Provided by VK_EXT_full_screen_exclusive
typedef struct VkSurfaceCapabilitiesFullScreenExclusiveEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           fullScreenExclusiveSupported;
} VkSurfaceCapabilitiesFullScreenExclusiveEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `fullScreenExclusiveControlSupported` is a boolean describing whether
  the surface is able to make use of exclusive full-screen access.

## <a href="#_description" class="anchor"></a>Description

This structure **can** be included in the `pNext` chain of
[VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html) to determine
support for exclusive full-screen access. If
`fullScreenExclusiveSupported` is `VK_FALSE`, it indicates that
exclusive full-screen access is not obtainable for this surface.

Applications **must** not attempt to create swapchains with
`VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT` set if
`fullScreenExclusiveSupported` is `VK_FALSE`.

Valid Usage (Implicit)

- <a href="#VUID-VkSurfaceCapabilitiesFullScreenExclusiveEXT-sType-sType"
  id="VUID-VkSurfaceCapabilitiesFullScreenExclusiveEXT-sType-sType"></a>
  VUID-VkSurfaceCapabilitiesFullScreenExclusiveEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_full_screen_exclusive](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_full_screen_exclusive.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfaceCapabilitiesFullScreenExclusiveEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
