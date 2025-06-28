# VkSurfaceFullScreenExclusiveInfoEXT(3) Manual Page

## Name

VkSurfaceFullScreenExclusiveInfoEXT - Structure specifying the preferred full-screen transition behavior



## [](#_c_specification)C Specification

If the `pNext` chain of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) includes a `VkSurfaceFullScreenExclusiveInfoEXT` structure, then that structure specifies the application’s preferred full-screen transition behavior.

The `VkSurfaceFullScreenExclusiveInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_full_screen_exclusive
typedef struct VkSurfaceFullScreenExclusiveInfoEXT {
    VkStructureType             sType;
    void*                       pNext;
    VkFullScreenExclusiveEXT    fullScreenExclusive;
} VkSurfaceFullScreenExclusiveInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `fullScreenExclusive` is a [VkFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFullScreenExclusiveEXT.html) value specifying the preferred full-screen transition behavior.

## [](#_description)Description

If this structure is not present, `fullScreenExclusive` is considered to be `VK_FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT`.

Valid Usage (Implicit)

- [](#VUID-VkSurfaceFullScreenExclusiveInfoEXT-sType-sType)VUID-VkSurfaceFullScreenExclusiveInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT`
- [](#VUID-VkSurfaceFullScreenExclusiveInfoEXT-fullScreenExclusive-parameter)VUID-VkSurfaceFullScreenExclusiveInfoEXT-fullScreenExclusive-parameter  
  `fullScreenExclusive` **must** be a valid [VkFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFullScreenExclusiveEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_full\_screen\_exclusive](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_full_screen_exclusive.html), [VkFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFullScreenExclusiveEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfaceFullScreenExclusiveInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0