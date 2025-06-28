# vkDestroySwapchainKHR(3) Manual Page

## Name

vkDestroySwapchainKHR - Destroy a swapchain object



## [](#_c_specification)C Specification

To destroy a swapchain object call:

```c++
// Provided by VK_KHR_swapchain
void vkDestroySwapchainKHR(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) associated with `swapchain`.
- `swapchain` is the swapchain to destroy.
- `pAllocator` is the allocator used for host memory allocated for the swapchain object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).

## [](#_description)Description

The application **must** not destroy a swapchain until after completion of all outstanding operations on images that were acquired from the swapchain. `swapchain` and all associated `VkImage` handles are destroyed, and **must** not be acquired or used any more by the application. The memory of each `VkImage` will only be freed after that image is no longer used by the presentation engine. For example, if one image of the swapchain is being displayed in a window, the memory for that image **may** not be freed until the window is destroyed, or another swapchain is created for the window. Destroying the swapchain does not invalidate the parent `VkSurfaceKHR`, and a new swapchain **can** be created with it.

When a swapchain associated with a display surface is destroyed, if the image most recently presented to the display surface is from the swapchain being destroyed, then either any display resources modified by presenting images from any swapchain associated with the display surface **must** be reverted by the implementation to their state prior to the first present performed on one of these swapchains, or such resources **must** be left in their current state.

If `swapchain` has exclusive full-screen access, it is released before the swapchain is destroyed.

Valid Usage

- [](#VUID-vkDestroySwapchainKHR-swapchain-01282)VUID-vkDestroySwapchainKHR-swapchain-01282  
  All uses of presentable images acquired from `swapchain` **must** have completed execution
- [](#VUID-vkDestroySwapchainKHR-swapchain-01283)VUID-vkDestroySwapchainKHR-swapchain-01283  
  If `VkAllocationCallbacks` were provided when `swapchain` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroySwapchainKHR-swapchain-01284)VUID-vkDestroySwapchainKHR-swapchain-01284  
  If no `VkAllocationCallbacks` were provided when `swapchain` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroySwapchainKHR-device-parameter)VUID-vkDestroySwapchainKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroySwapchainKHR-swapchain-parameter)VUID-vkDestroySwapchainKHR-swapchain-parameter  
  If `swapchain` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkDestroySwapchainKHR-pAllocator-parameter)VUID-vkDestroySwapchainKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroySwapchainKHR-swapchain-parent)VUID-vkDestroySwapchainKHR-swapchain-parent  
  If `swapchain` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroySwapchainKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0