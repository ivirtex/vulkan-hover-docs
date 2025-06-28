# VkIOSSurfaceCreateInfoMVK(3) Manual Page

## Name

VkIOSSurfaceCreateInfoMVK - Structure specifying parameters of a newly created iOS surface object



## [](#_c_specification)C Specification

The [VkIOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIOSSurfaceCreateInfoMVK.html) structure is defined as:

```c++
// Provided by VK_MVK_ios_surface
typedef struct VkIOSSurfaceCreateInfoMVK {
    VkStructureType               sType;
    const void*                   pNext;
    VkIOSSurfaceCreateFlagsMVK    flags;
    const void*                   pView;
} VkIOSSurfaceCreateInfoMVK;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `pView` is a reference to either a [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html) object or a `UIView` object.

## [](#_description)Description

Valid Usage

- [](#VUID-VkIOSSurfaceCreateInfoMVK-pView-04143)VUID-VkIOSSurfaceCreateInfoMVK-pView-04143  
  If `pView` is a [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html) object, it **must** be a valid [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html)
- [](#VUID-VkIOSSurfaceCreateInfoMVK-pView-01316)VUID-VkIOSSurfaceCreateInfoMVK-pView-01316  
  If `pView` is a `UIView` object, it **must** be a valid `UIView`, **must** be backed by a `CALayer` object of type [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html), and [vkCreateIOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIOSSurfaceMVK.html) **must** be called on the main thread

Valid Usage (Implicit)

- [](#VUID-VkIOSSurfaceCreateInfoMVK-sType-sType)VUID-VkIOSSurfaceCreateInfoMVK-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IOS_SURFACE_CREATE_INFO_MVK`
- [](#VUID-VkIOSSurfaceCreateInfoMVK-pNext-pNext)VUID-VkIOSSurfaceCreateInfoMVK-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkIOSSurfaceCreateInfoMVK-flags-zerobitmask)VUID-VkIOSSurfaceCreateInfoMVK-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_MVK\_ios\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_MVK_ios_surface.html), [VkIOSSurfaceCreateFlagsMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIOSSurfaceCreateFlagsMVK.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateIOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIOSSurfaceMVK.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIOSSurfaceCreateInfoMVK)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0