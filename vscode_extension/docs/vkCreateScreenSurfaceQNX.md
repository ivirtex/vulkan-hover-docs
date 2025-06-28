# vkCreateScreenSurfaceQNX(3) Manual Page

## Name

vkCreateScreenSurfaceQNX - Create a slink:VkSurfaceKHR object for a QNX Screen window



## [](#_c_specification)C Specification

To create a `VkSurfaceKHR` object for a QNX Screen surface, call:

```c++
// Provided by VK_QNX_screen_surface
VkResult vkCreateScreenSurfaceQNX(
    VkInstance                                  instance,
    const VkScreenSurfaceCreateInfoQNX*         pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## [](#_parameters)Parameters

- `instance` is the instance to associate the surface with.
- `pCreateInfo` is a pointer to a [VkScreenSurfaceCreateInfoQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenSurfaceCreateInfoQNX.html) structure containing parameters affecting the creation of the surface object.
- `pAllocator` is the allocator used for host memory allocated for the surface object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).
- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle in which the created surface object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateScreenSurfaceQNX-instance-parameter)VUID-vkCreateScreenSurfaceQNX-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkCreateScreenSurfaceQNX-pCreateInfo-parameter)VUID-vkCreateScreenSurfaceQNX-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkScreenSurfaceCreateInfoQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenSurfaceCreateInfoQNX.html) structure
- [](#VUID-vkCreateScreenSurfaceQNX-pAllocator-parameter)VUID-vkCreateScreenSurfaceQNX-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateScreenSurfaceQNX-pSurface-parameter)VUID-vkCreateScreenSurfaceQNX-pSurface-parameter  
  `pSurface` **must** be a valid pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_QNX\_screen\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QNX_screen_surface.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkScreenSurfaceCreateInfoQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenSurfaceCreateInfoQNX.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateScreenSurfaceQNX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0