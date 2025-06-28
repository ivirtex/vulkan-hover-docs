# vkWaitSemaphores(3) Manual Page

## Name

vkWaitSemaphores - Wait for timeline semaphores on the host



## [](#_c_specification)C Specification

To wait for a set of semaphores created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE` to reach particular counter values on the host, call:

```c++
// Provided by VK_VERSION_1_2
VkResult vkWaitSemaphores(
    VkDevice                                    device,
    const VkSemaphoreWaitInfo*                  pWaitInfo,
    uint64_t                                    timeout);
```

or the equivalent command

```c++
// Provided by VK_KHR_timeline_semaphore
VkResult vkWaitSemaphoresKHR(
    VkDevice                                    device,
    const VkSemaphoreWaitInfo*                  pWaitInfo,
    uint64_t                                    timeout);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the semaphores.
- `pWaitInfo` is a pointer to a [VkSemaphoreWaitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitInfo.html) structure containing information about the wait condition.
- `timeout` is the timeout period in units of nanoseconds. `timeout` is adjusted to the closest value allowed by the implementation-dependent timeout accuracy, which **may** be substantially longer than one nanosecond, and **may** be longer than the requested period.

## [](#_description)Description

If the condition is satisfied when `vkWaitSemaphores` is called, then `vkWaitSemaphores` returns immediately. If the condition is not satisfied at the time `vkWaitSemaphores` is called, then `vkWaitSemaphores` will block and wait until the condition is satisfied or the `timeout` has expired, whichever is sooner.

If `timeout` is zero, then `vkWaitSemaphores` does not wait, but simply returns information about the current state of the semaphores. `VK_TIMEOUT` will be returned in this case if the condition is not satisfied, even though no actual wait was performed.

If the condition is satisfied before the `timeout` has expired, `vkWaitSemaphores` returns `VK_SUCCESS`. Otherwise, `vkWaitSemaphores` returns `VK_TIMEOUT` after the `timeout` has expired.

If device loss occurs (see [Lost Device](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-lost-device)) before the timeout has expired, `vkWaitSemaphores` **must** return in finite time with either `VK_SUCCESS` or `VK_ERROR_DEVICE_LOST`.

Valid Usage (Implicit)

- [](#VUID-vkWaitSemaphores-device-parameter)VUID-vkWaitSemaphores-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkWaitSemaphores-pWaitInfo-parameter)VUID-vkWaitSemaphores-pWaitInfo-parameter  
  `pWaitInfo` **must** be a valid pointer to a valid [VkSemaphoreWaitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitInfo.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_TIMEOUT`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_DEVICE_LOST`

## [](#_see_also)See Also

[VK\_KHR\_timeline\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_timeline_semaphore.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSemaphoreWaitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkWaitSemaphores)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0