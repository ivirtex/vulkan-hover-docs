# vkCreateIOSSurfaceMVK(3) Manual Page

## Name

vkCreateIOSSurfaceMVK - Create a VkSurfaceKHR object for an iOS UIView



## [](#_c_specification)C Specification

To create a `VkSurfaceKHR` object for an iOS `UIView` or [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html), call:

```c++
// Provided by VK_MVK_ios_surface
VkResult vkCreateIOSSurfaceMVK(
    VkInstance                                  instance,
    const VkIOSSurfaceCreateInfoMVK*            pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## [](#_parameters)Parameters

- `instance` is the instance with which to associate the surface.
- `pCreateInfo` is a pointer to a [VkIOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIOSSurfaceCreateInfoMVK.html) structure containing parameters affecting the creation of the surface object.
- `pAllocator` is the allocator used for host memory allocated for the surface object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).
- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle in which the created surface object is returned.

## [](#_description)Description

Note

The `vkCreateIOSSurfaceMVK` function is considered deprecated and has been superseded by [vkCreateMetalSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMetalSurfaceEXT.html) from the `VK_EXT_metal_surface` extension.

Valid Usage (Implicit)

- [](#VUID-vkCreateIOSSurfaceMVK-instance-parameter)VUID-vkCreateIOSSurfaceMVK-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkCreateIOSSurfaceMVK-pCreateInfo-parameter)VUID-vkCreateIOSSurfaceMVK-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkIOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIOSSurfaceCreateInfoMVK.html) structure
- [](#VUID-vkCreateIOSSurfaceMVK-pAllocator-parameter)VUID-vkCreateIOSSurfaceMVK-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateIOSSurfaceMVK-pSurface-parameter)VUID-vkCreateIOSSurfaceMVK-pSurface-parameter  
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

[VK\_MVK\_ios\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_MVK_ios_surface.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkIOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIOSSurfaceCreateInfoMVK.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateIOSSurfaceMVK)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0