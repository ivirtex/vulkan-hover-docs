# vkAcquireNextImageKHR(3) Manual Page

## Name

vkAcquireNextImageKHR - Retrieve the index of the next available
presentable image



## <a href="#_c_specification" class="anchor"></a>C Specification

To acquire an available presentable image to use, and retrieve the index
of that image, call:

``` c
// Provided by VK_KHR_swapchain
VkResult vkAcquireNextImageKHR(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    uint64_t                                    timeout,
    VkSemaphore                                 semaphore,
    VkFence                                     fence,
    uint32_t*                                   pImageIndex);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the non-retired swapchain from which an image is being
  acquired.

- `timeout` specifies how long the function waits, in nanoseconds, if no
  image is available.

- `semaphore` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a semaphore to
  signal.

- `fence` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a fence to signal.

- `pImageIndex` is a pointer to a `uint32_t` in which the index of the
  next image to use (i.e. an index into the array of images returned by
  `vkGetSwapchainImagesKHR`) is returned.

## <a href="#_description" class="anchor"></a>Description

If the `swapchain` has been created with the
`VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT` flag, the image
whose index is returned in `pImageIndex` will be fully backed by memory
before this call returns to the application, as if it is bound
completely and contiguously to a single `VkDeviceMemory` object.

Valid Usage

- <a href="#VUID-vkAcquireNextImageKHR-swapchain-01285"
  id="VUID-vkAcquireNextImageKHR-swapchain-01285"></a>
  VUID-vkAcquireNextImageKHR-swapchain-01285  
  `swapchain` **must** not be in the retired state

- <a href="#VUID-vkAcquireNextImageKHR-semaphore-01286"
  id="VUID-vkAcquireNextImageKHR-semaphore-01286"></a>
  VUID-vkAcquireNextImageKHR-semaphore-01286  
  If `semaphore` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) it
  **must** be unsignaled

- <a href="#VUID-vkAcquireNextImageKHR-semaphore-01779"
  id="VUID-vkAcquireNextImageKHR-semaphore-01779"></a>
  VUID-vkAcquireNextImageKHR-semaphore-01779  
  If `semaphore` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) it
  **must** not have any uncompleted signal or wait operations pending

- <a href="#VUID-vkAcquireNextImageKHR-fence-01287"
  id="VUID-vkAcquireNextImageKHR-fence-01287"></a>
  VUID-vkAcquireNextImageKHR-fence-01287  
  If `fence` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) it **must** be
  unsignaled and **must** not be associated with any other queue command
  that has not yet completed execution on that queue

- <a href="#VUID-vkAcquireNextImageKHR-semaphore-01780"
  id="VUID-vkAcquireNextImageKHR-semaphore-01780"></a>
  VUID-vkAcquireNextImageKHR-semaphore-01780  
  `semaphore` and `fence` **must** not both be equal to
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkAcquireNextImageKHR-surface-07783"
  id="VUID-vkAcquireNextImageKHR-surface-07783"></a>
  VUID-vkAcquireNextImageKHR-surface-07783  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#swapchain-acquire-forward-progress"
  target="_blank" rel="noopener">forward progress</a> cannot be
  guaranteed for the `surface` used to create the `swapchain` member of
  `pAcquireInfo`, the `timeout` member of `pAcquireInfo` **must** not be
  `UINT64_MAX`

- <a href="#VUID-vkAcquireNextImageKHR-semaphore-03265"
  id="VUID-vkAcquireNextImageKHR-semaphore-03265"></a>
  VUID-vkAcquireNextImageKHR-semaphore-03265  
  `semaphore` **must** have a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_BINARY`

Valid Usage (Implicit)

- <a href="#VUID-vkAcquireNextImageKHR-device-parameter"
  id="VUID-vkAcquireNextImageKHR-device-parameter"></a>
  VUID-vkAcquireNextImageKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkAcquireNextImageKHR-swapchain-parameter"
  id="VUID-vkAcquireNextImageKHR-swapchain-parameter"></a>
  VUID-vkAcquireNextImageKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkAcquireNextImageKHR-semaphore-parameter"
  id="VUID-vkAcquireNextImageKHR-semaphore-parameter"></a>
  VUID-vkAcquireNextImageKHR-semaphore-parameter  
  If `semaphore` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a href="#VUID-vkAcquireNextImageKHR-fence-parameter"
  id="VUID-vkAcquireNextImageKHR-fence-parameter"></a>
  VUID-vkAcquireNextImageKHR-fence-parameter  
  If `fence` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `fence`
  **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handle

- <a href="#VUID-vkAcquireNextImageKHR-pImageIndex-parameter"
  id="VUID-vkAcquireNextImageKHR-pImageIndex-parameter"></a>
  VUID-vkAcquireNextImageKHR-pImageIndex-parameter  
  `pImageIndex` **must** be a valid pointer to a `uint32_t` value

- <a href="#VUID-vkAcquireNextImageKHR-swapchain-parent"
  id="VUID-vkAcquireNextImageKHR-swapchain-parent"></a>
  VUID-vkAcquireNextImageKHR-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

- <a href="#VUID-vkAcquireNextImageKHR-semaphore-parent"
  id="VUID-vkAcquireNextImageKHR-semaphore-parent"></a>
  VUID-vkAcquireNextImageKHR-semaphore-parent  
  If `semaphore` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

- <a href="#VUID-vkAcquireNextImageKHR-fence-parent"
  id="VUID-vkAcquireNextImageKHR-fence-parent"></a>
  VUID-vkAcquireNextImageKHR-fence-parent  
  If `fence` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

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

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkAcquireNextImageKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
