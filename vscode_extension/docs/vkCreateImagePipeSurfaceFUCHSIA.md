# vkCreateImagePipeSurfaceFUCHSIA(3) Manual Page

## Name

vkCreateImagePipeSurfaceFUCHSIA - Create a slink:VkSurfaceKHR object for a Fuchsia ImagePipe



## [](#_c_specification)C Specification

To create a `VkSurfaceKHR` object for a Fuchsia ImagePipe, call:

```c++
// Provided by VK_FUCHSIA_imagepipe_surface
VkResult vkCreateImagePipeSurfaceFUCHSIA(
    VkInstance                                  instance,
    const VkImagePipeSurfaceCreateInfoFUCHSIA*  pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## [](#_parameters)Parameters

- `instance` is the instance to associate with the surface.
- `pCreateInfo` is a pointer to a [VkImagePipeSurfaceCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePipeSurfaceCreateInfoFUCHSIA.html) structure containing parameters affecting the creation of the surface object.
- `pAllocator` is the allocator used for host memory allocated for the surface object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).
- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle in which the created surface object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateImagePipeSurfaceFUCHSIA-instance-parameter)VUID-vkCreateImagePipeSurfaceFUCHSIA-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkCreateImagePipeSurfaceFUCHSIA-pCreateInfo-parameter)VUID-vkCreateImagePipeSurfaceFUCHSIA-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkImagePipeSurfaceCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePipeSurfaceCreateInfoFUCHSIA.html) structure
- [](#VUID-vkCreateImagePipeSurfaceFUCHSIA-pAllocator-parameter)VUID-vkCreateImagePipeSurfaceFUCHSIA-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateImagePipeSurfaceFUCHSIA-pSurface-parameter)VUID-vkCreateImagePipeSurfaceFUCHSIA-pSurface-parameter  
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

[VK\_FUCHSIA\_imagepipe\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_imagepipe_surface.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkImagePipeSurfaceCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePipeSurfaceCreateInfoFUCHSIA.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateImagePipeSurfaceFUCHSIA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0