# vkSignalSemaphore(3) Manual Page

## Name

vkSignalSemaphore - Signal a timeline semaphore on the host



## [](#_c_specification)C Specification

To signal a semaphore created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE` with a particular counter value, on the host, call:

```c++
// Provided by VK_VERSION_1_2
VkResult vkSignalSemaphore(
    VkDevice                                    device,
    const VkSemaphoreSignalInfo*                pSignalInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_timeline_semaphore
VkResult vkSignalSemaphoreKHR(
    VkDevice                                    device,
    const VkSemaphoreSignalInfo*                pSignalInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the semaphore.
- `pSignalInfo` is a pointer to a [VkSemaphoreSignalInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSignalInfo.html) structure containing information about the signal operation.

## [](#_description)Description

When `vkSignalSemaphore` is executed on the host, it defines and immediately executes a [*semaphore signal operation*](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-signaling) which sets the timeline semaphore to the given value.

The first synchronization scope is defined by the host execution model, but includes execution of `vkSignalSemaphore` on the host and anything that happened-before it.

The second synchronization scope is empty.

Valid Usage (Implicit)

- [](#VUID-vkSignalSemaphore-device-parameter)VUID-vkSignalSemaphore-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSignalSemaphore-pSignalInfo-parameter)VUID-vkSignalSemaphore-pSignalInfo-parameter  
  `pSignalInfo` **must** be a valid pointer to a valid [VkSemaphoreSignalInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSignalInfo.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_timeline\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_timeline_semaphore.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSemaphoreSignalInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreSignalInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSignalSemaphore)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0