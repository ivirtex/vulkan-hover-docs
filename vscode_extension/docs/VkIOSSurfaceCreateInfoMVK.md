# VkIOSSurfaceCreateInfoMVK(3) Manual Page

## Name

VkIOSSurfaceCreateInfoMVK - Structure specifying parameters of a newly
created iOS surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkIOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIOSSurfaceCreateInfoMVK.html)
structure is defined as:

``` c
// Provided by VK_MVK_ios_surface
typedef struct VkIOSSurfaceCreateInfoMVK {
    VkStructureType               sType;
    const void*                   pNext;
    VkIOSSurfaceCreateFlagsMVK    flags;
    const void*                   pView;
} VkIOSSurfaceCreateInfoMVK;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `pView` is a reference to either a [CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html)
  object or a `UIView` object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkIOSSurfaceCreateInfoMVK-pView-04143"
  id="VUID-VkIOSSurfaceCreateInfoMVK-pView-04143"></a>
  VUID-VkIOSSurfaceCreateInfoMVK-pView-04143  
  If `pView` is a [CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html) object, it **must**
  be a valid [CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html)

- <a href="#VUID-VkIOSSurfaceCreateInfoMVK-pView-01316"
  id="VUID-VkIOSSurfaceCreateInfoMVK-pView-01316"></a>
  VUID-VkIOSSurfaceCreateInfoMVK-pView-01316  
  If `pView` is a `UIView` object, it **must** be a valid `UIView`,
  **must** be backed by a `CALayer` object of type
  [CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html), and
  [vkCreateIOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateIOSSurfaceMVK.html) **must** be called
  on the main thread

Valid Usage (Implicit)

- <a href="#VUID-VkIOSSurfaceCreateInfoMVK-sType-sType"
  id="VUID-VkIOSSurfaceCreateInfoMVK-sType-sType"></a>
  VUID-VkIOSSurfaceCreateInfoMVK-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IOS_SURFACE_CREATE_INFO_MVK`

- <a href="#VUID-VkIOSSurfaceCreateInfoMVK-pNext-pNext"
  id="VUID-VkIOSSurfaceCreateInfoMVK-pNext-pNext"></a>
  VUID-VkIOSSurfaceCreateInfoMVK-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkIOSSurfaceCreateInfoMVK-flags-zerobitmask"
  id="VUID-VkIOSSurfaceCreateInfoMVK-flags-zerobitmask"></a>
  VUID-VkIOSSurfaceCreateInfoMVK-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_MVK_ios_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MVK_ios_surface.html),
[VkIOSSurfaceCreateFlagsMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIOSSurfaceCreateFlagsMVK.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateIOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateIOSSurfaceMVK.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkIOSSurfaceCreateInfoMVK"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
