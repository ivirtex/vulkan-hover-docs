# vkCreateViSurfaceNN(3) Manual Page

## Name

vkCreateViSurfaceNN - Create a slink:VkSurfaceKHR object for a VI layer



## [](#_c_specification)C Specification

To create a `VkSurfaceKHR` object for an `nn`::`vi`::`Layer`, query the layer’s native handle using `nn`::`vi`::`GetNativeWindow`, and then call:

```c++
// Provided by VK_NN_vi_surface
VkResult vkCreateViSurfaceNN(
    VkInstance                                  instance,
    const VkViSurfaceCreateInfoNN*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## [](#_parameters)Parameters

- `instance` is the instance with which to associate the surface.
- `pCreateInfo` is a pointer to a [VkViSurfaceCreateInfoNN](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViSurfaceCreateInfoNN.html) structure containing parameters affecting the creation of the surface object.
- `pAllocator` is the allocator used for host memory allocated for the surface object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).
- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle in which the created surface object is returned.

## [](#_description)Description

During the lifetime of a surface created using a particular `nn`::`vi`::`NativeWindowHandle`, applications **must** not attempt to create another surface for the same `nn`::`vi`::`Layer` or attempt to connect to the same `nn`::`vi`::`Layer` through other platform mechanisms.

If the native window is created with a specified size, `currentExtent` will reflect that size. In this case, applications should use the same size for the swapchain’s `imageExtent`. Otherwise, the `currentExtent` will have the special value (0xFFFFFFFF, 0xFFFFFFFF), indicating that applications are expected to choose an appropriate size for the swapchain’s `imageExtent` (e.g., by matching the result of a call to `nn`::`vi`::`GetDisplayResolution`).

Valid Usage (Implicit)

- [](#VUID-vkCreateViSurfaceNN-instance-parameter)VUID-vkCreateViSurfaceNN-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkCreateViSurfaceNN-pCreateInfo-parameter)VUID-vkCreateViSurfaceNN-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkViSurfaceCreateInfoNN](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViSurfaceCreateInfoNN.html) structure
- [](#VUID-vkCreateViSurfaceNN-pAllocator-parameter)VUID-vkCreateViSurfaceNN-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateViSurfaceNN-pSurface-parameter)VUID-vkCreateViSurfaceNN-pSurface-parameter  
  `pSurface` **must** be a valid pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_NATIVE_WINDOW_IN_USE_KHR`

## [](#_see_also)See Also

[VK\_NN\_vi\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NN_vi_surface.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html), [VkViSurfaceCreateInfoNN](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViSurfaceCreateInfoNN.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateViSurfaceNN)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0