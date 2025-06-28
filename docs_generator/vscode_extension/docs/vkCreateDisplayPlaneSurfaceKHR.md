# vkCreateDisplayPlaneSurfaceKHR(3) Manual Page

## Name

vkCreateDisplayPlaneSurfaceKHR - Create a slink:VkSurfaceKHR structure representing a display plane and mode



## [](#_c_specification)C Specification

A complete display configuration includes a mode, one or more display planes and any parameters describing their behavior, and parameters describing some aspects of the images associated with those planes. Display surfaces describe the configuration of a single plane within a complete display configuration. To create a `VkSurfaceKHR` object for a display plane, call:

```c++
// Provided by VK_KHR_display
VkResult vkCreateDisplayPlaneSurfaceKHR(
    VkInstance                                  instance,
    const VkDisplaySurfaceCreateInfoKHR*        pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## [](#_parameters)Parameters

- `instance` is the instance corresponding to the physical device the targeted display is on.
- `pCreateInfo` is a pointer to a [VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceCreateInfoKHR.html) structure specifying which mode, plane, and other parameters to use, as described below.
- `pAllocator` is the allocator used for host memory allocated for the surface object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).
- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle in which the created surface is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateDisplayPlaneSurfaceKHR-instance-parameter)VUID-vkCreateDisplayPlaneSurfaceKHR-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkCreateDisplayPlaneSurfaceKHR-pCreateInfo-parameter)VUID-vkCreateDisplayPlaneSurfaceKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceCreateInfoKHR.html) structure
- [](#VUID-vkCreateDisplayPlaneSurfaceKHR-pAllocator-parameter)VUID-vkCreateDisplayPlaneSurfaceKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateDisplayPlaneSurfaceKHR-pSurface-parameter)VUID-vkCreateDisplayPlaneSurfaceKHR-pSurface-parameter  
  `pSurface` **must** be a valid pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_display.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceCreateInfoKHR.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateDisplayPlaneSurfaceKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0