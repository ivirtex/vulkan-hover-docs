# vkCreateSurfaceOHOS(3) Manual Page

## Name

vkCreateSurfaceOHOS - Create a slink:VkSurfaceKHR object for an Open Harmony OS native window



## [](#_c_specification)C Specification

To create a `VkSurfaceKHR` object on Open Harmony OS platform, call:

```c++
// Provided by VK_OHOS_surface
VkResult vkCreateSurfaceOHOS(
    VkInstance                                  instance,
    const VkSurfaceCreateInfoOHOS*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## [](#_parameters)Parameters

- `instance` is the instance to associate the surface with.
- `pCreateInfo` is a pointer to a [VkOHSurfaceCreateInfoOHOS](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOHSurfaceCreateInfoOHOS.html) structure containing parameters affecting the creation of the surface object.
- `pAllocator` is the allocator used for host memory allocated for the surface object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).
- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle in which the created surface object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateSurfaceOHOS-instance-parameter)VUID-vkCreateSurfaceOHOS-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkCreateSurfaceOHOS-pCreateInfo-parameter)VUID-vkCreateSurfaceOHOS-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkSurfaceCreateInfoOHOS](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCreateInfoOHOS.html) structure
- [](#VUID-vkCreateSurfaceOHOS-pAllocator-parameter)VUID-vkCreateSurfaceOHOS-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateSurfaceOHOS-pSurface-parameter)VUID-vkCreateSurfaceOHOS-pSurface-parameter  
  `pSurface` **must** be a valid pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_SURFACE_LOST_KHR`

## [](#_see_also)See Also

[VK\_OHOS\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_OHOS_surface.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkSurfaceCreateInfoOHOS](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCreateInfoOHOS.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateSurfaceOHOS)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0