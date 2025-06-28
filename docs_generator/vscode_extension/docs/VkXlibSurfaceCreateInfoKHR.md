# VkXlibSurfaceCreateInfoKHR(3) Manual Page

## Name

VkXlibSurfaceCreateInfoKHR - Structure specifying parameters of a newly created Xlib surface object



## [](#_c_specification)C Specification

The `VkXlibSurfaceCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_xlib_surface
typedef struct VkXlibSurfaceCreateInfoKHR {
    VkStructureType                sType;
    const void*                    pNext;
    VkXlibSurfaceCreateFlagsKHR    flags;
    Display*                       dpy;
    Window                         window;
} VkXlibSurfaceCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `dpy` is a pointer to an Xlib `Display` connection to the X server.
- `window` is an Xlib `Window` to associate the surface with.

## [](#_description)Description

Valid Usage

- [](#VUID-VkXlibSurfaceCreateInfoKHR-dpy-01313)VUID-VkXlibSurfaceCreateInfoKHR-dpy-01313  
  `dpy` **must** point to a valid Xlib `Display`
- [](#VUID-VkXlibSurfaceCreateInfoKHR-window-01314)VUID-VkXlibSurfaceCreateInfoKHR-window-01314  
  `window` **must** be a valid Xlib `Window`

Valid Usage (Implicit)

- [](#VUID-VkXlibSurfaceCreateInfoKHR-sType-sType)VUID-VkXlibSurfaceCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_XLIB_SURFACE_CREATE_INFO_KHR`
- [](#VUID-VkXlibSurfaceCreateInfoKHR-pNext-pNext)VUID-VkXlibSurfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkXlibSurfaceCreateInfoKHR-flags-zerobitmask)VUID-VkXlibSurfaceCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_KHR\_xlib\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_xlib_surface.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkXlibSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkXlibSurfaceCreateFlagsKHR.html), [vkCreateXlibSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateXlibSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkXlibSurfaceCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0