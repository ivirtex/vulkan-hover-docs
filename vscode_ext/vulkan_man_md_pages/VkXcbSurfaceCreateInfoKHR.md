# VkXcbSurfaceCreateInfoKHR(3) Manual Page

## Name

VkXcbSurfaceCreateInfoKHR - Structure specifying parameters of a newly
created Xcb surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkXcbSurfaceCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_xcb_surface
typedef struct VkXcbSurfaceCreateInfoKHR {
    VkStructureType               sType;
    const void*                   pNext;
    VkXcbSurfaceCreateFlagsKHR    flags;
    xcb_connection_t*             connection;
    xcb_window_t                  window;
} VkXcbSurfaceCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `connection` is a pointer to an `xcb_connection_t` to the X server.

- `window` is the `xcb_window_t` for the X11 window to associate the
  surface with.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkXcbSurfaceCreateInfoKHR-connection-01310"
  id="VUID-VkXcbSurfaceCreateInfoKHR-connection-01310"></a>
  VUID-VkXcbSurfaceCreateInfoKHR-connection-01310  
  `connection` **must** point to a valid X11 `xcb_connection_t`

- <a href="#VUID-VkXcbSurfaceCreateInfoKHR-window-01311"
  id="VUID-VkXcbSurfaceCreateInfoKHR-window-01311"></a>
  VUID-VkXcbSurfaceCreateInfoKHR-window-01311  
  `window` **must** be a valid X11 `xcb_window_t`

Valid Usage (Implicit)

- <a href="#VUID-VkXcbSurfaceCreateInfoKHR-sType-sType"
  id="VUID-VkXcbSurfaceCreateInfoKHR-sType-sType"></a>
  VUID-VkXcbSurfaceCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_XCB_SURFACE_CREATE_INFO_KHR`

- <a href="#VUID-VkXcbSurfaceCreateInfoKHR-pNext-pNext"
  id="VUID-VkXcbSurfaceCreateInfoKHR-pNext-pNext"></a>
  VUID-VkXcbSurfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkXcbSurfaceCreateInfoKHR-flags-zerobitmask"
  id="VUID-VkXcbSurfaceCreateInfoKHR-flags-zerobitmask"></a>
  VUID-VkXcbSurfaceCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_xcb_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_xcb_surface.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkXcbSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXcbSurfaceCreateFlagsKHR.html),
[vkCreateXcbSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateXcbSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkXcbSurfaceCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
