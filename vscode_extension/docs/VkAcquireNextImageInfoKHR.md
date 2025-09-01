# VkAcquireNextImageInfoKHR(3) Manual Page

## Name

VkAcquireNextImageInfoKHR - Structure specifying parameters of the acquire



## [](#_c_specification)C Specification

The `VkAcquireNextImageInfoKHR` structure is defined as:

```c++
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_swapchain
typedef struct VkAcquireNextImageInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkSwapchainKHR     swapchain;
    uint64_t           timeout;
    VkSemaphore        semaphore;
    VkFence            fence;
    uint32_t           deviceMask;
} VkAcquireNextImageInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `swapchain` is a non-retired swapchain from which an image is acquired.
- `timeout` specifies how long the function waits, in nanoseconds, if no image is available.
- `semaphore` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a semaphore to signal.
- `fence` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a fence to signal.
- `deviceMask` is a mask of physical devices for which the swapchain image will be ready to use when the semaphore or fence is signaled.

## [](#_description)Description

If [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImageKHR.html) is used, the device mask is considered to include all physical devices in the logical device.

Note

[vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImage2KHR.html) signals at most one semaphore, even if the application requests waiting for multiple physical devices to be ready via the `deviceMask`. However, only a single physical device **can** wait on that semaphore, since the semaphore becomes unsignaled when the wait succeeds. For other physical devices to wait for the image to be ready, it is necessary for the application to submit semaphore signal operation(s) to that first physical device to signal additional semaphore(s) after the wait succeeds, which the other physical device(s) **can** wait upon.

Valid Usage

- [](#VUID-VkAcquireNextImageInfoKHR-swapchain-01675)VUID-VkAcquireNextImageInfoKHR-swapchain-01675  
  `swapchain` **must** not be in the retired state
- [](#VUID-VkAcquireNextImageInfoKHR-semaphore-01288)VUID-VkAcquireNextImageInfoKHR-semaphore-01288  
  If `semaphore` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** be unsignaled
- [](#VUID-VkAcquireNextImageInfoKHR-semaphore-01781)VUID-VkAcquireNextImageInfoKHR-semaphore-01781  
  If `semaphore` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** not have any uncompleted signal or wait operations pending
- [](#VUID-VkAcquireNextImageInfoKHR-fence-01289)VUID-VkAcquireNextImageInfoKHR-fence-01289  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** be unsignaled
- [](#VUID-VkAcquireNextImageInfoKHR-fence-10067)VUID-VkAcquireNextImageInfoKHR-fence-10067  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** not be associated with any other queue command that has not yet completed execution on that queue
- [](#VUID-VkAcquireNextImageInfoKHR-semaphore-01782)VUID-VkAcquireNextImageInfoKHR-semaphore-01782  
  `semaphore` and `fence` **must** not both be equal to [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkAcquireNextImageInfoKHR-deviceMask-01290)VUID-VkAcquireNextImageInfoKHR-deviceMask-01290  
  `deviceMask` **must** be a valid device mask
- [](#VUID-VkAcquireNextImageInfoKHR-deviceMask-01291)VUID-VkAcquireNextImageInfoKHR-deviceMask-01291  
  `deviceMask` **must** not be zero
- [](#VUID-VkAcquireNextImageInfoKHR-semaphore-03266)VUID-VkAcquireNextImageInfoKHR-semaphore-03266  
  `semaphore` **must** have a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY`

Valid Usage (Implicit)

- [](#VUID-VkAcquireNextImageInfoKHR-sType-sType)VUID-VkAcquireNextImageInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR`
- [](#VUID-VkAcquireNextImageInfoKHR-pNext-pNext)VUID-VkAcquireNextImageInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkAcquireNextImageInfoKHR-swapchain-parameter)VUID-VkAcquireNextImageInfoKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-VkAcquireNextImageInfoKHR-semaphore-parameter)VUID-VkAcquireNextImageInfoKHR-semaphore-parameter  
  If `semaphore` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handle
- [](#VUID-VkAcquireNextImageInfoKHR-fence-parameter)VUID-VkAcquireNextImageInfoKHR-fence-parameter  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-VkAcquireNextImageInfoKHR-commonparent)VUID-VkAcquireNextImageInfoKHR-commonparent  
  Each of `fence`, `semaphore`, and `swapchain` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized
- Host access to `semaphore` **must** be externally synchronized
- Host access to `fence` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_device\_group](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_device_group.html), [VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html), [vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImage2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAcquireNextImageInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0