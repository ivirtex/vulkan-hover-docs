# VkMacOSSurfaceCreateInfoMVK(3) Manual Page

## Name

VkMacOSSurfaceCreateInfoMVK - Structure specifying parameters of a newly
created macOS surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkMacOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMacOSSurfaceCreateInfoMVK.html)
structure is defined as:

``` c
// Provided by VK_MVK_macos_surface
typedef struct VkMacOSSurfaceCreateInfoMVK {
    VkStructureType                 sType;
    const void*                     pNext;
    VkMacOSSurfaceCreateFlagsMVK    flags;
    const void*                     pView;
} VkMacOSSurfaceCreateInfoMVK;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `pView` is a reference to either a [CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html)
  object or an `NSView` object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkMacOSSurfaceCreateInfoMVK-pView-04144"
  id="VUID-VkMacOSSurfaceCreateInfoMVK-pView-04144"></a>
  VUID-VkMacOSSurfaceCreateInfoMVK-pView-04144  
  If `pView` is a [CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html) object, it **must**
  be a valid [CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html)

- <a href="#VUID-VkMacOSSurfaceCreateInfoMVK-pView-01317"
  id="VUID-VkMacOSSurfaceCreateInfoMVK-pView-01317"></a>
  VUID-VkMacOSSurfaceCreateInfoMVK-pView-01317  
  If `pView` is an `NSView` object, it **must** be a valid `NSView`,
  **must** be backed by a `CALayer` object of type
  [CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html), and
  [vkCreateMacOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateMacOSSurfaceMVK.html) **must** be
  called on the main thread

Valid Usage (Implicit)

- <a href="#VUID-VkMacOSSurfaceCreateInfoMVK-sType-sType"
  id="VUID-VkMacOSSurfaceCreateInfoMVK-sType-sType"></a>
  VUID-VkMacOSSurfaceCreateInfoMVK-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK`

- <a href="#VUID-VkMacOSSurfaceCreateInfoMVK-pNext-pNext"
  id="VUID-VkMacOSSurfaceCreateInfoMVK-pNext-pNext"></a>
  VUID-VkMacOSSurfaceCreateInfoMVK-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMacOSSurfaceCreateInfoMVK-flags-zerobitmask"
  id="VUID-VkMacOSSurfaceCreateInfoMVK-flags-zerobitmask"></a>
  VUID-VkMacOSSurfaceCreateInfoMVK-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_MVK_macos_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MVK_macos_surface.html),
[VkMacOSSurfaceCreateFlagsMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMacOSSurfaceCreateFlagsMVK.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateMacOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateMacOSSurfaceMVK.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMacOSSurfaceCreateInfoMVK"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
