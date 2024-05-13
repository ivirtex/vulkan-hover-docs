# VkViSurfaceCreateInfoNN(3) Manual Page

## Name

VkViSurfaceCreateInfoNN - Structure specifying parameters of a newly
created VI surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkViSurfaceCreateInfoNN` structure is defined as:

``` c
// Provided by VK_NN_vi_surface
typedef struct VkViSurfaceCreateInfoNN {
    VkStructureType             sType;
    const void*                 pNext;
    VkViSurfaceCreateFlagsNN    flags;
    void*                       window;
} VkViSurfaceCreateInfoNN;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `window` is the `nn`::`vi`::`NativeWindowHandle` for the
  `nn`::`vi`::`Layer` with which to associate the surface.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkViSurfaceCreateInfoNN-window-01318"
  id="VUID-VkViSurfaceCreateInfoNN-window-01318"></a>
  VUID-VkViSurfaceCreateInfoNN-window-01318  
  `window` **must** be a valid `nn`::`vi`::`NativeWindowHandle`

Valid Usage (Implicit)

- <a href="#VUID-VkViSurfaceCreateInfoNN-sType-sType"
  id="VUID-VkViSurfaceCreateInfoNN-sType-sType"></a>
  VUID-VkViSurfaceCreateInfoNN-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VI_SURFACE_CREATE_INFO_NN`

- <a href="#VUID-VkViSurfaceCreateInfoNN-pNext-pNext"
  id="VUID-VkViSurfaceCreateInfoNN-pNext-pNext"></a>
  VUID-VkViSurfaceCreateInfoNN-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkViSurfaceCreateInfoNN-flags-zerobitmask"
  id="VUID-VkViSurfaceCreateInfoNN-flags-zerobitmask"></a>
  VUID-VkViSurfaceCreateInfoNN-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NN_vi_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NN_vi_surface.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkViSurfaceCreateFlagsNN](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViSurfaceCreateFlagsNN.html),
[vkCreateViSurfaceNN](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateViSurfaceNN.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkViSurfaceCreateInfoNN"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
