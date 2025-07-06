# vkAcquireNextImageKHR(3) Manual Page

## Name

vkAcquireNextImageKHR - Retrieve the index of the next available presentable image



## [](#_c_specification)C Specification

To acquire an available presentable image to use, and retrieve the index of that image, call:

```c++
// Provided by VK_KHR_swapchain
VkResult vkAcquireNextImageKHR(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    uint64_t                                    timeout,
    VkSemaphore                                 semaphore,
    VkFence                                     fence,
    uint32_t*                                   pImageIndex);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `swapchain` is the non-retired swapchain from which an image is being acquired.
- `timeout` specifies how long the function waits, in nanoseconds, if no image is available.
- `semaphore` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a semaphore to signal.
- `fence` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a fence to signal.
- `pImageIndex` is a pointer to a `uint32_t` in which the index of the next image to use (i.e. an index into the array of images returned by `vkGetSwapchainImagesKHR`) is returned.

## [](#_description)Description

If the `swapchain` has been created with the `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_KHR` flag, the image whose index is returned in `pImageIndex` will be fully backed by memory before this call returns to the application, as if it is bound completely and contiguously to a single `VkDeviceMemory` object.

Valid Usage

- [](#VUID-vkAcquireNextImageKHR-swapchain-01285)VUID-vkAcquireNextImageKHR-swapchain-01285  
  `swapchain` **must** not be in the retired state
- [](#VUID-vkAcquireNextImageKHR-semaphore-01286)VUID-vkAcquireNextImageKHR-semaphore-01286  
  If `semaphore` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** be unsignaled
- [](#VUID-vkAcquireNextImageKHR-semaphore-01779)VUID-vkAcquireNextImageKHR-semaphore-01779  
  If `semaphore` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** not have any uncompleted signal or wait operations pending
- [](#VUID-vkAcquireNextImageKHR-fence-01287)VUID-vkAcquireNextImageKHR-fence-01287  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** be unsignaled
- [](#VUID-vkAcquireNextImageKHR-fence-10066)VUID-vkAcquireNextImageKHR-fence-10066  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** not be associated with any other queue command that has not yet completed execution on that queue
- [](#VUID-vkAcquireNextImageKHR-semaphore-01780)VUID-vkAcquireNextImageKHR-semaphore-01780  
  `semaphore` and `fence` **must** not both be equal to [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkAcquireNextImageKHR-surface-07783)VUID-vkAcquireNextImageKHR-surface-07783  
  If [forward progress](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#swapchain-acquire-forward-progress) cannot be guaranteed for the `surface` used to create the `swapchain` member of `pAcquireInfo`, `timeout` **must** not be `UINT64_MAX`
- [](#VUID-vkAcquireNextImageKHR-semaphore-03265)VUID-vkAcquireNextImageKHR-semaphore-03265  
  `semaphore` **must** have a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY`

Valid Usage (Implicit)

- [](#VUID-vkAcquireNextImageKHR-device-parameter)VUID-vkAcquireNextImageKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkAcquireNextImageKHR-swapchain-parameter)VUID-vkAcquireNextImageKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkAcquireNextImageKHR-semaphore-parameter)VUID-vkAcquireNextImageKHR-semaphore-parameter  
  If `semaphore` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handle
- [](#VUID-vkAcquireNextImageKHR-fence-parameter)VUID-vkAcquireNextImageKHR-fence-parameter  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-vkAcquireNextImageKHR-pImageIndex-parameter)VUID-vkAcquireNextImageKHR-pImageIndex-parameter  
  `pImageIndex` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkAcquireNextImageKHR-swapchain-parent)VUID-vkAcquireNextImageKHR-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from `device`
- [](#VUID-vkAcquireNextImageKHR-semaphore-parent)VUID-vkAcquireNextImageKHR-semaphore-parent  
  If `semaphore` is a valid handle, it **must** have been created, allocated, or retrieved from `device`
- [](#VUID-vkAcquireNextImageKHR-fence-parent)VUID-vkAcquireNextImageKHR-fence-parent  
  If `fence` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized
- Host access to `semaphore` **must** be externally synchronized
- Host access to `fence` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_TIMEOUT`
- `VK_NOT_READY`
- `VK_SUBOPTIMAL_KHR`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_DEVICE_LOST`
- `VK_ERROR_OUT_OF_DATE_KHR`
- `VK_ERROR_SURFACE_LOST_KHR`
- `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT`

## [](#_see_also)See Also

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkAcquireNextImageKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0