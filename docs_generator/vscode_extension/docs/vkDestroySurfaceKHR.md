# vkDestroySurfaceKHR(3) Manual Page

## Name

vkDestroySurfaceKHR - Destroy a VkSurfaceKHR object



## [](#_c_specification)C Specification

To destroy a `VkSurfaceKHR` object, call:

```c++
// Provided by VK_KHR_surface
void vkDestroySurfaceKHR(
    VkInstance                                  instance,
    VkSurfaceKHR                                surface,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `instance` is the instance used to create the surface.
- `surface` is the surface to destroy.
- `pAllocator` is the allocator used for host memory allocated for the surface object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).

## [](#_description)Description

Destroying a `VkSurfaceKHR` merely severs the connection between Vulkan and the native surface, and does not imply destroying the native surface, closing a window, or similar behavior.

Valid Usage

- [](#VUID-vkDestroySurfaceKHR-surface-01266)VUID-vkDestroySurfaceKHR-surface-01266  
  All `VkSwapchainKHR` objects created for `surface` **must** have been destroyed prior to destroying `surface`
- [](#VUID-vkDestroySurfaceKHR-surface-01267)VUID-vkDestroySurfaceKHR-surface-01267  
  If `VkAllocationCallbacks` were provided when `surface` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroySurfaceKHR-surface-01268)VUID-vkDestroySurfaceKHR-surface-01268  
  If no `VkAllocationCallbacks` were provided when `surface` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroySurfaceKHR-instance-parameter)VUID-vkDestroySurfaceKHR-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkDestroySurfaceKHR-surface-parameter)VUID-vkDestroySurfaceKHR-surface-parameter  
  If `surface` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle
- [](#VUID-vkDestroySurfaceKHR-pAllocator-parameter)VUID-vkDestroySurfaceKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroySurfaceKHR-surface-parent)VUID-vkDestroySurfaceKHR-surface-parent  
  If `surface` is a valid handle, it **must** have been created, allocated, or retrieved from `instance`

Host Synchronization

- Host access to `surface` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroySurfaceKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0