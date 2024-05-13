# VkWaylandSurfaceCreateInfoKHR(3) Manual Page

## Name

VkWaylandSurfaceCreateInfoKHR - Structure specifying parameters of a
newly created Wayland surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkWaylandSurfaceCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_wayland_surface
typedef struct VkWaylandSurfaceCreateInfoKHR {
    VkStructureType                   sType;
    const void*                       pNext;
    VkWaylandSurfaceCreateFlagsKHR    flags;
    struct wl_display*                display;
    struct wl_surface*                surface;
} VkWaylandSurfaceCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `display` and `surface` are pointers to the Wayland `wl_display` and
  `wl_surface` to associate the surface with.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkWaylandSurfaceCreateInfoKHR-display-01304"
  id="VUID-VkWaylandSurfaceCreateInfoKHR-display-01304"></a>
  VUID-VkWaylandSurfaceCreateInfoKHR-display-01304  
  `display` **must** point to a valid Wayland `wl_display`

- <a href="#VUID-VkWaylandSurfaceCreateInfoKHR-surface-01305"
  id="VUID-VkWaylandSurfaceCreateInfoKHR-surface-01305"></a>
  VUID-VkWaylandSurfaceCreateInfoKHR-surface-01305  
  `surface` **must** point to a valid Wayland `wl_surface`

Valid Usage (Implicit)

- <a href="#VUID-VkWaylandSurfaceCreateInfoKHR-sType-sType"
  id="VUID-VkWaylandSurfaceCreateInfoKHR-sType-sType"></a>
  VUID-VkWaylandSurfaceCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_WAYLAND_SURFACE_CREATE_INFO_KHR`

- <a href="#VUID-VkWaylandSurfaceCreateInfoKHR-pNext-pNext"
  id="VUID-VkWaylandSurfaceCreateInfoKHR-pNext-pNext"></a>
  VUID-VkWaylandSurfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkWaylandSurfaceCreateInfoKHR-flags-zerobitmask"
  id="VUID-VkWaylandSurfaceCreateInfoKHR-flags-zerobitmask"></a>
  VUID-VkWaylandSurfaceCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_wayland_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_wayland_surface.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkWaylandSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWaylandSurfaceCreateFlagsKHR.html),
[vkCreateWaylandSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateWaylandSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkWaylandSurfaceCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
