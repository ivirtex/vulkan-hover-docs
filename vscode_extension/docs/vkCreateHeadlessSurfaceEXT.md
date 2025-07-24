# vkCreateHeadlessSurfaceEXT(3) Manual Page

## Name

vkCreateHeadlessSurfaceEXT - Create a headless slink:VkSurfaceKHR object



## [](#_c_specification)C Specification

To create a headless `VkSurfaceKHR` object, call:

```c++
// Provided by VK_EXT_headless_surface
VkResult vkCreateHeadlessSurfaceEXT(
    VkInstance                                  instance,
    const VkHeadlessSurfaceCreateInfoEXT*       pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## [](#_parameters)Parameters

- `instance` is the instance to associate the surface with.
- `pCreateInfo` is a pointer to a [VkHeadlessSurfaceCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHeadlessSurfaceCreateInfoEXT.html) structure containing parameters affecting the creation of the surface object.
- `pAllocator` is the allocator used for host memory allocated for the surface object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).
- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle in which the created surface object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateHeadlessSurfaceEXT-instance-parameter)VUID-vkCreateHeadlessSurfaceEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkCreateHeadlessSurfaceEXT-pCreateInfo-parameter)VUID-vkCreateHeadlessSurfaceEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkHeadlessSurfaceCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHeadlessSurfaceCreateInfoEXT.html) structure
- [](#VUID-vkCreateHeadlessSurfaceEXT-pAllocator-parameter)VUID-vkCreateHeadlessSurfaceEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateHeadlessSurfaceEXT-pSurface-parameter)VUID-vkCreateHeadlessSurfaceEXT-pSurface-parameter  
  `pSurface` **must** be a valid pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_headless\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_headless_surface.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkHeadlessSurfaceCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHeadlessSurfaceCreateInfoEXT.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateHeadlessSurfaceEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0