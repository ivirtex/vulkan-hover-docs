# VkScreenSurfaceCreateInfoQNX(3) Manual Page

## Name

VkScreenSurfaceCreateInfoQNX - Structure specifying parameters of a newly created QNX Screen surface object



## [](#_c_specification)C Specification

The `VkScreenSurfaceCreateInfoQNX` structure is defined as:

```c++
// Provided by VK_QNX_screen_surface
typedef struct VkScreenSurfaceCreateInfoQNX {
    VkStructureType                  sType;
    const void*                      pNext;
    VkScreenSurfaceCreateFlagsQNX    flags;
    struct _screen_context*          context;
    struct _screen_window*           window;
} VkScreenSurfaceCreateInfoQNX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `context` and `window` are QNX Screen `context` and `window` to associate the surface with.

## [](#_description)Description

Valid Usage

- [](#VUID-VkScreenSurfaceCreateInfoQNX-context-04741)VUID-VkScreenSurfaceCreateInfoQNX-context-04741  
  `context` **must** point to a valid QNX Screen `struct` \_screen\_context
- [](#VUID-VkScreenSurfaceCreateInfoQNX-window-04742)VUID-VkScreenSurfaceCreateInfoQNX-window-04742  
  `window` **must** point to a valid QNX Screen `struct` \_screen\_window

Valid Usage (Implicit)

- [](#VUID-VkScreenSurfaceCreateInfoQNX-sType-sType)VUID-VkScreenSurfaceCreateInfoQNX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SCREEN_SURFACE_CREATE_INFO_QNX`
- [](#VUID-VkScreenSurfaceCreateInfoQNX-pNext-pNext)VUID-VkScreenSurfaceCreateInfoQNX-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkScreenSurfaceCreateInfoQNX-flags-zerobitmask)VUID-VkScreenSurfaceCreateInfoQNX-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_QNX\_screen\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QNX_screen_surface.html), [VkScreenSurfaceCreateFlagsQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenSurfaceCreateFlagsQNX.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateScreenSurfaceQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateScreenSurfaceQNX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkScreenSurfaceCreateInfoQNX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0