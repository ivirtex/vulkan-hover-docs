# VkScreenSurfaceCreateInfoQNX(3) Manual Page

## Name

VkScreenSurfaceCreateInfoQNX - Structure specifying parameters of a
newly created QNX Screen surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkScreenSurfaceCreateInfoQNX` structure is defined as:

``` c
// Provided by VK_QNX_screen_surface
typedef struct VkScreenSurfaceCreateInfoQNX {
    VkStructureType                  sType;
    const void*                      pNext;
    VkScreenSurfaceCreateFlagsQNX    flags;
    struct _screen_context*          context;
    struct _screen_window*           window;
} VkScreenSurfaceCreateInfoQNX;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `context` and `window` are QNX Screen `context` and `window` to
  associate the surface with.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkScreenSurfaceCreateInfoQNX-context-04741"
  id="VUID-VkScreenSurfaceCreateInfoQNX-context-04741"></a>
  VUID-VkScreenSurfaceCreateInfoQNX-context-04741  
  `context` **must** point to a valid QNX Screen `struct`
  \_screen_context

- <a href="#VUID-VkScreenSurfaceCreateInfoQNX-window-04742"
  id="VUID-VkScreenSurfaceCreateInfoQNX-window-04742"></a>
  VUID-VkScreenSurfaceCreateInfoQNX-window-04742  
  `window` **must** point to a valid QNX Screen `struct` \_screen_window

Valid Usage (Implicit)

- <a href="#VUID-VkScreenSurfaceCreateInfoQNX-sType-sType"
  id="VUID-VkScreenSurfaceCreateInfoQNX-sType-sType"></a>
  VUID-VkScreenSurfaceCreateInfoQNX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SCREEN_SURFACE_CREATE_INFO_QNX`

- <a href="#VUID-VkScreenSurfaceCreateInfoQNX-pNext-pNext"
  id="VUID-VkScreenSurfaceCreateInfoQNX-pNext-pNext"></a>
  VUID-VkScreenSurfaceCreateInfoQNX-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkScreenSurfaceCreateInfoQNX-flags-zerobitmask"
  id="VUID-VkScreenSurfaceCreateInfoQNX-flags-zerobitmask"></a>
  VUID-VkScreenSurfaceCreateInfoQNX-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QNX_screen_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QNX_screen_surface.html),
[VkScreenSurfaceCreateFlagsQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenSurfaceCreateFlagsQNX.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateScreenSurfaceQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateScreenSurfaceQNX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkScreenSurfaceCreateInfoQNX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
