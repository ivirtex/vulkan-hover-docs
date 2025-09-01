# vkCreateMacOSSurfaceMVK(3) Manual Page

## Name

vkCreateMacOSSurfaceMVK - Create a VkSurfaceKHR object for a macOS NSView



## [](#_c_specification)C Specification

To create a `VkSurfaceKHR` object for a macOS `NSView` or [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html), call:

```c++
// Provided by VK_MVK_macos_surface
VkResult vkCreateMacOSSurfaceMVK(
    VkInstance                                  instance,
    const VkMacOSSurfaceCreateInfoMVK*          pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## [](#_parameters)Parameters

- `instance` is the instance with which to associate the surface.
- `pCreateInfo` is a pointer to a [VkMacOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMacOSSurfaceCreateInfoMVK.html) structure containing parameters affecting the creation of the surface object.
- `pAllocator` is the allocator used for host memory allocated for the surface object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).
- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle in which the created surface object is returned.

## [](#_description)Description

Note

The `vkCreateMacOSSurfaceMVK` function is considered deprecated and has been superseded by [vkCreateMetalSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMetalSurfaceEXT.html) from the `VK_EXT_metal_surface` extension.

Valid Usage (Implicit)

- [](#VUID-vkCreateMacOSSurfaceMVK-instance-parameter)VUID-vkCreateMacOSSurfaceMVK-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkCreateMacOSSurfaceMVK-pCreateInfo-parameter)VUID-vkCreateMacOSSurfaceMVK-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkMacOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMacOSSurfaceCreateInfoMVK.html) structure
- [](#VUID-vkCreateMacOSSurfaceMVK-pAllocator-parameter)VUID-vkCreateMacOSSurfaceMVK-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateMacOSSurfaceMVK-pSurface-parameter)VUID-vkCreateMacOSSurfaceMVK-pSurface-parameter  
  `pSurface` **must** be a valid pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_NATIVE_WINDOW_IN_USE_KHR`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_MVK\_macos\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_MVK_macos_surface.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkMacOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMacOSSurfaceCreateInfoMVK.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateMacOSSurfaceMVK).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0