# VkMacOSSurfaceCreateInfoMVK(3) Manual Page

## Name

VkMacOSSurfaceCreateInfoMVK - Structure specifying parameters of a newly created macOS surface object



## [](#_c_specification)C Specification

The [VkMacOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMacOSSurfaceCreateInfoMVK.html) structure is defined as:

```c++
// Provided by VK_MVK_macos_surface
typedef struct VkMacOSSurfaceCreateInfoMVK {
    VkStructureType                 sType;
    const void*                     pNext;
    VkMacOSSurfaceCreateFlagsMVK    flags;
    const void*                     pView;
} VkMacOSSurfaceCreateInfoMVK;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `pView` is a reference to either a [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html) object or an `NSView` object.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMacOSSurfaceCreateInfoMVK-pView-04144)VUID-VkMacOSSurfaceCreateInfoMVK-pView-04144  
  If `pView` is a [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html) object, it **must** be a valid [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html)
- [](#VUID-VkMacOSSurfaceCreateInfoMVK-pView-01317)VUID-VkMacOSSurfaceCreateInfoMVK-pView-01317  
  If `pView` is an `NSView` object, it **must** be a valid `NSView`, **must** be backed by a `CALayer` object of type [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html), and [vkCreateMacOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMacOSSurfaceMVK.html) **must** be called on the main thread

Valid Usage (Implicit)

- [](#VUID-VkMacOSSurfaceCreateInfoMVK-sType-sType)VUID-VkMacOSSurfaceCreateInfoMVK-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK`
- [](#VUID-VkMacOSSurfaceCreateInfoMVK-pNext-pNext)VUID-VkMacOSSurfaceCreateInfoMVK-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMacOSSurfaceCreateInfoMVK-flags-zerobitmask)VUID-VkMacOSSurfaceCreateInfoMVK-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_MVK\_macos\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_MVK_macos_surface.html), [VkMacOSSurfaceCreateFlagsMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMacOSSurfaceCreateFlagsMVK.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateMacOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMacOSSurfaceMVK.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMacOSSurfaceCreateInfoMVK).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0