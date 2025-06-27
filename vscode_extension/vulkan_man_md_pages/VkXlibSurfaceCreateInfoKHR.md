# VkXlibSurfaceCreateInfoKHR(3) Manual Page

## Name

VkXlibSurfaceCreateInfoKHR - Structure specifying parameters of a newly
created Xlib surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkXlibSurfaceCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_xlib_surface
typedef struct VkXlibSurfaceCreateInfoKHR {
    VkStructureType                sType;
    const void*                    pNext;
    VkXlibSurfaceCreateFlagsKHR    flags;
    Display*                       dpy;
    Window                         window;
} VkXlibSurfaceCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `dpy` is a pointer to an Xlib `Display` connection to the X server.

- `window` is an Xlib `Window` to associate the surface with.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkXlibSurfaceCreateInfoKHR-dpy-01313"
  id="VUID-VkXlibSurfaceCreateInfoKHR-dpy-01313"></a>
  VUID-VkXlibSurfaceCreateInfoKHR-dpy-01313  
  `dpy` **must** point to a valid Xlib `Display`

- <a href="#VUID-VkXlibSurfaceCreateInfoKHR-window-01314"
  id="VUID-VkXlibSurfaceCreateInfoKHR-window-01314"></a>
  VUID-VkXlibSurfaceCreateInfoKHR-window-01314  
  `window` **must** be a valid Xlib `Window`

Valid Usage (Implicit)

- <a href="#VUID-VkXlibSurfaceCreateInfoKHR-sType-sType"
  id="VUID-VkXlibSurfaceCreateInfoKHR-sType-sType"></a>
  VUID-VkXlibSurfaceCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_XLIB_SURFACE_CREATE_INFO_KHR`

- <a href="#VUID-VkXlibSurfaceCreateInfoKHR-pNext-pNext"
  id="VUID-VkXlibSurfaceCreateInfoKHR-pNext-pNext"></a>
  VUID-VkXlibSurfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkXlibSurfaceCreateInfoKHR-flags-zerobitmask"
  id="VUID-VkXlibSurfaceCreateInfoKHR-flags-zerobitmask"></a>
  VUID-VkXlibSurfaceCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_xlib_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_xlib_surface.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkXlibSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXlibSurfaceCreateFlagsKHR.html),
[vkCreateXlibSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateXlibSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkXlibSurfaceCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
