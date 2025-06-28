# vkGetSemaphoreCounterValue(3) Manual Page

## Name

vkGetSemaphoreCounterValue - Query the current state of a timeline semaphore



## [](#_c_specification)C Specification

To query the current counter value of a semaphore created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE` from the host, call:

```c++
// Provided by VK_VERSION_1_2
VkResult vkGetSemaphoreCounterValue(
    VkDevice                                    device,
    VkSemaphore                                 semaphore,
    uint64_t*                                   pValue);
```

or the equivalent command

```c++
// Provided by VK_KHR_timeline_semaphore
VkResult vkGetSemaphoreCounterValueKHR(
    VkDevice                                    device,
    VkSemaphore                                 semaphore,
    uint64_t*                                   pValue);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the semaphore.
- `semaphore` is the handle of the semaphore to query.
- `pValue` is a pointer to a 64-bit integer value in which the current counter value of the semaphore is returned.

## [](#_description)Description

Note

If a [queue submission](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-submission) command is pending execution, then the value returned by this command **may** immediately be out of date.

Valid Usage

- [](#VUID-vkGetSemaphoreCounterValue-semaphore-03255)VUID-vkGetSemaphoreCounterValue-semaphore-03255  
  `semaphore` **must** have been created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`

Valid Usage (Implicit)

- [](#VUID-vkGetSemaphoreCounterValue-device-parameter)VUID-vkGetSemaphoreCounterValue-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetSemaphoreCounterValue-semaphore-parameter)VUID-vkGetSemaphoreCounterValue-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handle
- [](#VUID-vkGetSemaphoreCounterValue-pValue-parameter)VUID-vkGetSemaphoreCounterValue-pValue-parameter  
  `pValue` **must** be a valid pointer to a `uint64_t` value
- [](#VUID-vkGetSemaphoreCounterValue-semaphore-parent)VUID-vkGetSemaphoreCounterValue-semaphore-parent  
  `semaphore` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_DEVICE_LOST`

## [](#_see_also)See Also

[VK\_KHR\_timeline\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_timeline_semaphore.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetSemaphoreCounterValue)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0