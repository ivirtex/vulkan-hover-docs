# vkCreateXcbSurfaceKHR(3) Manual Page

## Name

vkCreateXcbSurfaceKHR - Create a slink:VkSurfaceKHR object for a X11 window, using the XCB client-side library



## [](#_c_specification)C Specification

To create a `VkSurfaceKHR` object for an X11 window, using the XCB client-side library, call:

```c++
// Provided by VK_KHR_xcb_surface
VkResult vkCreateXcbSurfaceKHR(
    VkInstance                                  instance,
    const VkXcbSurfaceCreateInfoKHR*            pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## [](#_parameters)Parameters

- `instance` is the instance to associate the surface with.
- `pCreateInfo` is a pointer to a [VkXcbSurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkXcbSurfaceCreateInfoKHR.html) structure containing parameters affecting the creation of the surface object.
- `pAllocator` is the allocator used for host memory allocated for the surface object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).
- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle in which the created surface object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateXcbSurfaceKHR-instance-parameter)VUID-vkCreateXcbSurfaceKHR-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkCreateXcbSurfaceKHR-pCreateInfo-parameter)VUID-vkCreateXcbSurfaceKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkXcbSurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkXcbSurfaceCreateInfoKHR.html) structure
- [](#VUID-vkCreateXcbSurfaceKHR-pAllocator-parameter)VUID-vkCreateXcbSurfaceKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateXcbSurfaceKHR-pSurface-parameter)VUID-vkCreateXcbSurfaceKHR-pSurface-parameter  
  `pSurface` **must** be a valid pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_xcb\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_xcb_surface.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html), [VkXcbSurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkXcbSurfaceCreateInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateXcbSurfaceKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0