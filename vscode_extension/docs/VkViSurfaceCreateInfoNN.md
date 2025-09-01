# VkViSurfaceCreateInfoNN(3) Manual Page

## Name

VkViSurfaceCreateInfoNN - Structure specifying parameters of a newly created VI surface object



## [](#_c_specification)C Specification

The `VkViSurfaceCreateInfoNN` structure is defined as:

```c++
// Provided by VK_NN_vi_surface
typedef struct VkViSurfaceCreateInfoNN {
    VkStructureType             sType;
    const void*                 pNext;
    VkViSurfaceCreateFlagsNN    flags;
    void*                       window;
} VkViSurfaceCreateInfoNN;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `window` is the `nn`::`vi`::`NativeWindowHandle` for the `nn`::`vi`::`Layer` with which to associate the surface.

## [](#_description)Description

Valid Usage

- [](#VUID-VkViSurfaceCreateInfoNN-window-01318)VUID-VkViSurfaceCreateInfoNN-window-01318  
  `window` **must** be a valid `nn`::`vi`::`NativeWindowHandle`

Valid Usage (Implicit)

- [](#VUID-VkViSurfaceCreateInfoNN-sType-sType)VUID-VkViSurfaceCreateInfoNN-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VI_SURFACE_CREATE_INFO_NN`
- [](#VUID-VkViSurfaceCreateInfoNN-pNext-pNext)VUID-VkViSurfaceCreateInfoNN-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkViSurfaceCreateInfoNN-flags-zerobitmask)VUID-VkViSurfaceCreateInfoNN-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_NN\_vi\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NN_vi_surface.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkViSurfaceCreateFlagsNN](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViSurfaceCreateFlagsNN.html), [vkCreateViSurfaceNN](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateViSurfaceNN.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkViSurfaceCreateInfoNN).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0