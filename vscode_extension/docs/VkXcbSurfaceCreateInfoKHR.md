# VkXcbSurfaceCreateInfoKHR(3) Manual Page

## Name

VkXcbSurfaceCreateInfoKHR - Structure specifying parameters of a newly created Xcb surface object



## [](#_c_specification)C Specification

The `VkXcbSurfaceCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_xcb_surface
typedef struct VkXcbSurfaceCreateInfoKHR {
    VkStructureType               sType;
    const void*                   pNext;
    VkXcbSurfaceCreateFlagsKHR    flags;
    xcb_connection_t*             connection;
    xcb_window_t                  window;
} VkXcbSurfaceCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `connection` is a pointer to an `xcb_connection_t` to the X server.
- `window` is the `xcb_window_t` for the X11 window to associate the surface with.

## [](#_description)Description

Valid Usage

- [](#VUID-VkXcbSurfaceCreateInfoKHR-connection-01310)VUID-VkXcbSurfaceCreateInfoKHR-connection-01310  
  `connection` **must** point to a valid X11 `xcb_connection_t`
- [](#VUID-VkXcbSurfaceCreateInfoKHR-window-01311)VUID-VkXcbSurfaceCreateInfoKHR-window-01311  
  `window` **must** be a valid X11 `xcb_window_t`

Valid Usage (Implicit)

- [](#VUID-VkXcbSurfaceCreateInfoKHR-sType-sType)VUID-VkXcbSurfaceCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR`
- [](#VUID-VkXcbSurfaceCreateInfoKHR-pNext-pNext)VUID-VkXcbSurfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkXcbSurfaceCreateInfoKHR-flags-zerobitmask)VUID-VkXcbSurfaceCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_KHR\_xcb\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_xcb_surface.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkXcbSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkXcbSurfaceCreateFlagsKHR.html), [vkCreateXcbSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateXcbSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkXcbSurfaceCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0