# VkWaylandSurfaceCreateInfoKHR(3) Manual Page

## Name

VkWaylandSurfaceCreateInfoKHR - Structure specifying parameters of a newly created Wayland surface object



## [](#_c_specification)C Specification

The `VkWaylandSurfaceCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_wayland_surface
typedef struct VkWaylandSurfaceCreateInfoKHR {
    VkStructureType                   sType;
    const void*                       pNext;
    VkWaylandSurfaceCreateFlagsKHR    flags;
    struct wl_display*                display;
    struct wl_surface*                surface;
} VkWaylandSurfaceCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `display` and `surface` are pointers to the Wayland `wl_display` and `wl_surface` to associate the surface with.

## [](#_description)Description

Valid Usage

- [](#VUID-VkWaylandSurfaceCreateInfoKHR-display-01304)VUID-VkWaylandSurfaceCreateInfoKHR-display-01304  
  `display` **must** point to a valid Wayland `wl_display`
- [](#VUID-VkWaylandSurfaceCreateInfoKHR-surface-01305)VUID-VkWaylandSurfaceCreateInfoKHR-surface-01305  
  `surface` **must** point to a valid Wayland `wl_surface`

Valid Usage (Implicit)

- [](#VUID-VkWaylandSurfaceCreateInfoKHR-sType-sType)VUID-VkWaylandSurfaceCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WAYLAND_SURFACE_CREATE_INFO_KHR`
- [](#VUID-VkWaylandSurfaceCreateInfoKHR-pNext-pNext)VUID-VkWaylandSurfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkWaylandSurfaceCreateInfoKHR-flags-zerobitmask)VUID-VkWaylandSurfaceCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_KHR\_wayland\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_wayland_surface.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkWaylandSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWaylandSurfaceCreateFlagsKHR.html), [vkCreateWaylandSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateWaylandSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkWaylandSurfaceCreateInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0